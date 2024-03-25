package config

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/caarlos0/env/v6"
	consulapi "github.com/hashicorp/consul/api"
	"net"
	"strconv"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/22 9:08
 * @Desc:
 */

var (
	consulConfig     *ConsulConfig
	enableConsulFlag = false
)

type ConsulConfig struct {
	Host     string `mapstructure:"host" json:"host" env:"CONSUL_HOST"`
	Port     int    `mapstructure:"port" json:"port" env:"CONSUL_PORT"`
	ACLToken string `mapstructure:"acl_token" json:"acl_token" env:"CONSUL_ACL_TOKEN"`
	Key      string `mapstructure:"key" json:"key" env:"CONSUL_KEY"`
}

func newConsulConfig() *ConsulConfig {
	return &ConsulConfig{
		Host: "127.0.0.1",
		Port: 8500,
		Key:  "wxcommerce/pmsmodel",
	}
}

func initConsulConfig() (*ConsulConfig, error) {
	conf := newConsulConfig()
	if err := env.Parse(conf); err != nil {
		return nil, fmt.Errorf("fatal error parsing env config: %s", err.Error())
	}

	if conf.Key == "" || conf.Host == "" || conf.Port == 0 {
		return nil, fmt.Errorf("consul params error, all params must not null")
	}
	enableConsulFlag = true
	return conf, nil
}

func loadConfigFromConsul() error {
	if conf, err := initConsulConfig(); err != nil {
		return err
	} else {
		consulConfig = conf
	}

	cfg := consulapi.DefaultConfig()
	cfg.Address = net.JoinHostPort(consulConfig.Host, strconv.Itoa(consulConfig.Port))
	consulClient, err := consulapi.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("new consul client failed: %s", err.Error())
	}

	content, _, err := consulClient.KV().Get(consulConfig.Key, nil)
	if err != nil {
		return fmt.Errorf("consul kv failed: %s", err.Error())
	}

	err = sonic.Unmarshal(content.Value, serverConfig)
	if err != nil {
		return fmt.Errorf("sonic unmarshal config failed: %s", err.Error())
	}

	if serverConfig.Host == "" {
		serverConfig.Host = "0.0.0.0"
	}

	if err := env.Parse(serverConfig); err != nil {
		return fmt.Errorf("fatal error parsing env config: %w", err)
	}

	return validate.Struct(serverConfig)
}
