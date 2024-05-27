package service

/*
	info(
	desc: "预约挂号逻辑"
	author: "济源尚鑫平"
	email: "As1433223@163.com"
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

func (c *Registration) CreateAppointment(ctx context.Context, in *proto.CreateAppointmentReq) (*proto.Empty, error) {
	parsedTime, err := time.Parse(time.DateTime, in.AppointmentData)
	if err != nil {
		fmt.Println("解析时间出错:", err)
		return nil, err
	}
	var (
		tx          = mysql.DB.Begin()
		appointment = &mysql.Appointment{
			UserID:          int(in.UserId),
			AppointmentType: int(in.AppointmentType),
			AppointmentData: parsedTime.Format(time.DateOnly),
			AppointmentTime: parsedTime.Format(time.TimeOnly),
			Status:          int(in.Status),
		}
	)

	//添加预约
	if err = mysql.CreateAppointment(tx, appointment); err != nil {
		tx.Rollback()
		return &proto.Empty{}, err
	}
	//添加挂号表
	if err = mysql.CreateAttendingPhysician(tx, &mysql.AttendingPhysician{
		UserID:        int(in.UserId),
		AppointmentId: int(appointment.ID),
		DoctorID:      int(in.DoctorId),
		OfficeID:      int(in.OfficeId),
		RealName:      in.RealName,
		IDNumber:      in.IdNumber,
	}); err != nil {
		tx.Rollback()
		return &proto.Empty{}, err
	}
	tx.Commit()
	return &proto.Empty{}, nil
}

// 取消预约
func (c *Registration) CancelAppointmentRegistration(ctx context.Context, in *proto.CancelAppointmentRegistrationReq) (*proto.Empty, error) {
	var (
		tx = mysql.DB.Begin()
	)
	//删除预约表
	if err := mysql.DeleteAppointment(tx, int(in.AppointmentId)); err != nil {
		tx.Rollback()
		return &proto.Empty{}, err
	}
	//删除挂号表
	if err := mysql.DeleteAttendingPhysician(tx, int(in.AppointmentId)); err != nil {
		tx.Rollback()
		return &proto.Empty{}, err
	}
	return &proto.Empty{}, nil
}

// 获取预约信息
func (c *Registration) GetAppointmentRegistrationById(ctx context.Context, in *proto.GetAppointmentRegistrationByIdReq) (*proto.GetAppointmentRegistrationByIdRes, error) {
	res, err := mysql.GetAttendingPhysicianById(mysql.DB, int(in.AttendingPhysicianId))
	if err != nil {
		return &proto.GetAppointmentRegistrationByIdRes{}, err
	}
	return &proto.GetAppointmentRegistrationByIdRes{
		Data: &proto.AttendingPhysician{
			Id:            int32(res.ID),
			UserId:        int32(res.UserID),
			AppointmentId: int32(res.AppointmentId),
			DoctorID:      int32(res.DoctorID),
			OfficeID:      int32(res.OfficeID),
			RealName:      res.RealName,
			IDNumber:      res.IDNumber,
			Symptoms:      res.Symptoms,
			Diagnosis:     res.Diagnosis,
			Prescription:  res.Prescription,
		},
	}, nil
}
