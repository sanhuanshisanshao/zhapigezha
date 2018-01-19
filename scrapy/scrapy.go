package scrapy

import (
	"fmt"
	"strings"
	"time"
	"zhapigezha/downloader"
	http "zhapigezha/httpClient"
	"zhapigezha/scheduler"
	"zhapigezha/spiders"
)

//Scrapy 调度各模块的工作
type Scrapy struct {
	saveSourceUrl string
	down          *downloader.Downloader
	sche          *scheduler.Scheduler
	spider        *spiders.Spider
}

func NewScrapy(url string, down *downloader.Downloader, sche *scheduler.Scheduler, spid *spiders.Spider) *Scrapy {
	return &Scrapy{
		saveSourceUrl: url,
		down:          down,
		sche:          sche,
		spider:        spid,
	}
}

func (s *Scrapy) Start() {

	//scheduler将源地址传给downloader下载该html文件
	go func() {
		for url := range s.sche.GetUrl() {
			s.down.Download(url)
			s.sche.RemoveKey(url)
		}
	}()

	//downloader将html文件的每行解析出来，传给spider分析
	go func() {
		for row := range s.down.GetRowChan() {
			s.spider.SetRow(row)
		}
	}()

	//将spider提取出的源文件传给scheduler
	go func() {
		for url := range s.spider.GetUrlChan() {
			s.sche.PutUrl(url)
		}
	}()

	//
	go func() {
		for sourceUrl := range s.spider.GetSourceChan() {
			bigSizeImg := strings.Replace(sourceUrl, "sqxs", "l", -1)
			bts, err := http.HttpGet(bigSizeImg)
			if err != nil {
				//TODO:
				break
			}
			err = s.down.ToFile(fmt.Sprintf(s.saveSourceUrl+"image_%v.jpg", time.Now().UnixNano()), bts)
			if err != nil {
				//TODO:
				break
			}
		}
	}()
}
