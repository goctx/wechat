# http-wechat
微信公众平台消息接口服务中间件


## 接入代码

```go
package main

import (
	"flag"
	"fmt"
	"github.com/goctx/http-wechat"
	"github.com/goctx/http-wechat/io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config := &http_wechat.Config{
		Token: "token",
	}
	handleFunc := http_wechat.NewWechat(config).Middleware(handleRequest)
	http.HandleFunc("/", handleFunc)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handleRequest(req *io.Request) interface{} {
	switch req.MsgType {
	case "text":
		return &io.TextResponse{
			Content: "您发送的消息是: " + req.Content,
		}
	case "image":
		return &io.ImageResponse{
			Image: io.Media{
				MediaId: req.MediaId,
			},
		}
	case "voice":
		return &io.VoiceResponse{
			Voice: io.Media{
				MediaId: req.MediaId,
			},
		}
	case "video":
		return &io.VideoResponse{
			Video: io.Video{
				MediaId:     req.MediaId,
				Title:       req.ThumbMediaId,
				Description: strconv.FormatInt(req.MsgId, 10),
			},
		}
	case "shortvideo":
		return &io.VideoResponse{
			Video: io.Video{
				MediaId:     req.MediaId,
				Title:       req.ThumbMediaId,
				Description: strconv.FormatInt(req.MsgId, 10),
			},
		}
	case "link":
		return &io.NewsResponse{
			Articles: io.Articles{
				Articles: []io.Article{
					{
						Title:       req.Title,
						Description: req.Description,
						Url:         req.Url,
						PicUrl:      "https://alicdn.qun.hk/static/img/logo.bc9fd88.png",
					},
				},
			},
		}
	case "event":
		return handleEvent(req)
	default:
		return "success"
	}
}

func handleEvent(req *io.Request) interface{} {
	switch req.Event {
	case "subscribe":
		return &io.TextResponse{
			Content: "谢谢关注: " + req.EventKey,
		}
	case "SCAN":
		return &io.TextResponse{
			Content: "您扫了: " + req.EventKey,
		}
	case "CLICK":
		return &io.TextResponse{
			Content: "您点击了: " + req.EventKey,
		}
	case "VIEW":
		return &io.TextResponse{
			Content: "您点击了: " + req.EventKey,
		}
	default:
		return "success"
	}
}
```