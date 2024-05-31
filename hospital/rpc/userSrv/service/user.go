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
	"demo/rpc/common"
	"demo/rpc/model/mysql"
	"demo/rpc/userSrv/etc"
	proto "demo/rpc/userSrv/userclient"
	"errors"
	"github.com/mervick/aes-everywhere/go/aes256"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct {
	proto.UnimplementedUserServer
}

func (u User) GetUserAuth(ctx context.Context, request *proto.GetUserAuthRequest) (*proto.GetUserAuthResponse, error) {
	//TODO 获取用户实名信息
	auth, err := mysql.GetUserAuth(int(request.UserId))
	if err != nil {
		zap.S().Info("获取用户实名信息失败", err)
		return nil, errors.New("获取用户实名信息失败")
	}
	return &proto.GetUserAuthResponse{
		AuthInfo: &proto.UserAuthInfo{
			RealName: auth.RealName,
			IdNumber: auth.IdNumber,
			Status:   int32(auth.Status),
		},
	}, nil
}

func (u User) AddUserAuth(ctx context.Context, request *proto.AddUserAuthRequest) (*proto.AddUserAuthResponse, error) {
	//TODO 实名认证接口
	auth := mysql.UserAuth{
		UserID:   int(request.UserId),
		RealName: request.RealName,
		IdNumber: request.IdNumber,
	}
	// todo: 判断用户是否已经实名
	userAuth, err := mysql.GetUserAuth(int(request.UserId))
	if err != nil {
		zap.S().Info("用户信息获取失败", err)
		return nil, errors.New("用户信息获取失败")
	}
	if userAuth.ID != 0 {
		return nil, errors.New("用户已经实名认证")
	}
	tx := mysql.DB.Begin()
	res, err := common.AuthUser(auth.IdNumber, auth.RealName)
	if err != nil {
		tx.Rollback()
		zap.S().Info("实名信息认证失败", err)
		return nil, errors.New("实名信息认证失败")
	}
	if int(res["error_code"].(float64)) != 0 {
		if !res["result"].(map[string]interface{})["isok"].(bool) {
			tx.Rollback()
			zap.S().Info("实名信息认证失败", err)
			return nil, errors.New("实名信息认证失败")
		}
	}
	err = tx.Create(&auth).Error
	if err != nil {
		tx.Rollback()
		zap.S().Info("实名信息认证失败", err)
		return nil, errors.New("实名信息认证失败")
	}
	err = tx.Model(mysql.User{}).Where("id = ?", request.UserId).Update("status", "3").Error
	if err != nil {
		tx.Rollback()
		zap.S().Info("用户状态修改失败", err)
		return nil, errors.New("获取用户状态失败")
	}
	tx.Commit()
	// 同步修改用户状态
	return &proto.AddUserAuthResponse{
		AuthInfo: &proto.UserAuthInfo{
			RealName: auth.RealName,
			IdNumber: auth.IdNumber,
			Status:   int32(auth.Status),
		},
	}, nil
}

func (u User) GetUserInfo(ctx context.Context, request *proto.GetUserInfoRequest) (*proto.GetUserInfoResponse, error) {
	//TODO 返回用户详情信息
	user, err := mysql.GetUserById(int(request.UserId))
	if err != nil {
		return nil, errors.New("用户信息查询失败")
	}
	return &proto.GetUserInfoResponse{
		User: &proto.UserInfo{
			ID:           int32(user.ID),
			Username:     user.Username,
			FullName:     user.FullName,
			Sex:          user.Sex,
			Age:          int32(user.Age),
			Mobile:       user.Mobile,
			HealthScore:  int32(user.HealthScore),
			Email:        user.Email,
			Subscription: int32(user.Subscription),
			Status:       int32(user.Status),
		},
	}, nil
}

func (u User) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	//TODO 修改用户信息
	user, err := mysql.GetUserById(int(request.ID))
	if err != nil {
		return nil, errors.New("用户信息获取失败")
	}
	user = &mysql.User{
		ID:           user.ID,
		Username:     request.Username,
		Password:     user.Password,
		FullName:     request.FullName,
		Sex:          request.Sex,
		Age:          int(request.Age),
		Mobile:       request.Mobile,
		HealthScore:  int(request.HealthScore),
		Email:        request.Email,
		Subscription: int8(request.Subscription),
		Status:       int8(request.Status),
		Model: gorm.Model{
			CreatedAt: user.CreatedAt,
		},
	}
	err = mysql.UpdateUser(user)
	if err != nil {
		return nil, errors.New("用户信息修改失败")
	}
	return &proto.UpdateUserResponse{
		User: &proto.UserInfo{
			ID:           int32(user.ID),
			Username:     user.Username,
			FullName:     user.FullName,
			Sex:          user.Sex,
			Age:          int32(user.Age),
			Mobile:       user.Mobile,
			HealthScore:  int32(user.HealthScore),
			Email:        user.Email,
			Subscription: int32(user.Subscription),
			Status:       int32(user.Status),
		},
	}, nil
}

func (u User) DeleteUser(ctx context.Context, request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	//TODO 注销用户
	err := mysql.DeleteUser(int(request.UserId))
	if err != nil {
		return nil, err
	}
	return &proto.DeleteUserResponse{
		IsOk: true,
	}, nil
}

func (u User) UpdatePassword(ctx context.Context, request *proto.UpdatePasswordRequest) (*proto.UpdatePasswordResponse, error) {
	//TODO 修改用户密码
	user, err := mysql.GetUserByMobile(request.Mobile)
	if err != nil {
		return nil, errors.New("未找到该用户")
	}
	user.Password = aes256.Encrypt(request.Password, etc.RpcConfig.Aes.Passport)
	err = mysql.UpdateUser(user)
	if err != nil {
		return nil, errors.New("密码修改失败")
	}
	return &proto.UpdatePasswordResponse{
		IsOk: true,
	}, nil
}

func (u User) GetHealthList(ctx context.Context, request *proto.GetHealthListRequest) (*proto.GetHealthListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	//TODO 用户注册
	user := mysql.User{
		Username: request.Username,
		Password: aes256.Encrypt(request.Password, etc.RpcConfig.Aes.Passport),
		FullName: request.FullName,
		Sex:      request.Sex,
		Age:      int(request.Age),
		Mobile:   request.Mobiles,
		Email:    request.Email,
	}
	// todo:判断用户是否已经存在
	users, _ := mysql.GetUserByUsername(request.Username)
	if users.ID != 0 {
		//zap.S().Info(err)
		return nil, errors.New("用户已经注册,无需重复注册")
	}
	err := mysql.AddUser(&user)
	if err != nil {
		zap.S().Info(err)
		return nil, errors.New("注册失败")
	}
	return &proto.RegisterResponse{
		UserId: int32(user.ID),
	}, nil
}

func (u User) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	//TODO 用户登录
	user, err := mysql.GetUserByUsername(request.Username)
	if err != nil {
		zap.S().Info(err)
		return nil, errors.New("用户信息获取失败")
	}
	if request.Password != aes256.Decrypt(user.Password, etc.RpcConfig.Aes.Passport) {
		zap.S().Info(err)
		return nil, errors.New("密码错误")
	}
	userinfo := &proto.UserInfo{
		ID:           int32(user.ID),
		Username:     user.Username,
		FullName:     user.FullName,
		Sex:          user.Sex,
		Age:          int32(user.Age),
		Mobile:       user.Mobile,
		HealthScore:  int32(user.HealthScore),
		Email:        user.Email,
		Subscription: int32(user.Subscription),
		Status:       int32(user.Status),
	}
	return &proto.LoginResponse{
		User: userinfo,
	}, nil
}
