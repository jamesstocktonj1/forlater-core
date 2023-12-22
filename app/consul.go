package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/jamesstocktonj1/forlater-core/internal/config"
)

type ConsulClient struct {
	*api.Client
}

type ConsulConfig struct {
	*api.KV
}

func NewConsulClient() (*ConsulClient, error) {
	conf := &api.Config{
		Address: os.Getenv("CONSUL_ADDR"),
	}
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &ConsulClient{client}, nil
}

func (c *ConsulClient) Register() error {
	host, _ := os.Hostname()

	serviceCheck := &api.AgentServiceCheck{
		Interval: "30s",
		Timeout:  "60s",
		HTTP:     "http://" + host + ":8000/ping",
	}

	agentService := &api.AgentServiceRegistration{
		ID:    host,
		Name:  ConsulServiceName,
		Check: serviceCheck,
	}

	return c.Agent().ServiceRegister(agentService)
}

func (c *ConsulClient) Deregister() error {
	return c.Agent().ServiceDeregister(ConsulServiceName)
}

func (c *ConsulClient) LoadConfig() (*config.ServerConfig, error) {

	return nil, nil
}

func (c *ConsulConfig) GetValue(key string) (string, error) {
	p, _, err := c.Get(key, nil)
	if err != nil {
		return "", err
	}
	return string(p.Value), nil
}

func (c *ConsulConfig) SetValue(key string, value string) error {
	p := api.KVPair{
		Key:   key,
		Value: []byte(value),
	}
	_, err := c.Put(&p, nil)
	return err
}

func (c *ConsulConfig) GetInt(key string) (int, error) {
	p, _, err := c.Get(key, nil)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(p.Value))
}

func (c *ConsulConfig) SetInt(key string, value int) error {
	p := api.KVPair{
		Key:   key,
		Value: []byte(fmt.Sprintf("%d", value)),
	}
	_, err := c.Put(&p, nil)
	return err
}
