package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Consul *ConsulConfig `json:"consul,omitempty"`
}
type ConsulConfig struct {
	Addr string `json:"addr,omitempty"`
}

var Cfg *Config

func InitFrom(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	Cfg = new(Config)
	if err := json.Unmarshal(buf, Cfg); err != nil {
		return err
	}
	return nil
}
