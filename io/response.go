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
	Image Media
}

type Media struct {
	MediaId string
}

type VoiceResponse struct {
	baseResponse
	Voice Media
}

type VideoResponse struct {
	baseResponse
	Video Video
}

type Video struct {
	MediaId     string
	Title       string `xml:"omitempty"`
	Description string `xml:"omitempty"`
}

type MusicResponse struct {
	baseResponse
	Music Music
}

type Music struct {
	Title        string `xml:"omitempty"`
	Description  string `xml:"omitempty"`
	MusicURL     string `xml:"omitempty"`
	HQMusicUrl   string `xml:"omitempty"`
	ThumbMediaId string
}

type NewsResponse struct {
	baseResponse
	ArticleCount int
	Articles     Articles
}

type Articles struct {
	Articles []Article
}

type Article struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type CustomerServiceResponse struct {
	TransInfo TransInfo
}

type TransInfo struct {
	KfAccount string
}
