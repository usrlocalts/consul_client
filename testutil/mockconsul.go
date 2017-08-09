package testutil

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
)

type MockConsul struct {
	response string
}

func (self *MockConsul) ResetHandlerResponse(response string) {
	self.response = response
}

func (self *MockConsul) GetMockServer() (*httptest.Server) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if self.response == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			io.WriteString(w, self.response)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	return server
}

func NewMockConsul(response string) *MockConsul {
	return &MockConsul{response: response}
}

func SetupMockConsulResponse(response string) *httptest.Server {
	return NewMockConsul(response).GetMockServer()
}

func MockConsulResponse(name string, tags []string, addresses ...net.Addr) *httptest.Server {
	var response []string
	for _, address := range addresses {
		addressData := strings.Split(address.String(), ":")
		tagData := strings.Join(tags, "\",\"")
		consulData := fmt.Sprintf("{\"Address\": \"%s\", \"ServiceName\": \"%s\", \"ServiceTags\": [\"%s\"], \"ServiceAddress\": \"%s\", \"ServicePort\": %s }", addressData[0], name, tagData, addressData[0], addressData[1])
		response = append(response, consulData)
	}
	return SetupMockConsulResponse(fmt.Sprintf("[%s]", strings.Join(response, ",")))
}
