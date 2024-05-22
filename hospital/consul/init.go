package consul

import (
	"demo/config"
	"fmt"
	"github.com/hashicorp/consul/api"
)

var Consul *api.Client

func init() {
	var err error
	cfg := api.DefaultConfig()
	// todo: 配置全局变量
	cfg.Address = fmt.Sprintf("%s:%d", config.GlobalConfig.Consul.Host, config.GlobalConfig.Consul.Port)
	Consul, err = api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
}
