package discovery

import (
	"fmt"
	"strings"

	"github.com/hashicorp/consul/api"
)

type ServiceDiscovery interface {
	Discover(name string, tags []string) ([]string, error)
}

type serviceDiscovery struct {
	consulAgentAddress string
}

func New(consulAgentAddress string) ServiceDiscovery {
	return &serviceDiscovery{consulAgentAddress: consulAgentAddress}
}

func (self serviceDiscovery) Discover(name string, tags []string) ([]string, error) {

	consulConfig := api.DefaultConfig()
	consulConfig.Address = self.consulAgentAddress
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return []string{}, err
	}
	catalog := client.Catalog()
	discoveredServices, _, err := catalog.Service(name, strings.Join(tags[:], ","), &api.QueryOptions{})
	if err != nil {
		return []string{}, err
	}
	urls := make([]string, len(discoveredServices))
	for i, service := range discoveredServices {
		urls[i] = fmt.Sprintf("%s:%d", service.Address, service.ServicePort)
	}
	return urls, nil
}
