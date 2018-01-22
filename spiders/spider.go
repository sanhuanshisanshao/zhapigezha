package spiders

import (
	"fmt"
	"regexp"
	"sync"
	"zhapigezha/models"
)

//提取正则表达式的第二条匹配项既为jpg图片链接
var webp = regexp.MustCompile(`(.+)"(https://.+/view/.+\.jpg)"(.+)`)

//提取下一页的超链接
var href = regexp.MustCompile(`<(.+)href="(https://.+#title-anchor)"(.+)>`)

//Spider分析出来的结果有两种：一种是需要进一步抓取的链接
//例如之前分析的“下一页”的链接，这些东西会被传回 Scheduler
//另一种是需要保存的数据
type Spider struct {
	sync.RWMutex
	PageRow    chan models.RowInfo
	urls       map[string]int
	urlChan    chan models.SourceInfo
	source     map[string]int
	sourceChan chan string
}

func NewSpider() *Spider {
	s := &Spider{
		PageRow:    make(chan models.RowInfo, 1),
		urlChan:    make(chan models.SourceInfo, 1),
		sourceChan: make(chan string, 1),
		urls:       make(map[string]int),
		source:     make(map[string]int),
	}
	go func() {
		s.Analysis()
	}()
	return s
}

func (s *Spider) SetRow(rowInfo models.RowInfo) {
	s.PageRow <- rowInfo
}

//SetSource
func (s *Spider) SetSource(str string) {
	s.Lock()
	defer s.Unlock()
	fmt.Printf("SetSource %v\n", str)
	if s.source[str] == 0 {
		s.source[str] = 1
		s.sourceChan <- str
	}
	return
}

func (s *Spider) GetSourceChan() chan string {
	return s.sourceChan
}

//SetUrls
func (s *Spider) SetUrls(url string, st models.SourceType) {
	s.Lock()
	defer s.Unlock()
	fmt.Printf("SetUrls %v\n", url)
	if s.urls[url] == 0 {
		s.urls[url] = 1
		s.urlChan <- models.SourceInfo{Url: url, SourceType: st}
	}
	return
}

func (s *Spider) GetUrlChan() chan models.SourceInfo {
	return s.urlChan
}

func (s *Spider) Analysis() {
	for row := range s.PageRow {
		if row.SourceType.Code == models.DOUBANIMAGE {
			imgs := webp.FindAllStringSubmatch(row.Row, 100)
			urls := href.FindAllStringSubmatch(row.Row, 100)

			for _, v := range imgs {
				for i, val := range v {
					if i == 2 {
						s.SetSource(val)
					}
				}
			}
			for _, v := range urls {
				for i, val := range v {
					if i == 2 {
						s.SetUrls(val, models.SourceType{Code: row.SourceType.Code})
					}
				}
			}
		} else if row.SourceType.Code == models.WEIBO {
			//TODO:分析微博页面
		}

	}
	return
}
