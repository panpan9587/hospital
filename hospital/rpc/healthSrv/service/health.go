package service

import (
	"context"
	proto "demo/rpc/healthSrv/healthclient"
	"demo/rpc/model/mysql"
)

type HealthService struct {
	proto.UnimplementedHealthServiceServer
}

// 预约记录
func (h *HealthService) GetAppointment(ctx context.Context, in *proto.GetAppointmentReq) (*proto.GetAppointmentResp, error) {
	var app mysql.Appointment
	app.UserID = int(in.UserID)
	app.Mobile = in.Mobile
	app.AppointmentData = in.AppointmentData
	app.AppointmentTime = in.AppointmentTime
	app.AppointmentType = int(in.AppointmentType)
	app.Status = int(in.Status)
	err := mysql.DB.Table("appointment").Create(&app).Error
	if err != nil {
		return nil, err
	}
	info := &proto.AppointmentInfo{
		ID:              int64(app.ID),
		UserID:          int64(app.UserID),
		AppointmentType: int64(app.AppointmentType),
		Mobile:          app.Mobile,
		AppointmentData: app.AppointmentData,
		AppointmentTime: app.AppointmentTime,
		Status:          int64(app.Status),
	}
	return &proto.GetAppointmentResp{Appointment: info}, nil
}

// GetHealth 体检表记录
func (h *HealthService) GetHealth(ctx context.Context, in *proto.GetHealthReq) (*proto.GetHealthResp, error) {
	var health mysql.Health
	var project mysql.HealthInfo
	health.AppointmentId = in.AppointmentId
	health.UserID = in.UserID
	health.RealName = in.RealName
	health.IdNumber = in.IdNumber
	t := mysql.DB.Begin()
	err := mysql.DB.Table("health").Create(&health).Error
	if err != nil {
		t.Rollback()
		return nil, err
	}
	project.HealthID = in.ProjectInfo.HealthId
	project.UserID = in.ProjectInfo.UserID
	project.DoctorId = in.ProjectInfo.DoctorId
	project.Height = float64(in.ProjectInfo.Height)
	project.Weight = float64(in.ProjectInfo.Weight)
	project.HeartRate = in.ProjectInfo.HeartRate
	project.Hearing = in.ProjectInfo.Hearing
	project.BloodPressur = in.ProjectInfo.BloodPressur
	project.BloodSugar = in.ProjectInfo.BloodSugar
	project.Urine = in.ProjectInfo.Urine
	project.Ct = in.ProjectInfo.Ct

	err = t.Table("health_info").Create(&project).Error
	if err != nil {
		t.Rollback()
		return nil, err
	}
	Info := &proto.HealthInfo{
		ID:            health.ID,
		AppointmentId: health.AppointmentId,
		UserID:        health.UserID,
		RealName:      health.RealName,
		IdNumber:      health.IdNumber,
		ProjectInfo: &proto.HealthProjectInfo{
			ID:           project.ID,
			HealthId:     project.HealthID,
			UserID:       project.UserID,
			DoctorId:     project.DoctorId,
			Height:       uint64(project.Height),
			Weight:       uint64(project.Weight),
			HeartRate:    project.HeartRate,
			Hearing:      project.Hearing,
			BloodPressur: project.BloodPressur,
			BloodSugar:   project.BloodSugar,
			Urine:        project.Urine,
			Ct:           project.Ct,
		},
	}
	t.Commit()
	return &proto.GetHealthResp{Health: Info}, nil
}

// GetHealthId 根据userid查询体检详情列表
func (h *HealthService) GetHealthId(ctx context.Context, in *proto.GetHealthIdReq) (*proto.GetHealthIdResp, error) {
	var health []mysql.Health
	err := mysql.DB.Table("health").Where("appointment_id=?", in.AppointmentId).Find(&health).Error
	if err != nil {
		return nil, err
	}
	var rsp []*proto.HealthInfo
	for _, info := range health {
		tmp := &proto.HealthInfo{
			ID:            info.ID,
			AppointmentId: info.AppointmentId,
			UserID:        info.UserID,
			RealName:      info.RealName,
			IdNumber:      string(info.IdNumber),
		}
		rsp = append(rsp, tmp)
	}
	return &proto.GetHealthIdResp{Health: rsp}, nil
}

// HealthProjectId 根据user_id查询体检项目详情
func (h *HealthService) HealthProjectId(ctx context.Context, in *proto.HealthProjectIdReq) (*proto.HealthProjectIdResp, error) {
	var projectId mysql.HealthInfo
	err := mysql.DB.Table("health_info").Where("user_id=?", in.UserID).First(&projectId).Error
	if err != nil {
		return nil, err
	}
	Info := &proto.HealthProjectInfo{
		ID:           projectId.ID,
		HealthId:     projectId.HealthID,
		UserID:       projectId.UserID,
		DoctorId:     projectId.DoctorId,
		Height:       uint64(projectId.Height),
		Weight:       uint64(projectId.Weight),
		HeartRate:    projectId.HeartRate,
		Hearing:      projectId.Hearing,
		BloodPressur: projectId.BloodPressur,
		BloodSugar:   projectId.BloodSugar,
		Urine:        projectId.Urine,
		Ct:           projectId.Ct,
	}
	return &proto.HealthProjectIdResp{ProjectInfo: Info}, nil
}

// GetDoctorOffice 科室详情
func (h *HealthService) GetDoctorOffice(ctx context.Context, in *proto.GetDoctorOfficeReq) (*proto.GetDoctorOfficeResp, error) {
	var office mysql.DoctorOffice
	err := mysql.DB.Table("doctor_office").Where("id=?", in.ID).First(&office).Error
	if err != nil {
		return nil, err
	}
	return &proto.GetDoctorOfficeResp{DoctorOffice: &proto.DoctorOfficeInfo{
		ID:         office.Id,
		OfficeName: office.OfficeName,
		Status:     int64(office.Status),
	}}, nil
}

// GetPackage  套餐详情
func (h *HealthService) GetPackage(ctx context.Context, in *proto.GetPackageReq) (*proto.GetPackageResp, error) {
	var pack mysql.Package
	var app mysql.Appointment
	err := mysql.DB.Preload("app").Where("id=?", in.ID).First(&pack).Error
	if err != nil {
		return nil, err
	}
	info := &proto.GetPackageResp{
		PackageName:   pack.PackageName,
		PackageAmount: pack.PackageAmount,
		CheckupDate: &proto.CheckupDate{
			AppointmentData: app.AppointmentData,
			AppointmentTime: app.AppointmentTime,
		},
	}
	return info, nil
}
