package downloader

import (
	"fmt"
	"testing"
)

func TestNewDownloader(t *testing.T) {
	d := NewDownloader()
	d.Download("https://movie.douban.com/photos/photo/2501683809/")
	rows := d.GetRowChan()
	for v := range rows {
		fmt.Println(v)
	}
}
