package io

import "encoding/xml"

type Response struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
}

type TextResponse struct {
	Response
	Content string
}

type ImageResponse struct {
	Response
	MediaId string
}

type VoiceResponse struct {
	Response
	MediaId string
}

type VideoResponse struct {
	Response
	MediaId     string
	Title       string `xml:"omitempty"`
	Description string `xml:"omitempty"`
}

type MusicResponse struct {
	Response
	Title        string `xml:"omitempty"`
	Description  string `xml:"omitempty"`
	MusicURL     string `xml:"omitempty"`
	HQMusicUrl   string `xml:"omitempty"`
	ThumbMediaId string
}

type NewsResponse struct {
	Response
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
