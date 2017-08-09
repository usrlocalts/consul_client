# Consul Client

This Go Client library allows a service to register and deregister to consul. This library also allows to discover a service from consul by supplying app name and tags

## Getting Started

You can include this library in your glide.yaml/project by using:

`import github.com/usrlocalts/consul_client`

If you want to avoid entering username and password, in your glide.yaml, specifiy as follows:

```
- package: github.com/usrlocalts/consul_client
  repo: git@github.com:usrlocalts/consul_client.git
  vcs: git

```

## Contributing to the library

- `mkdir -p github.com/usrlocalts`
- Clone the repository `git clone git@github.com:usrlocalts/consul_client.git`
- `make test` to run tests

## Setting up go

This service runs on go.

- Install go
    - On OSX run `brew install go`.
    - Follow instructions on https://golang.org/doc/install for other OSes.
- Setup go
      - Make sure that the executable `go` is in your shell's path.
      - Add the following in your .zshrc or .bashrc: (where `<workspace_dir>` is the directory in
        which you'll checkout your code)
- Run Test
    make test

```
GOPATH=<workspace_dir>
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

## Usage

After you include the library in your application,

##### Create a client by doing the following:

```
func ConsulClientConfig() *consulClientConfig.Config {
	consulClientConfig := consulClientConfig.NewConfig()
	clientConfig.NodeID = "example-service"
    clientConfig.NodeIP = "example-ip"
    clientConfig.AppName = "example-service"
	return consulClientConfig
}

consulClient, err := client.NewConsul(ConsulClientConfig())


Assuming all these confgurations are in your config file

```

##### The default configurations when you do a `consulClientConfig.NewConfig()` are:

```
	Port:                      3000,
	ConsulAddress:             "localhost:8500",
	NodeIP:                    "localhost",
	ConsulCheckURL:            "http://localhost:3000/",
	ConsulCheckInterval:       "3s",
	ConsulTags:                []string{"grpc"},
	ConsulServiceCheckTimeout: 100,

```

##### You can edit these config by editing the configs:
 
```
	Port                      int
	ConsulAddress             string
	NodeID                    string
	NodeIP                    string
	AppName                   string
	ConsulCheckURL            string
	ConsulCheckInterval       string
	ConsulTags                []string
	ConsulServiceCheckTimeout int
	
	by
	
	clientConfig.Port = 3100

```

##### Methods available

```
consulClient.RegisterToConsul()
consulClient.DeRegisterFromConsul(serviceID)
consulClient.Discover("example-service", []string{"grpc"})
```