package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config Application configuration structure
type Config struct {
	RPCURL string
	IP     string
	Port   uint16
}

// EnvConfig stores config from ENV
type EnvConfig struct {
	RPCURL     string `envconfig:"rpc_url" required:"false" default:"http://127.0.0.1:9001/RPC2"`
	ListenIP   string `envconfig:"listen_ip" required:"false" default:"0.0.0.0"`
	ListenPort uint16 `envconfig:"listen_port" required:"false" default:"8080"`
}

// GetConfig get Application configuration
func GetConfig() *Config {
	var s EnvConfig
	err := envconfig.Process("status", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	config := Config{
		RPCURL: s.RPCURL,
		IP:     s.ListenIP,
		Port:   s.ListenPort,
	}

	return &config
}

// ListenAddress returns listen string
func (c *Config) ListenAddress() string {
	return fmt.Sprintf("%s:%d", c.IP, c.Port)
}
