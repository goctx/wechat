package io

import "encoding/xml"

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

type ImageResponse struct {
	baseResponse
	MediaId string
}

type VoiceResponse struct {
	baseResponse
	MediaId string
}

type VideoResponse struct {
	baseResponse
	MediaId     string
	Title       string `xml:"omitempty"`
	Description string `xml:"omitempty"`
}

type MusicResponse struct {
	baseResponse
	Title        string `xml:"omitempty"`
	Description  string `xml:"omitempty"`
	MusicURL     string `xml:"omitempty"`
	HQMusicUrl   string `xml:"omitempty"`
	ThumbMediaId string
}

type NewsResponse struct {
	baseResponse
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
