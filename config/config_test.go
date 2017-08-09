package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usrlocalts/consul_client/config"
)

func TestConfig(t *testing.T) {
	config.NewConfig()
	assert.Equal(t, 3000, config.Port())
	assert.Equal(t, "localhost:8500", config.ConsulAddress())
	assert.Equal(t, "localhost", config.NodeIP())
	assert.Equal(t, "http://localhost:3000/", config.ConsulCheckURL())
	assert.Equal(t, "3s", config.ConsulCheckInterval())
	assert.Equal(t, []string{"grpc"}, config.ConsulTags())
	assert.Equal(t, 100, config.ConsulServiceCheckTimeout())
}

func TestConfigOverride(t *testing.T) {
	clientConfig := config.NewConfig()
	clientConfig.NodeID = "example-service"
	clientConfig.NodeIP = "example-ip"
	clientConfig.AppName = "example-service"
	assert.Equal(t, "example-service", config.NodeID())
	assert.Equal(t, "example-ip", config.NodeIP())
	assert.Equal(t, "example-service", config.AppName())
}

