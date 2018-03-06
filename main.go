package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"zhapigezha/downloader"
	"zhapigezha/httpServer"
	"zhapigezha/models"
	"zhapigezha/scheduler"
	"zhapigezha/scrapy"
	"zhapigezha/spiders"
)

//定义命令行参数
var (
	//种子地址
	URLS []string
	//保存资源的路径
	SAVEPATH string
	//种子地址类型
	SOURCETYPE models.Source
)

func init() {
	var c string
	var sourceType string
	flag.StringVar(&SAVEPATH, "path", "C:/Users/Gao/Desktop/photos/", "")
	flag.StringVar(&c, "url", "https://movie.douban.com/photos/photo/2167242744/", "")
	flag.StringVar(&sourceType, "type", "1", "")

	if sourceType == "1" {
		SOURCETYPE = models.DOUBANIMAGE
	} else if sourceType == "2" {
		SOURCETYPE = models.WEIBO
	} else {
		//默认
		SOURCETYPE = models.DOUBANIMAGE
	}

	slice := strings.Split(c, ",")
	URLS = slice
}

func main() {
	var st = models.SourceType{Code: SOURCETYPE}
	sche := scheduler.NewScheduler()
	down := downloader.NewDownloader()
	spider := spiders.NewSpider()

	go func() {
		server := httpServer.NewHTTPServer(sche)
		if err := http.ListenAndServe(":5665", server.GetRouter()); err != nil {
			fmt.Println("start http server error:%v", err)
		}
	}()

	for _, v := range URLS {
		//豆瓣,微博 资源入口
		sche.PutUrl(v, st)
	}

	s := scrapy.NewScrapy(SAVEPATH, down, sche, spider)
	s.Start()

	//quit when receive end signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	fmt.Printf("receice quit signal %v\n", <-ch)
}
