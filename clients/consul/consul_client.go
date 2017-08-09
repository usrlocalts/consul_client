package consul

import (
	consul "github.com/hashicorp/consul/api"
	"github.com/usrlocalts/consul_client/config"
	"github.com/usrlocalts/consul_client/clients/consul/discovery"
)

type Client interface {
	RegisterToConsul() error
	DeRegisterFromConsul(string) error
	Discover(name string, tags []string) ([]string, error)
}

type client struct {
	consul           *consul.Client
	consulConf       *config.Config
	serviceDiscovery discovery.ServiceDiscovery
}

func NewConsulClient(consulConf *config.Config) (Client, error) {
	config := consul.DefaultConfig()
	config.Address = consulConf.ConsulAddress
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c, consulConf: consulConf, serviceDiscovery: discovery.New(consulConf.ConsulAddress)}, nil
}

func (c *client) RegisterToConsul() error {
	reg := &consul.AgentServiceRegistration{
		ID:      c.consulConf.NodeID,
		Name:    c.consulConf.AppName,
		Port:    c.consulConf.Port,
		Address: c.consulConf.NodeIP,
		Tags:    c.consulConf.ConsulTags,
		Check: &consul.AgentServiceCheck{
			Interval: c.consulConf.ConsulCheckInterval,
			HTTP:     c.consulConf.ConsulCheckURL,
		},
	}

	return c.consul.Agent().ServiceRegister(reg)
}

func (c *client) DeRegisterFromConsul(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

func (c *client) Discover(name string, tags []string) ([]string, error){
	return c.serviceDiscovery.Discover(name, tags)
}