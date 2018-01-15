package io

import "encoding/xml"

type Response interface {
	ServeWechat()
}

type baseResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
}

type TextResponse struct {
	baseResponse
	Content string
}

func (*TextResponse) ServeWechat() {

}

type ImageResponse struct {
	baseResponse
	MediaId string
}

func (*ImageResponse) ServeWechat() {

}

type VoiceResponse struct {
	baseResponse
	MediaId string
}

func (*VoiceResponse) ServeWechat() {

}

type VideoResponse struct {
	baseResponse
	MediaId     string
	Title       string `xml:"omitempty"`
	Description string `xml:"omitempty"`
}

func (*VideoResponse) ServeWechat() {

}

type MusicResponse struct {
	baseResponse
	Title        string `xml:"omitempty"`
	Description  string `xml:"omitempty"`
	MusicURL     string `xml:"omitempty"`
	HQMusicUrl   string `xml:"omitempty"`
	ThumbMediaId string
}

func (*MusicResponse) ServeWechat() {

}

type NewsResponse struct {
	baseResponse
	ArticleCount int
	Articles     []Article
}

func (*NewsResponse) ServeWechat() {

}

type Article struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}
