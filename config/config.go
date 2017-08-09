package config

type Config struct {
	Port                      int
	ConsulAddress             string
	NodeID                    string
	NodeIP                    string
	AppName                   string
	ConsulCheckURL            string
	ConsulCheckInterval       string
	ConsulTags                []string
	ConsulServiceCheckTimeout int
}

var appConfig *Config

func load(config *Config) {
	appConfig = config
}

func Port() int {
	return appConfig.Port
}

func ConsulAddress() string {
	return appConfig.ConsulAddress
}

func NodeID() string {
	return appConfig.NodeID
}

func NodeIP() string {
	return appConfig.NodeIP
}

func AppName() string {
	return appConfig.AppName
}

func ConsulTags() []string {
	return appConfig.ConsulTags
}

func ConsulCheckURL() string {
	return appConfig.ConsulCheckURL
}

func ConsulCheckInterval() string {
	return appConfig.ConsulCheckInterval
}

func ConsulServiceCheckTimeout() int {
	return appConfig.ConsulServiceCheckTimeout
}

func NewConfig() *Config {
	newConfig := &Config{
		Port:                      3000,
		ConsulAddress:             "localhost:8500",
		NodeIP:                    "localhost",
		ConsulCheckURL:            "http://localhost:3000/",
		ConsulCheckInterval:       "3s",
		ConsulTags:                []string{"grpc"},
		ConsulServiceCheckTimeout: 100,
	}
	load(newConfig)
	return newConfig
}
