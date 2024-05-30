package etc

import (
	"demo/config"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
	"log"
)

type Config struct {
	Server struct {
		Host    string `yaml:"Host"`
		Port    int    `yaml:"Port"`
		Name    string `yaml:"Name"`
		Tags    string `yaml:"Tags"`
		LogPath string `yaml:"LogPath"`
	} `yaml:"server"`
}

var RpcConfig Config

// 接入对应的配置文件
func init() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.GlobalConfig.NacosRpcCaseConfig.NamespaceId, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      config.GlobalConfig.NacosRpcCaseConfig.Host,
			ContextPath: "/nacos",
			Port:        uint64(config.GlobalConfig.NacosRpcCaseConfig.Port),
			Scheme:      "http",
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: config.GlobalConfig.NacosRpcCaseConfig.DataId,
		Group:  config.GlobalConfig.NacosRpcCaseConfig.Group})

	fmt.Println(config.GlobalConfig.NacosRpcCaseConfig.Host, "1111")

	err = yaml.Unmarshal([]byte(content), &RpcConfig)
	if err != nil {
		log.Fatalln("nacos caserpc配置文件解码失败")
	}
}
