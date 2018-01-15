package io

import "encoding/xml"

type Response interface {
	ServeWechat()
}

type BaseResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
}

func (*BaseResponse) ServeWechat() {

}

type TextResponse struct {
	BaseResponse
	Content string
}

type ImageResponse struct {
	BaseResponse
	MediaId string
}

type VoiceResponse struct {
	BaseResponse
	MediaId string
}

type VideoResponse struct {
	BaseResponse
	MediaId     string
	Title       string `xml:"omitempty"`
	Description string `xml:"omitempty"`
}

type MusicResponse struct {
	BaseResponse
	Title        string `xml:"omitempty"`
	Description  string `xml:"omitempty"`
	MusicURL     string `xml:"omitempty"`
	HQMusicUrl   string `xml:"omitempty"`
	ThumbMediaId string
}

type NewsResponse struct {
	BaseResponse
	ArticleCount int
	Articles     []Article
}

type Article struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}
