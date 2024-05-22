package service

/*
	info(
	desc: "用户中心逻辑"
	author: "panpan"
	email: "918716975@qq.com"
)
*/
import (
	"context"
	proto "demo/rpc/userSrv/userclient"
)

type User struct {
	proto.UnimplementedUserServer
}

func (u User) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	//TODO 用户注册
	return nil, nil
}

func (u User) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	//TODO 用户登录
	return nil, nil
}
