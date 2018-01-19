package httpServer

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"zhapigezha/scheduler"
)

type httpServer struct {
	sche   *scheduler.Scheduler
	router http.Handler
}

func NewHTTPServer(sche *scheduler.Scheduler) *httpServer {
	router := httprouter.New()
	router.HandleMethodNotAllowed = true

	s := &httpServer{
		sche:   sche,
		router: router,
	}

	router.Handle("GET", "/ping",
		func(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			s.pingHandle(resp, req)
		})

	router.Handle("GET", "/put/:url", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		url := params.ByName("url")
		//TODO:参数校验
		url = strings.Replace(url, ";", "/", -1)
		fmt.Println("url:", url)
		s.put(writer, request, url)
	})

	router.Handle("GET", "/get/sche", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		s.getScheduler(writer, request)
	})

	return s
}

func (hs *httpServer) GetRouter() http.Handler {
	return hs.router
}

func (s *httpServer) pingHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
}

func (s *httpServer) put(w http.ResponseWriter, req *http.Request, url string) {
	s.sche.PutUrl(url)
	fmt.Fprintf(w, "Success")
}

func (s *httpServer) getScheduler(w http.ResponseWriter, req *http.Request) {
	list := s.sche.RangeMap()
	fmt.Println(list)
	bytes, _ := json.Marshal(&list)
	w.Write(bytes)
}
