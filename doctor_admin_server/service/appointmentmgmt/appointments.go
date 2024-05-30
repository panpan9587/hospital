package appointmentmgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/appointmentmgmt"
	appointmentmgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/appointmentmgmt/request"
)

type AppointmentService struct {
}

// CreateAppointment 创建预约记录表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appointmentService *AppointmentService) CreateAppointment(appointment *appointmentmgmt.Appointment) (err error) {
	err = global.GVA_DB.Create(appointment).Error
	return err
}

// DeleteAppointment 删除预约记录表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appointmentService *AppointmentService) DeleteAppointment(ID string) (err error) {
	err = global.GVA_DB.Delete(&appointmentmgmt.Appointment{}, "id = ?", ID).Error
	return err
}

// DeleteAppointmentByIds 批量删除预约记录表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appointmentService *AppointmentService) DeleteAppointmentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]appointmentmgmt.Appointment{}, "id in ?", IDs).Error
	return err
}

// UpdateAppointment 更新预约记录表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appointmentService *AppointmentService) UpdateAppointment(appointment appointmentmgmt.Appointment) (err error) {
	err = global.GVA_DB.Model(&appointmentmgmt.Appointment{}).Where("id = ?", appointment.ID).Updates(&appointment).Error
	return err
}

// GetAppointment 根据ID获取预约记录表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appointmentService *AppointmentService) GetAppointment(ID string) (appointment appointmentmgmt.Appointment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&appointment).Error
	return
}

// GetAppointmentInfoList 分页获取预约记录表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appointmentService *AppointmentService) GetAppointmentInfoList(info appointmentmgmtReq.AppointmentSearch) (list []appointmentmgmt.Appointment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&appointmentmgmt.Appointment{})
	var appointments []appointmentmgmt.Appointment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&appointments).Error
	return appointments, total, err
}
func (appointmentService *AppointmentService) GetAppointmentDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	appointmentType := make([]map[string]any, 0)
	global.GVA_DB.Table("appointment_type").Select("appointment_type_title as label,id as value").Scan(&appointmentType)
	res["appointmentType"] = appointmentType
	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("user").Select("full_name as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}
