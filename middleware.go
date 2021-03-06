package http_wechat

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/goctx/http-wechat/io"
	"log"
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
	EnabledLog     bool
}

type Wechat struct {
	token          string
	appid          string
	encodingAESKey string
	enableLog      bool
}

func NewWechat(config *Config) *Wechat {
	return &Wechat{
		token:          config.Token,
		appid:          config.AppId,
		encodingAESKey: config.EncodingAESKey,
		enableLog:      config.EnabledLog,
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
		if w.enableLog {
			log.Printf("%s %s\n", r.Method, r.URL.String())
		}
		nonce := r.URL.Query().Get("nonce")
		timestamp := r.URL.Query().Get("timestamp")
		signatureGet := r.URL.Query().Get("signature")
		if w.MakeSign(nonce, timestamp) != signatureGet {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		if r.Method == "GET" {
			echoStr := r.URL.Query().Get("echostr")
			fmt.Fprint(rw, echoStr)
			return
		}
		if r.Method == "POST" {
			// POST处理
			rw.Header().Set("Content-Type", "text/xml;charset=UTF-8")
			var req io.Request
			decoder := xml.NewDecoder(r.Body)
			r.Body.Close()
			if err := decoder.Decode(&req); err != nil {
				http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
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
			case *io.ImageResponse:
				v.MsgType = "image"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = time.Now().Unix()
				encoder := xml.NewEncoder(rw)
				encoder.Encode(&v)
			case *io.VoiceResponse:
				v.MsgType = "voice"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = time.Now().Unix()
				encoder := xml.NewEncoder(rw)
				encoder.Encode(&v)
			case *io.VideoResponse:
				v.MsgType = "video"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = time.Now().Unix()
				encoder := xml.NewEncoder(rw)
				encoder.Encode(&v)
			case *io.MusicResponse:
				v.MsgType = "music"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = time.Now().Unix()
				encoder := xml.NewEncoder(rw)
				encoder.Encode(&v)
			case *io.NewsResponse:
				v.MsgType = "news"
				v.ArticleCount = len(v.Articles.Articles)
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = time.Now().Unix()
				encoder := xml.NewEncoder(rw)
				encoder.Encode(&v)
			case *io.CustomerServiceResponse:
				v.MsgType = "transfer_customer_service"
				v.FromUserName = req.ToUserName
				v.ToUserName = req.FromUserName
				v.CreateTime = time.Now().Unix()
				encoder := xml.NewEncoder(rw)
				encoder.Encode(&v)
			case string:
				fmt.Fprint(rw, v)
			default:
				fmt.Fprint(rw, "success")
			}
		}
	}
}
