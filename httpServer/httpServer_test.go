package httpServer

import (
	"net/http"
	"testing"
)

func TestNewHTTPServer(t *testing.T) {
	server := NewHTTPServer()
	err := http.ListenAndServe(":3636", server.GetRouter())
	if err != nil {
		t.Fatalf("fail")
	}
}
