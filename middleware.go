package http_wechat

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/xialeistudio/wechatmp/io"
	"net/http"
	"sort"
	"strings"
)

type HandleFunc func(request *io.Request) *io.Response

type Config struct {
	Token          string
	AppId          string
	EncodingAESKey string
}

type Wechat struct {
	Token          string
	AppId          string
	EncodingAESKey string
}

func NewWechat(config *Config) *Wechat {
	return &Wechat{
		Token:          config.Token,
		AppId:          config.AppId,
		EncodingAESKey: config.EncodingAESKey,
	}
}

func (w *Wechat) MakeSign(nonce, timestamp string) string {
	sl := []string{w.Token, nonce, timestamp}
	sort.Strings(sl)
	s := sha1.New()
	s.Write([]byte(strings.Join(sl, "")))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func (w *Wechat) Middleware(handleFunc HandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		nonce := r.URL.Query().Get("nonce")
		timestamp := r.URL.Query().Get("timestamp")
		signatureGet := r.URL.Query().Get("signature")
		if w.MakeSign(nonce, timestamp) != signatureGet {
			rw.WriteHeader(401)
			fmt.Fprint(rw, "Invalid Signature")
			return
		}
		if r.Method == "GET" {
			echoStr := r.URL.Query().Get("echostr")
			fmt.Fprint(rw, echoStr)
			return
		}
		// POST处理
		rw.Header().Set("Content-Type", "text/xml;charset=UTF-8")
		var req io.Request
		decoder := xml.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			panic(err)
		}
		response := handleFunc(&req)
		encoder := xml.NewEncoder(rw)
		if err := encoder.Encode(&response); err != nil {
			panic(err)
		}
	}
}
