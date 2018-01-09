package downloader

import (
	"fmt"
	"regexp"
	"sync"
	http "zhapigezha/httpClient"
)

type Downloader struct {
	sync.WaitGroup
	resultChan chan string
	rowChan    chan string
}

func NewDownloader() *Downloader {
	d := &Downloader{
		resultChan: make(chan string, 1),
		rowChan:    make(chan string, 1),
	}

	go func() {
		d.split()
	}()
	return d
}

//ToFile 将下载的html文档写入文件
//func (d *Downloader) ToFile(name string, s string) error {
//	f, err := os.Create(name)
//	defer f.Close()
//
//	if err != nil {
//		return err
//	}
//	_, err = f.WriteString(s + "\n")
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (d *Downloader) GetRowChan() chan string {
	return d.rowChan
}

//Download 下载指定的url html资源，存放至resultChan
func (d *Downloader) Download(url string) error {
	bytes, err := http.HttpGet(url)
	if err != nil {
		return err
	}
	if len(bytes) > 0 {
		d.resultChan <- string(bytes)
		return nil
	}
	return fmt.Errorf("download result is nil")
}

func (d *Downloader) split() {
	//提取html的行
	reg := regexp.MustCompile(`\n`)
	for v := range d.resultChan {
		list := reg.Split(v, -1)
		for _, v := range list {
			d.rowChan <- v
		}
	}
}
