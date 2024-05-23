package main

import (
	"demo/consul"
	proto "demo/rpc/advisorySrv/advisoryclient"
	"demo/rpc/advisorySrv/etc"
	"demo/rpc/advisorySrv/service"
	"demo/zaplog"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	zaplog.InitZapLog(etc.RpcConfig.Server.LogPath)

	//初始化数据库
	s := grpc.NewServer()
	proto.RegisterAdvisoryServer(s, service.Advisory{})
	//端口监听
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", etc.RpcConfig.Server.Host, etc.RpcConfig.Server.Port))
	if err != nil {
		zap.S().Info("端口监听失败", err)
	}
	//注册consul
	client := consul.NewRegistryClient(etc.RpcConfig.Server.Host, etc.RpcConfig.Server.Port)
	result := client.FilterService(etc.RpcConfig.Server.Name)
	if len(result) > 0 {
		zap.S().Info("服务已注册，无需重复注册")
	} else {
		err = client.RegisterConsul(etc.RpcConfig.Server.Name, etc.RpcConfig.Server.Tags)
		if err != nil {
			zap.S().Info("服务注册失败")
			panic(err)
		} else {
			zap.S().Info("服务注册成功")
		}
	}
	zap.S().Info("advisory Srv 服务启动", etc.RpcConfig.Server.Port)
	err = s.Serve(lis)
	if err != nil {
		zap.S().Info("advisory Srv 服务启动失败")
	}

}
