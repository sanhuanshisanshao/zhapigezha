package models

import (
	"github.com/PuerkitoBio/goquery"
)

var (
	DOUBANIMAGE  Source = 1
	WEIBO        Source = 2
	WEIBO_COOKIE string
)

type Source int

type SourceType struct {
	//code=1 豆瓣电影图片
	//code=2 新浪微博
	//TODO:weibo friends,images,fans
	Code Source `json:"code"`
}

type Document struct {
	Doc        goquery.Document `json:"doc"`
	SourceType SourceType       `json:"source_type"`
}

type SourceInfo struct {
	Url        string     `json:"url"`
	SourceType SourceType `json:"source_type"`
}

type RowInfo struct {
	Row        string     `json:"row"`
	SourceType SourceType `json:"source_type"`
}

type HtmlInfo struct {
	Html       string     `json:"html"`
	SourceType SourceType `json:"source_type"`
}
