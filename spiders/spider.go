package spiders

import (
	"fmt"
	"regexp"
	"sync"
)

//提取正则表达式的第二条匹配项既为webp图片链接
var webp = regexp.MustCompile(`(.+)"(https://.+\.jpg)"(.+)`)

//提取下一页的超链接
var href = regexp.MustCompile(`<(.+)href="(https://.+#title-anchor)"(.+)>`)

//Spider分析出来的结果有两种：一种是需要进一步抓取的链接
//例如之前分析的“下一页”的链接，这些东西会被传回 Scheduler
//另一种是需要保存的数据
type Spider struct {
	sync.RWMutex

	//待分析的page row信息
	htmlPage chan string
	//待存入scheduler的链接
	urls map[string]int
	//待保存的资源文件链接
	source map[string]int
}

func NewSpider() *Spider {
	s := &Spider{
		htmlPage: make(chan string, 100),
		urls:     make(map[string]int),
		source:   make(map[string]int),
	}
	go func() {
		s.Analysis()
	}()
	return s
}

func (s *Spider) SetRow(str string) {
	s.htmlPage <- str
}

//SetSource
func (s *Spider) SetSource(str string) {
	s.Lock()
	defer s.Unlock()

	s.source[str] = 1
}

func (s *Spider) GetSource() string {
	s.Lock()
	defer s.Unlock()

	for k, _ := range s.source {
		delete(s.source, k)
		return k
	}
	return ""
}

//SetUrls
func (s *Spider) SetUrls(url string) {
	s.Lock()
	defer s.Unlock()

	s.urls[url] = 1
}

//GetUrl
func (s *Spider) GetUrl() string {
	s.Lock()
	defer s.Unlock()

	for k, _ := range s.urls {
		delete(s.urls, k)
		return k
	}
	return ""
}

//Analysis 筛选出页面中的资源文件和源文件（html）
//1:<img src 是图片(jpg,png一般是icon,webp是大图正是我们需要的)
//2:href 超链接，若不是https/http开头，则是当前url+href字段为新链接
func (s *Spider) Analysis() {
	for row := range s.htmlPage {
		imgs := webp.FindAllStringSubmatch(row, 100)
		urls := href.FindAllStringSubmatch(row, 100)

		//FIXME:
		for _, v := range imgs {
			for i, val := range v {
				if i == 2 {
					fmt.Println("get images ", val)
					s.SetSource(val)
				}
			}
		}

		for _, v := range urls {
			for i, val := range v {
				if i == 2 {
					fmt.Println("get url ", val)
					s.SetUrls(val)
				}
			}
		}
	}
	return
}
