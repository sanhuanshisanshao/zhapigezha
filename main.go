package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"zhapigezha/downloader"
	"zhapigezha/scheduler"
	"zhapigezha/scrapy"
	"zhapigezha/spiders"
)

var URLS = []string{
	`https://movie.douban.com/photos/photo/2508831489/`,
	`https://movie.douban.com/photos/photo/2506035906/`}

var SAVEURL = "C:/Users/Gao/Desktop/jinyan/"

//A user-defined flag type, a slice of durations.
type interval []string

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *interval) Set(value string) error {

	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, dt)
	}
	return nil
}

var urlsFlag interval

func init() {
	flag.Var(&urlsFlag, "c", "for url list")
}

func main() {

	fmt.Printf("%v\n", urlsFlag.String())

	sche := scheduler.NewScheduler()
	down := downloader.NewDownloader()
	spider := spiders.NewSpider()

	for _, v := range URLS {
		sche.PutUrl(v)
	}

	s := scrapy.NewScrapy(SAVEURL, down, sche, spider)
	s.Start()

	//quit when receive end signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	fmt.Printf("receice quit signal %v\n", <-ch)

	fmt.Println("----end-----")

}
