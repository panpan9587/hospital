package router

import (
	"demo/api/etc"
	"demo/config"
	"demo/consul"
	"demo/zaplog"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 启动项目

func Run() {
	zaplog.InitZapLog(etc.ApiConfig.Server.LogPath)
	// 配置中间件
	router := gin.Default()
	// todo: 中间件配置跨域问题
	// router.Use(cors)
	initRouter(router)
	//注册consul
	client := consul.NewRegistryClient(config.GlobalConfig.NacosApiConfig.Host, config.GlobalConfig.NacosApiConfig.Port)
	result := client.FilterService(etc.ApiConfig.Server.Name)
	if len(result) > 0 {
		zap.S().Info("服务已注册，无需重复注册")
	} else {
		err := client.RegisterConsul(etc.ApiConfig.Server.Name, etc.ApiConfig.Server.Tags)
		if err != nil {
			zap.S().Info("服务注册失败")
			panic(err)
		} else {
			zap.S().Info("服务注册成功")
		}
	}
	zap.S().Info("API 服务启动", etc.ApiConfig.Server.Port)
	err := router.Run(fmt.Sprintf(":%d", etc.ApiConfig.Server.Port))
	if err != nil {
		zap.S().Info("API 服务启动失败", err, etc.ApiConfig.Server.Port)
	}
}
