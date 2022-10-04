package service

import (
	"github.com/hashicorp/consul/api"
	"github.com/vu0a/learn-consul/config"
)

func All(conCfg *config.ConsulConfig) (map[string]*api.AgentService, error) {
	cfg := api.DefaultConfig()
	cfg.Address = conCfg.Addr

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return client.Agent().Services()
}
