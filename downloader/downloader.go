package downloader

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
	http "zhapigezha/httpClient"
	"zhapigezha/models"
)

type Downloader struct {
	sync.WaitGroup
	resultChan chan models.HtmlInfo
	rowChan    chan models.RowInfo
}

func NewDownloader() *Downloader {
	d := &Downloader{
		resultChan: make(chan models.HtmlInfo, 1),
		rowChan:    make(chan models.RowInfo, 1),
	}

	go func() {
		d.split()
	}()
	return d
}

//TODO:批量生成
//ToFile 将下载的html文档写入文件
//name  xxx.jpg
func (d *Downloader) ToFile(name string, s []byte) error {
	f, err := os.Create(name)
	defer f.Close()

	if err != nil {
		return err
	}
	reader := bytes.NewReader(s)
	_, err = io.Copy(f, reader)
	return nil
}

func (d *Downloader) GetRowChan() chan models.RowInfo {
	return d.rowChan
}

//Download 下载指定的url html资源，将整个页面作为string存放至resultChan 中
func (d *Downloader) Download(url string, st models.SourceType) error {
	bts, err := http.HttpGet(url)
	if err != nil {
		return err
	}
	if len(bts) > 0 {
		d.resultChan <- models.HtmlInfo{Html: string(bts), SourceType: st}
		return nil
	}
	return fmt.Errorf("download result is nil")
}

//split 如果是豆瓣的资源则,按行切割html页面,其他网站的html页面暂时不切割
func (d *Downloader) split() {
	//提取html的行
	reg := regexp.MustCompile(`\n`)
	for v := range d.resultChan {
		//豆瓣
		if v.SourceType.Code == models.DOUBANIMAGE {
			list := reg.Split(v.Html, -1)
			for _, val := range list {
				d.rowChan <- models.RowInfo{Row: val, SourceType: v.SourceType}
			}
		} else if v.SourceType.Code == models.WEIBO {
			//TODO:微博
		} else {
			//其他
		}

	}
}
