package registermgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/registermgmt"
	registermgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/registermgmt/request"
)

type AttendingphysicianService struct {
}

// CreateAttendingphysician 创建挂号管理表记录
// Author [piexlmax](https://github.com/piexlmax)
func (attendingphysicianService *AttendingphysicianService) CreateAttendingphysician(attendingphysician *registermgmt.Attendingphysician) (err error) {
	err = global.GVA_DB.Create(attendingphysician).Error
	return err
}

// DeleteAttendingphysician 删除挂号管理表记录
// Author [piexlmax](https://github.com/piexlmax)
func (attendingphysicianService *AttendingphysicianService) DeleteAttendingphysician(ID string) (err error) {
	err = global.GVA_DB.Delete(&registermgmt.Attendingphysician{}, "id = ?", ID).Error
	return err
}

// DeleteAttendingphysicianByIds 批量删除挂号管理表记录
// Author [piexlmax](https://github.com/piexlmax)
func (attendingphysicianService *AttendingphysicianService) DeleteAttendingphysicianByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]registermgmt.Attendingphysician{}, "id in ?", IDs).Error
	return err
}

// UpdateAttendingphysician 更新挂号管理表记录
// Author [piexlmax](https://github.com/piexlmax)
func (attendingphysicianService *AttendingphysicianService) UpdateAttendingphysician(attendingphysician registermgmt.Attendingphysician) (err error) {
	err = global.GVA_DB.Model(&registermgmt.Attendingphysician{}).Where("id = ?", attendingphysician.ID).Updates(&attendingphysician).Error
	return err
}

// GetAttendingphysician 根据ID获取挂号管理表记录
// Author [piexlmax](https://github.com/piexlmax)
func (attendingphysicianService *AttendingphysicianService) GetAttendingphysician(ID string) (attendingphysician registermgmt.Attendingphysician, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&attendingphysician).Error
	return
}

// GetAttendingphysicianInfoList 分页获取挂号管理表记录
// Author [piexlmax](https://github.com/piexlmax)
func (attendingphysicianService *AttendingphysicianService) GetAttendingphysicianInfoList(info registermgmtReq.AttendingphysicianSearch) (list []registermgmt.Attendingphysician, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&registermgmt.Attendingphysician{})
	var attendingphysicians []registermgmt.Attendingphysician
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

	err = db.Find(&attendingphysicians).Error
	return attendingphysicians, total, err
}
func (attendingphysicianService *AttendingphysicianService) GetAttendingphysicianDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	doctorId := make([]map[string]any, 0)
	global.GVA_DB.Table("doctor").Select("doctor_name as label,id as value").Scan(&doctorId)
	res["doctorId"] = doctorId
	officeId := make([]map[string]any, 0)
	global.GVA_DB.Table("doctor_office").Select("office_name as label, id as value").Scan(&officeId)
	res["officeId"] = officeId
	shiftId := make([]map[string]any, 0)
	global.GVA_DB.Table("shift_time").Select("title as label,shift_id as value").Scan(&shiftId)
	res["shiftId"] = shiftId
	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("user").Select("full_name as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}
