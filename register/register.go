package register

import (
	"github.com/hashicorp/vault/api"
	"github.com/vu0a/learn-consul/config"
)

func Consul(conCfg *config.ConsulConfig, name, id string, tags []string) error {
	cfg := api.DefaultConfig()
	return nil
}
