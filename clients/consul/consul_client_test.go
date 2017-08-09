package consul_test

import (
	"testing"
	"github.com/usrlocalts/consul_client/config"
	"github.com/usrlocalts/consul_client/clients/consul"
	"github.com/usrlocalts/consul_client/testutil"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestNewConsulClient(t *testing.T) {
	clientConfig := config.NewConfig()
	consulClient, err := consul.NewConsulClient(clientConfig)
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(consulClient).String(), "*consul.client")
}

func TestConsulClientWithWrongProtocolShouldReturnError(t *testing.T) {
	clientConfig := config.NewConfig()
	clientConfig.ConsulAddress = "blah://blah.com:8600"
	_, err := consul.NewConsulClient(clientConfig)
	assert.Error(t, err)
}

func TestConsulClientShouldRegisterToConsul(t *testing.T) {
	clientConfig := config.NewConfig()
	clientConfig.NodeID = "example-service"
	clientConfig.NodeIP = "example-ip"
	clientConfig.AppName = "example-service"

	lis := testutil.StartTestConsulServer()
	address := lis.Addr()
	server := testutil.MockConsulResponse("example-service", []string{"grpc"}, address)
	clientConfig.ConsulAddress = server.Listener.Addr().String()
	defer server.Close()

	consulClient, err := consul.NewConsulClient(clientConfig)
	assert.NoError(t, err)
	err = consulClient.RegisterToConsul()
	assert.NoError(t, err)
}

func TestConsulClientShouldNotRegisterToConsulWithNoServerRunningError(t *testing.T) {
	clientConfig := config.NewConfig()
	clientConfig.NodeID = "example-service"
	clientConfig.NodeIP = "example-ip"
	clientConfig.AppName = "example-service"

	consulClient, err := consul.NewConsulClient(clientConfig)
	assert.NoError(t, err)
	err = consulClient.RegisterToConsul()
	assert.Error(t, err)
}

func TestConsulClientShouldDeRegisterFromConsul(t *testing.T) {
	clientConfig := config.NewConfig()
	clientConfig.NodeID = "example-service"
	clientConfig.NodeIP = "example-ip"
	clientConfig.AppName = "example-service"

	lis := testutil.StartTestConsulServer()
	address := lis.Addr()
	server := testutil.MockConsulResponse("example-service", []string{"grpc"}, address)
	clientConfig.ConsulAddress = server.Listener.Addr().String()
	defer server.Close()

	consulClient, err := consul.NewConsulClient(clientConfig)
	assert.NoError(t, err)
	err = consulClient.DeRegisterFromConsul("12345")
	assert.NoError(t, err)
}

func TestConsulClientShouldNotDeRegisterFromConsulWithNoServerRunningError(t *testing.T) {
	clientConfig := config.NewConfig()
	clientConfig.NodeID = "example-service"
	clientConfig.NodeIP = "example-ip"
	clientConfig.AppName = "example-service"

	consulClient, err := consul.NewConsulClient(clientConfig)
	assert.NoError(t, err)
	err = consulClient.DeRegisterFromConsul("12345")
	assert.Error(t, err)
}

func TestDiscoverMultipleServicesByNameAndTag(t *testing.T) {
	mockConsulResponse := "[{\"Address\": \"127.0.0.1\", \"ServiceName\": \"example-service\", \"ServiceTags\": [], \"ServiceAddress\": \"localhost\", \"ServicePort\": 5000 }, {\"Address\": \"127.0.0.1\", \"ServiceName\": \"example-service\", \"ServiceTags\": [], \"ServiceAddress\": \"localhost\", \"ServicePort\": 5001 }]"
	server := testutil.SetupMockConsulResponse(mockConsulResponse)
	defer server.Close()

	clientConfig := config.NewConfig()
	clientConfig.NodeID = "example-service"
	clientConfig.NodeIP = "example-ip"
	clientConfig.AppName = "example-service"

	clientConfig.ConsulAddress = server.Listener.Addr().String()
	defer server.Close()

	consulClient, err := consul.NewConsulClient(clientConfig)

	exampleServices, _ := consulClient.Discover("example-service", []string{"grpc"})

	assert.NoError(t, err)
	assert.Equal(t, 2, len(exampleServices))
	assert.Equal(t, "127.0.0.1:5000", exampleServices[0])
	assert.Equal(t, "127.0.0.1:5001", exampleServices[1])
}

func TestDiscoverServiceIfAddressNotFound(t *testing.T) {
	clientConfig := config.NewConfig()
	clientConfig.NodeID = "example-service"
	clientConfig.NodeIP = "example-ip"
	clientConfig.AppName = "example-service"

	consulClient, err := consul.NewConsulClient(clientConfig)
	assert.NoError(t, err)

	exampleServices, _ := consulClient.Discover("example-service", []string{"grpc"})
	assert.Equal(t, 0, len(exampleServices))
}
