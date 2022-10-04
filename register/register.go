package register

import (
	"github.com/hashicorp/consul/api"
	"github.com/vu0a/learn-consul/config"
)

func Consul(conCfg *config.ConsulConfig, address, name, id string, tags []string, port int) error {
	cfg := api.DefaultConfig()
	cfg.Address = conCfg.Addr

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	// 注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Tags = tags
	registration.Address = address
	registration.Port = port

	// // 检查对象
	// check:=&api.AgentServiceCheck{
	// 	HTTP: "http://foo.bar/health",
	// 	Timeout: "5s",
	// 	Interval: "5s",
	// 	DeregisterCriticalServiceAfter: "10s",
	// }
	// registration.Check = check

	return client.Agent().ServiceRegister(registration)
}
