package service

import (
	"context"
	proto "demo/rpc/healthSrv/healthclient"
	"demo/rpc/model/mysql"
	"errors"
)

type HealthService struct {
	proto.UnimplementedHealthServiceServer
}

// AddBodyInspect 预约体检信息记录
func (h HealthService) AddBodyInspect(ctx context.Context, request *proto.BodyInspectRequest) (*proto.BodyInspectResponse, error) {
	var body mysql.BodyInspect
	body.PeopleMsgId = request.PeopleMsgId
	body.Height = int(request.Height)
	body.Weight = int(request.Weight)
	body.Inheritance = request.Inheritance
	body.DoctorId = request.DoctorId
	err := mysql.DB.Table("body_inspect").Create(&body).Error
	if err != nil {
		return nil, errors.New("体检信息添加失败")
	}
	Info := &proto.BodyInspectInfo{
		ID:          body.ID,
		PeopleMsgId: body.PeopleMsgId,
		Height:      int64(body.Height),
		Weight:      int64(body.Weight),
		Inheritance: body.Inheritance,
		DoctorId:    body.DoctorId,
	}
	return &proto.BodyInspectResponse{BodyInspect: Info}, nil
}

// GetMedicalItems 体检项目详情
func (h HealthService) GetMedicalItems(ctx context.Context, req *proto.MedicalItemsRequest) (*proto.MedicalItemsResponse, error) {
	var medical mysql.MedicalItems
	err := mysql.DB.Table("medical_items").Where("id = ?", req.Id).First(&medical).Error
	if err != nil {
		return nil, errors.New("查询失败")
	}
	info := &proto.MedicalItemsInfo{
		Id:            medical.ID,
		Visual:        medical.Visual,
		Hearing:       medical.Hearing,
		BloodPressure: medical.BloodPressure,
		BloodSugar:    medical.BloodSugar,
		Urine:         medical.Urine,
		Ct:            medical.Ct,
	}
	return &proto.MedicalItemsResponse{MedicalItems: info}, nil
}

// GetBodyInspect 预约体检信息详情
func (h HealthService) GetBodyInspect(ctx context.Context, req *proto.GetBodyInspectRequest) (*proto.GetBodyInspectResponse, error) {
	var body mysql.BodyInspect
	err := mysql.DB.Table("body_inspect").Where("id =?", req.Id).First(&body).Error
	if err != nil {
		return nil, errors.New("查询失败")
	}
	Info := &proto.BodyInspectInfo{
		ID:          body.ID,
		PeopleMsgId: body.PeopleMsgId,
		Height:      int64(body.Height),
		Weight:      int64(body.Weight),
		Inheritance: body.Inheritance,
		DoctorId:    body.DoctorId,
	}
	return &proto.GetBodyInspectResponse{BodyInspect: Info}, nil
}

func (h HealthService) GetSignIn(ctx context.Context, req *proto.GetSignInRequest) (*proto.GetSignInResponse, error) {
	var signIn mysql.SignIn
	signIn.UserId = int(req.UserId)
	signIn.Status = int(req.Status)
	signIn.SignInMethod = req.SignInMethod
	err := mysql.DB.Table("sign_in").Create(&signIn).Error
	if err != nil {
		return nil, errors.New("签到失败")
	}
	Info := &proto.SignInInfo{
		UserId:       int64(signIn.UserId),
		Status:       int64(signIn.Status),
		SignInMethod: signIn.SignInMethod,
	}
	return &proto.GetSignInResponse{SignIn: Info}, nil
}
