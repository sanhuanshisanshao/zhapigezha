package downloader

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"os"
	"regexp"
	"sync"
	http "zhapigezha/httpClient"
	"zhapigezha/models"
)

type Downloader struct {
	sync.WaitGroup
	//resultChan 保存整个HTML文件（豆瓣）
	resultChan chan models.HtmlInfo
	//rowChan 保存resultChan的每行字符串（豆瓣）
	rowChan chan models.RowInfo
	//goquery document
	document chan models.Document
}

func NewDownloader() *Downloader {
	d := &Downloader{
		resultChan: make(chan models.HtmlInfo, 1),
		rowChan:    make(chan models.RowInfo, 1),
		document:   make(chan models.Document, 1),
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

//Download
func (d *Downloader) Download(url string, st models.SourceType) error {
	if st.Code == models.DOUBANIMAGE {
		bts, err := http.HttpGet(url)
		if err != nil {
			return err
		}
		if len(bts) > 0 {
			d.resultChan <- models.HtmlInfo{Html: string(bts), SourceType: st}
			return nil
		}
		return fmt.Errorf("download result is nil")
	} else if st.Code == models.WEIBO {
		//new document
		doc, err := goquery.NewDocument(url, models.WEIBO_COOKIE)
		if err != nil {
			return fmt.Errorf("goquery new document error %v", err)
		}

		doc.Find(".c").Each(func(i int, selection *goquery.Selection) {
			// For each item found, get the band and title

			//只转发而没用评论改转发的微博，包含class="cmt",class="ctt"
			cmt := selection.Find("div").Find(".cmt").Text()
			if cmt == "" {
				//原创微博
				ctt := selection.Find("div").Find(".ctt").Text()
				if ctt != "" {
					fmt.Printf("原创微博: %d: %s \n", i, ctt)
				}
			} else {
				//转发出处
				from := selection.Find("div").Find(".cmt").First().Text()
				//转发理由
				reason := selection.Find("div").Last().Text()
				//转发的微博正文
				ctt := selection.Find("div").Find(".ctt").Text()
				//转发微博图片地址
				image, err := selection.Find("div").First().Last().Html()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("转发微博: %d: %s == %s == %s == %s \n", i, reason, ctt, from, image)
				//fmt.Printf("图片div：%s", image)
			}
		})

	}
	return fmt.Errorf("unknown error type")
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
		}
	}
}
