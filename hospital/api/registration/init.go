package registration

import (
	"demo/api/etc"
	"demo/consul"
	proto "demo/rpc/registerationSrv/registerationclient"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// todo: 初始化连接rpc
var RegistrationSrv proto.RegisterationClient

func init() {
	//注册consul
	client := consul.NewRegistryClient(etc.ApiConfig.Server.Host, etc.ApiConfig.Server.Port)
	result := client.FilterService(etc.ApiConfig.RegistrationSrv.Name)
	var host string
	var port int
	for _, service := range result {
		host = service.Address
		port = service.Port
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Info("Registration srv 连接失败", err)
	}
	RegistrationSrv = proto.NewRegisterationClient(conn)
}
