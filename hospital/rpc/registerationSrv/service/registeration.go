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
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Registration struct {
	proto.UnimplementedRegisterationServer
}

// 预约挂号
func (c *Registration) CreateAppointment(ctx context.Context, in *proto.CreateAppointmentReq) (*proto.Empty, error) {
	//判断是否有票  （前后端双重判断）
	shiftRes, err := mysql.GetShiftDoctorById(int(in.ShiftId))
	if err != nil {
		return &proto.Empty{}, err
	}
	if shiftRes.Count >= 50 && shiftRes.Status == 2 {
		return &proto.Empty{}, errors.New("没票了")
	}
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
		Mobile:        in.Mobile,
		ShiftId:       int(in.ShiftId),
		Model:         gorm.Model{},
	}); err != nil {
		tx.Rollback()
		return &proto.Empty{}, err
	}
	//修改预约票数
	shiftRes.Count += 1
	if shiftRes.Count >= 50 {
		shiftRes.Status = 2
	}
	err = mysql.UpdateShift(tx, shiftRes)
	if err != nil {
		return &proto.Empty{}, err
	}
	tx.Commit()
	return &proto.Empty{}, nil
}

// 取消预约
func (c *Registration) CancelAppointmentRegistration(ctx context.Context, in *proto.CancelAppointmentRegistrationReq) (*proto.Empty, error) {
	var (
		tx         = mysql.DB.Begin()
		wg         sync.WaitGroup
		resultChan = make(chan bool, 3)
		errRes     string
	)
	defer close(resultChan)
	wg.Add(3)
	go func() {
		defer wg.Done()
		//删除预约表
		if err := mysql.DeleteAppointment(tx, int(in.AppointmentId)); err != nil {
			resultChan <- false
			errRes += err.Error() + ";"
			return
		}
		resultChan <- true
	}()
	go func() {
		wg.Done()
		//删除挂号表
		if err := mysql.DeleteAttendingPhysician(tx, int(in.AppointmentId)); err != nil {
			resultChan <- false
			errRes += err.Error() + ";"
			return
		}
		errRes += "sssssss" + ";"
		resultChan <- true
	}()
	go func() {
		wg.Done()
		//归还预约票数
		shiftRes, err := mysql.GetShiftDoctorById(int(in.ShiftId))
		if err != nil {
			resultChan <- false
			errRes += err.Error() + ";"
			return
		}
		if shiftRes.Count >= 50 && shiftRes.Status == 2 {
			shiftRes.Status = 1
		}
		shiftRes.Count -= 1
		err = mysql.UpdateShift(tx, shiftRes)
		if err != nil {
			resultChan <- false
			errRes += err.Error() + ";"
			return
		}
		resultChan <- false
	}()
	//fmt.Println(<-resultChan, <-resultChan, <-resultChan, "///")
	wg.Wait()
	if !(<-resultChan && <-resultChan && <-resultChan) {
		tx.Rollback()
		return &proto.Empty{}, errors.New(errRes)
	}
	tx.Commit()
	return &proto.Empty{}, nil
}

// 获取预约信息  根据id
func (c *Registration) GetAppointmentById(ctx context.Context, in *proto.GetAppointmentByIdReq) (*proto.GetAppointmentByIdRes, error) {
	res, err := mysql.GetAppointmentById(mysql.DB, int(in.AppointmentId))
	if err != nil {
		return &proto.GetAppointmentByIdRes{}, err
	}
	return &proto.GetAppointmentByIdRes{
		Data: &proto.Appointment{
			Id:              int32(res.ID),
			UserId:          int32(res.UserID),
			AppointmentType: int32(res.AppointmentType),
			AppointmentData: res.AppointmentData,
			AppointmentTime: res.AppointmentTime,
			Status:          int32(res.Status),
		},
	}, nil
}

// 获取挂号信息  根据id
func (c *Registration) GetAttendingPhysicianByAppointmentId(ctx context.Context, in *proto.GetAttendingPhysicianByAppointmentIdReq) (*proto.GetAttendingPhysicianByAppointmentIdRes, error) {
	res, err := mysql.GetAttendingPhysicianById(mysql.DB, int(in.AppointmentId))
	if err != nil {
		return &proto.GetAttendingPhysicianByAppointmentIdRes{}, err
	}
	return &proto.GetAttendingPhysicianByAppointmentIdRes{
		Data: &proto.AttendingPhysician{
			Id:            int32(res.ID),
			UserId:        int32(res.UserID),
			AppointmentId: int32(res.AppointmentId),
			DoctorID:      int32(res.DoctorID),
			OfficeID:      int32(res.OfficeID),
			RealName:      res.RealName,
			IDNumber:      res.IDNumber,
			Mobile:        res.Mobile,
			Symptoms:      res.Symptoms,
			Diagnosis:     res.Diagnosis,
			Prescription:  res.Prescription,
			ShiftId:       int32(res.ShiftId),
		},
	}, nil
}

// 根据身份证号查询挂号记录
func (c *Registration) GetAttendingPhysicianByIdNumber(ctx context.Context, in *proto.GetAttendingPhysicianByIdNumberReq) (*proto.GetAttendingPhysicianByIdNumberRes, error) {
	var (
		list []*proto.AttendingPhysician
	)
	res, err := mysql.GetAttendingPhysicianByIdNumber(mysql.DB, in.IdNumber)
	if err != nil {
		return &proto.GetAttendingPhysicianByIdNumberRes{}, err
	}
	for _, v := range res {
		list = append(list, &proto.AttendingPhysician{
			Id:            int32(v.ID),
			UserId:        int32(v.UserID),
			AppointmentId: int32(v.AppointmentId),
			DoctorID:      int32(v.DoctorID),
			OfficeID:      int32(v.OfficeID),
			RealName:      v.RealName,
			IDNumber:      v.IDNumber,
			Mobile:        v.Mobile,
			Symptoms:      v.Symptoms,
			Diagnosis:     v.Diagnosis,
			Prescription:  v.Prescription,
			ShiftId:       int32(v.ShiftId),
		})
	}
	return &proto.GetAttendingPhysicianByIdNumberRes{
		Data: list,
	}, nil
}
