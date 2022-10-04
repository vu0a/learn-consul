package service

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/vu0a/learn-consul/config"
)

func Filter(conCfg *config.ConsulConfig, filter string) (map[string]*api.AgentService, error) {
	cfg := api.DefaultConfig()
	cfg.Address = conCfg.Addr

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return client.Agent().ServicesWithFilter(filter)
}

func EqNameFilter(serviceName string) string {
	return fmt.Sprintf(`Service == %q`, serviceName)
}
