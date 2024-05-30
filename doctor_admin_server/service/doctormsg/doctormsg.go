package doctormsg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/doctormsg"
	doctormsgReq "github.com/flipped-aurora/gin-vue-admin/server/model/doctormsg/request"
)

type DoctorService struct {
}

// CreateDoctor 创建医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) CreateDoctor(doctor *doctormsg.Doctor) (err error) {
	err = global.GVA_DB.Create(doctor).Error
	return err
}

// DeleteDoctor 删除医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) DeleteDoctor(ID string) (err error) {
	err = global.GVA_DB.Delete(&doctormsg.Doctor{}, "id = ?", ID).Error
	return err
}

// DeleteDoctorByIds 批量删除医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) DeleteDoctorByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]doctormsg.Doctor{}, "id in ?", IDs).Error
	return err
}

// UpdateDoctor 更新医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) UpdateDoctor(doctor doctormsg.Doctor) (err error) {
	err = global.GVA_DB.Model(&doctormsg.Doctor{}).Where("id = ?", doctor.ID).Updates(&doctor).Error
	return err
}

// GetDoctor 根据ID获取医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) GetDoctor(ID string) (doctor doctormsg.Doctor, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&doctor).Error
	return
}

// GetDoctorInfoList 分页获取医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) GetDoctorInfoList(info doctormsgReq.DoctorSearch) (list []doctormsg.Doctor, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&doctormsg.Doctor{})
	var doctors []doctormsg.Doctor
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

	err = db.Find(&doctors).Error
	return doctors, total, err
}
func (doctorService *DoctorService) GetDoctorDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	officeId := make([]map[string]any, 0)
	global.GVA_DB.Table("doctor_office").Select("office_name as label,id as value").Scan(&officeId)
	res["officeId"] = officeId
	return
}
