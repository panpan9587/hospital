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
	UserSrv struct {
		Name string `yaml:"Name"`
	} `yaml:"user_srv"`
	RegistrationSrv struct {
		Name string `yaml:"Name"`
	} `yaml:"registration_srv"`
	SearchSrv struct {
		Name string `yaml:"Name"`
	} `yaml:"search_srv"`
	Token struct {
		AccessSecret string `yaml:"accessSecret"`
		AccessExpire int64  `yaml:"accessExpire"`
	} `yaml:"token"`
	Sms struct {
		AccessKeyId     string `yaml:"AccessKeyId"`
		AccessKeySecret string `yaml:"AccessKeySecret"`
		Endpoint        string `yaml:"Endpoint"`
	}
	Auth struct {
		SecretID  string `yaml:"SecretID"`
		SecretKey string `yaml:"SecretKey"`
	} `yaml:"auth"`
}

var ApiConfig Config

// 接入对应的配置文件
func init() {
	fmt.Println(config.GlobalConfig, "ddd")
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.GlobalConfig.NacosApiConfig.NamespaceId, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      config.GlobalConfig.NacosApiConfig.Host,
			ContextPath: "/nacos",
			Port:        uint64(config.GlobalConfig.NacosApiConfig.Port),
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
		DataId: config.GlobalConfig.NacosApiConfig.DataId,
		Group:  config.GlobalConfig.NacosApiConfig.Group})
	err = yaml.Unmarshal([]byte(content), &ApiConfig)
	if err != nil {
		log.Fatalln("nacos 配置文件解码失败")
	}
}
