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
	"demo/rpc/model/mysql"
	proto "demo/rpc/registerationSrv/registerationclient"
	"fmt"
	"time"
)

type Registration struct {
	proto.UnimplementedRegisterationServer
}

func (c *Registration) AddRegisteration(ctx context.Context, in *proto.AddAppointmentRegisterReq) (*proto.Empty, error) {
	parsedTime, err := time.Parse(time.DateTime, in.AppointmentTime)
	if err != nil {
		fmt.Println("解析时间字符串出错:", err)
		return &proto.Empty{}, nil
	}
	registration := &mysql.PatientAppointment{
		PatientID:       int(in.PatientId),
		DoctorID:        int(in.DoctorId),
		RealName:        in.RealName,
		IdNumber:        in.IdNumber,
		Mobile:          in.Mobile,
		AppointmentTime: parsedTime,
	}
	res := mysql.DB.Create(registration)
	if res.Error != nil {
		return &proto.Empty{}, res.Error
	}
	return &proto.Empty{}, nil
}
func (c *Registration) CancelAppointmentRegistration(ctx context.Context, in *proto.CancelAppointmentRegistrationReq) (*proto.Empty, error) {
	res := mysql.DB.Where("id = ?", in.AppointmentId).
		Delete(&mysql.PatientAppointment{})
	if res.Error != nil {
		return &proto.Empty{}, res.Error
	}
	return &proto.Empty{}, nil
}
func (c *Registration) GetAppointmentRegistrationById(ctx context.Context, in *proto.GetAppointmentRegistrationByIdReq) (*proto.GetAppointmentRegistrationByIdRes, error) {
	var (
		appointment = new(mysql.PatientAppointment)
	)
	res := mysql.DB.Where("id = ?", in.AppointmentId).First(appointment)
	if res.Error != nil {
		return &proto.GetAppointmentRegistrationByIdRes{}, res.Error
	}
	return &proto.GetAppointmentRegistrationByIdRes{
		Data: &proto.AppointmentRegistration{
			Id:              int32(appointment.ID),
			PatientId:       int32(appointment.PatientID),
			DoctorId:        int32(appointment.DoctorID),
			RealName:        appointment.RealName,
			IdNumber:        appointment.IdNumber,
			Mobile:          appointment.Mobile,
			AppointmentTime: appointment.AppointmentTime.Format(time.DateTime),
		},
	}, nil
}
func (c *Registration) UpdateAppointmentRegistration(ctx context.Context, in *proto.UpdateAppointmentRegistrationReq) (*proto.Empty, error) {
	var (
		registration *mysql.PatientAppointment
	)
	if in.Data.AppointmentTime != "" {
		parsedTime, err := time.Parse(time.DateTime, in.Data.AppointmentTime)
		if err != nil {
			fmt.Println("解析时间字符串出错:", err)
			return &proto.Empty{}, nil
		}
		registration.AppointmentTime = parsedTime
	}

	if in.Data.PatientId > 0 {
		registration.PatientID = int(in.Data.PatientId)
	}
	if in.Data.DoctorId > 0 {
		registration.DoctorID = int(in.Data.DoctorId)
	}
	registration.Status = int(in.Data.Status)
	res := mysql.DB.Where("id = ?", in.Data.Id).Updates(registration)
	if res.Error != nil {
		return nil, res.Error
	}
	return &proto.Empty{}, nil
}
