package httpClient

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(urlStr string, reqBody string) (respBytes []byte, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(reqBody))
	req.Header.Add("content-type", "application/json")
	//req.Response.Header.Add("Content-Type", "text/html;charset=utf-8")

	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return body, nil
}
