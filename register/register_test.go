package register_test

import (
	"testing"

	"github.com/vu0a/learn-consul/config"
	"github.com/vu0a/learn-consul/register"
)

func TestRegister(t *testing.T) {
	if err := config.InitFrom("../config.json"); err != nil {
		t.Fatal(err)
	}
	if err := register.Consul(config.Cfg.Consul, "[2a:**:1]", "learn-consul", "learn-consul", []string{"learn-consul"}, 12345); err != nil {
		t.Fatal("register failed", err)
	}
}
