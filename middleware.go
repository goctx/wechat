package http_wechat

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/goctx/http-wechat/io"
	"net/http"
	"sort"
	"strings"
	"time"
)

type HandleFunc func(request *io.Request) interface{}

type Config struct {
	Token          string
	AppId          string
	EncodingAESKey string
}

type Wechat struct {
	token          string
	appid          string
	encodingAESKey string
}

func NewWechat(config *Config) *Wechat {
	return &Wechat{
		token:          config.Token,
		appid:          config.AppId,
		encodingAESKey: config.EncodingAESKey,
	}
}

func (w *Wechat) MakeSign(nonce, timestamp string) string {
	sl := []string{w.token, nonce, timestamp}
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
		switch v := response.(type) {
		case *io.TextResponse:
			v.MsgType = "text"
			v.FromUserName = req.ToUserName
			v.ToUserName = req.FromUserName
			v.CreateTime = time.Now().Unix()
			encoder := xml.NewEncoder(rw)
			encoder.Encode(&v)
			break
		case *io.ImageResponse:
			v.MsgType = "image"
			v.FromUserName = req.ToUserName
			v.ToUserName = req.FromUserName
			v.CreateTime = time.Now().Unix()
			encoder := xml.NewEncoder(rw)
			encoder.Encode(&v)
			break
		case *io.VoiceResponse:
			v.MsgType = "voice"
			v.FromUserName = req.ToUserName
			v.ToUserName = req.FromUserName
			v.CreateTime = time.Now().Unix()
			encoder := xml.NewEncoder(rw)
			encoder.Encode(&v)
			break
		case *io.VideoResponse:
			v.MsgType = "video"
			v.FromUserName = req.ToUserName
			v.ToUserName = req.FromUserName
			v.CreateTime = time.Now().Unix()
			encoder := xml.NewEncoder(rw)
			encoder.Encode(&v)
			break
		case *io.MusicResponse:
			v.MsgType = "music"
			v.FromUserName = req.ToUserName
			v.ToUserName = req.FromUserName
			v.CreateTime = time.Now().Unix()
			encoder := xml.NewEncoder(rw)
			encoder.Encode(&v)
			break
		case *io.NewsResponse:
			v.MsgType = "news"
			v.ArticleCount = len(v.Articles.Articles)
			v.FromUserName = req.ToUserName
			v.ToUserName = req.FromUserName
			v.CreateTime = time.Now().Unix()
			encoder := xml.NewEncoder(rw)
			encoder.Encode(&v)
			break
		case string:
			fmt.Fprint(rw, v)
		default:
			fmt.Fprint(rw, "success")
			break
		}
	}
}
