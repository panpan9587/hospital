package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	NacosApiConfig struct {
		NamespaceId string `yaml:"NamespaceId"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DataId      string `yaml:"DataId"`
		Group       string `yaml:"group"`
	} `yaml:"NacosApiConfig"`
	NacosRpcUserConfig struct {
		NamespaceId string `yaml:"NamespaceId"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DataId      string `yaml:"DataId"`
		Group       string `yaml:"group"`
	} `yaml:"NacosRpcUserConfig"`
	NacosRpcAdvisoryConfig struct {
		NamespaceId string `yaml:"NamespaceId"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DataId      string `yaml:"DataId"`
		Group       string `yaml:"group"`
	} `yaml:"NacosRpcAdvisoryConfig"`
	Consul struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"consul"`
	Mysql struct {
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DBName      string `yaml:"DBName"`
		MaxIdleConn int    `yaml:"MaxIdleConn"`
		MaxOpenConn int    `yaml:"MaxOpenConn"`
	} `yaml:"mysql"`
	Mongodb struct {
		Username   string `yaml:"username"`   //用户名
		Password   string `yaml:"password"`   //密码
		Host       string `yaml:"host"`       //主机
		Port       int    `yaml:"port"`       //端口
		Database   string `yaml:"database"`   //数据库
		Collection string `yaml:"collection"` //集合
	} `yaml:"mongodb"`
	Log struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"log"`
}

var GlobalConfig Config

// 初始化配置文件
func init() {
	//viper.SetConfigFile("D:/github_daima/hospital/config/config.yaml")
	viper.SetConfigFile("D:/go/src/hospital/hospital/config/config.yaml")
	viper.ReadInConfig()
	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalln("yaml etc parsing error", err)
	}
}