package discovery_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usrlocalts/consul_client/clients/consul/discovery"
	"github.com/usrlocalts/consul_client/testutil"
)

func TestDiscoverMultipleServicesByNameAndTag(t *testing.T) {
	mockConsulResponse := "[{\"Address\": \"127.0.0.1\", \"ServiceName\": \"example-service\", \"ServiceTags\": [], \"ServiceAddress\": \"localhost\", \"ServicePort\": 5000 }, {\"Address\": \"127.0.0.1\", \"ServiceName\": \"example-service\", \"ServiceTags\": [], \"ServiceAddress\": \"localhost\", \"ServicePort\": 5001 }]"
	server := testutil.SetupMockConsulResponse(mockConsulResponse)
	defer server.Close()
	sdd := discovery.New(server.Listener.Addr().String())

	exampleServices, err := sdd.Discover("example-service", []string{"grpc"})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(exampleServices))
	assert.Equal(t, "127.0.0.1:5000", exampleServices[0])
	assert.Equal(t, "127.0.0.1:5001", exampleServices[1])
}

func TestDiscoverServiceIfAddressNotFound(t *testing.T) {
	mockConsulResponse := "[]"
	server := testutil.SetupMockConsulResponse(mockConsulResponse)
	defer server.Close()
	sdd := discovery.New(server.Listener.Addr().String())

	exampleServices, _ := sdd.Discover("example-service", []string{"grpc"})
	assert.Equal(t, 0, len(exampleServices))
}
