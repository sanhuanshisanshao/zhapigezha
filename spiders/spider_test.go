package spiders

import (
	"fmt"

	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
	"zhapigezha/downloader"
	http "zhapigezha/httpClient"
	"zhapigezha/scheduler"
)

var url = `https://movie.douban.com/photos/photo/1016491895/`

func TestSpider_Analysis(t *testing.T) {
	sche := scheduler.NewScheduler()
	down := downloader.NewDownloader()
	spider := NewSpider()

	sche.PutUrl(url)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range sche.GetUrl() {
			fmt.Printf("get url %v\n", v)
			down.Download(v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range down.GetRowChan() {
			spider.SetRow(v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for reUrl := range spider.GetUrlChan() {
			fmt.Println("reurl:", reUrl)
			sche.PutUrl(reUrl)
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()

		for s := range spider.GetSourceChan() {

			times := time.Now().Nanosecond()
			if s != "" {
				f, _ := os.Create(fmt.Sprintf("C:/Users/Gao/Desktop/jinyan/images_%v.jpg", times))
				fmt.Printf("开始储存图片URL：%v at %v\n", s, time.Now())
				//大图地址
				bigImg := strings.Replace(s, "sqxs", "l", -1)
				bts, err := http.HttpGet(bigImg)

				if err != nil {
					fmt.Printf("http获取图片失败：%v\n", err)
				}

				br := bytes.NewReader(bts)
				_, err = io.Copy(f, br)
				if err != nil {
					fmt.Printf("存储图片失败：%v\n", err)
				}

				f.Close()
			}

		}
	}()
	wg.Wait()
}
