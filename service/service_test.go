package service_test

import (
	"testing"

	"github.com/vu0a/learn-consul/config"
	"github.com/vu0a/learn-consul/service"
)

func TestAll(t *testing.T) {
	if err := config.InitFrom("../config.json"); err != nil {
		t.Fatal(err)
	}
	svcs, err := service.All(config.Cfg.Consul)
	if err != nil {
		t.Fatal("get all services failed: ", err)
	}
	for key, val := range svcs {
		t.Logf("%s: %#v\n", key, val)
	}
}

func TestFilter(t *testing.T) {
	if err := config.InitFrom("../config.json"); err != nil {
		t.Fatal(err)
	}
	filter := service.EqNameFilter("learn-consul")
	svcs, err := service.Filter(config.Cfg.Consul, filter)
	if err != nil {
		t.Fatal("get all services failed: ", err)
	}
	for key, val := range svcs {
		t.Logf("%s: %#v\n", key, val)
	}
}
