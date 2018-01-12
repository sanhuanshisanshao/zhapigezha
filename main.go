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
	"zhapigezha/scheduler"
	"zhapigezha/scrapy"
	"zhapigezha/spiders"
)

var URLS []string

var SAVEURL string

func init() {
	var c string
	flag.StringVar(&SAVEURL, "path", "C:/Users/Gao/Desktop/jinyan/", "save file path")
	flag.StringVar(&c, "url", "", "url")
	slice := strings.Split(c, ",")
	URLS = slice
}

func main() {

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
		sche.PutUrl(v)
	}

	s := scrapy.NewScrapy(SAVEURL, down, sche, spider)
	s.Start()

	//quit when receive end signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	fmt.Printf("receice quit signal %v\n", <-ch)
}
