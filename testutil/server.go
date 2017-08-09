package testutil

import (
	"net/http/httptest"
	"net/http"
	"net"
)

func NewTestServer(router http.Handler) *httptest.Server {
	return httptest.NewServer(router)
}

func StartTestConsulServer() (net.Listener) {
	lis, _ := net.Listen("tcp", "localhost:0")
	return lis
}
