package httpClient

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHttpGet(t *testing.T) {
	url := "https://www.nasa.gov/topics/earth/images/index.html"
	bytes, err := HttpGet(url)
	if err != nil {
		t.Fatalf("http get url error %v", err)
	}
	s := string(bytes)

	fmt.Printf("%v", s)
}

func TestRegexp(t *testing.T) {
	str := `<img src="https://img3.doubanio.com/view/photo/l/public/p455467226.webp"  />`
	//str := `        "https://img1.doubanio.com/view/photo/l/public/p2182128677.webp",`
	//webp image
	reg := regexp.MustCompile(`(.+)"(https://.+/view/.+\.webp)"(.+)`)

	str2 := `<a class="mainphoto" href="https://movie.douban.com/photos/photo/2182128677/#title-anchor" title="点击查看下一张">`
	//href
	reg2 := regexp.MustCompile(`<(.+)href="(https://.+#title-anchor)"(.+)>`)

	match := reg.FindAllStringSubmatch(str, 1000)

	match2 := reg2.FindAllStringSubmatch(str2, 1000)

	for _, v := range match {
		for k, val := range v {
			if k == 2 {
				fmt.Printf("%v\n", val)
			}

		}
	}

	for _, v := range match2 {
		for k, val := range v {
			if k == 2 {
				fmt.Printf("%v\n", val)
			}
		}
	}

}
