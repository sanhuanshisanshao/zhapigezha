package spiders

import (
	"fmt"

	"source-finder/downloader"
	"source-finder/scheduler"
	"sync"
	"testing"
	"time"
)

var url = `https://movie.douban.com/photos/photo/2501683809/`

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
		//TODO:source to scheduler
		for true {
			reUrl := spider.GetUrl()
			fmt.Println("reurl:", reUrl)
			sche.PutUrl(reUrl)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		//times := time.Now().Nanosecond()
		//	f, _ := os.Create(fmt.Sprintf("C:/Users/Gao/Desktop/images_%v.txt", times))
		for true {
			s := spider.GetSource()
			fmt.Printf("储存图片URL：%v at %v\n", s, time.Now())
			//_, err := f.WriteString(fmt.Sprintf("%v\n", s))
			//if err != nil {
			//	fmt.Printf("write []byte to file error:%v", err)
			//}
		}
	}()
	wg.Wait()
}
