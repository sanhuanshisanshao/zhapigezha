package httpServer

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type httpServer struct {
	//ctx    *Context
	router http.Handler
}

func NewHTTPServer() *httpServer {

	router := httprouter.New()
	router.HandleMethodNotAllowed = true

	s := &httpServer{
		router: router,
	}

	router.Handle("GET", "/ping",
		func(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
			//TODO:
			s.pingHandle(resp, req)
		})

	return s
}

func (s *httpServer) pingHandle(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	return "OK", nil
}
