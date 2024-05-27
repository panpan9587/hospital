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
	NacosRpcRegistrationConfig struct {
		NamespaceId string `yaml:"NamespaceId"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DataId      string `yaml:"DataId"`
		Group       string `yaml:"group"`
	} `yaml:"NacosRpcRegistrationConfig"`
	NacosRpcHealthConfig struct {
		NamespaceId string `yaml:"NamespaceId"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DataId      string `yaml:"DataId"`
		Group       string `yaml:"group"`
	} `yaml:"NacosRpcHealthConfig"`
	NacosRpcDiagnosisConfig struct {
		NamespaceId string `yaml:"NamespaceId"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DataId      string `yaml:"DataId"`
		Group       string `yaml:"group"`
	} `yaml:"NacosRpcDiagnosisConfig"`
	NacosRpcCaseConfig struct {
		NamespaceId string `yaml:"NamespaceId"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		DataId      string `yaml:"DataId"`
		Group       string `yaml:"group"`
	} `yaml:"NacosRpcCaseConfig"`


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
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Db   int    `yaml:"db"`
	} `yaml:"redis"`
	Rabbitmq struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"` //用户名
		Password string `yaml:"password"` //密码
		Exchange string `yaml:"exchange"` //交换机
	} `yaml:"rabbitmq"`
	Es struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"es"`

	Log struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"log"`
}

var GlobalConfig Config

// 初始化配置文件
func init() {
	viper.SetConfigFile("D:/gos/src/bookProject/hospital/hospital/config/config.yaml")
	viper.ReadInConfig()
	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalln("yaml etc parsing error", err)
	}
}
