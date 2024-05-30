package doctormgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/doctormgmt"
	doctormgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/doctormgmt/request"
)

type DoctorService struct {
}

// CreateDoctor 创建医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) CreateDoctor(doctor *doctormgmt.Doctor) (err error) {
	err = global.GVA_DB.Create(doctor).Error
	return err
}

// DeleteDoctor 删除医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) DeleteDoctor(ID string) (err error) {
	err = global.GVA_DB.Delete(&doctormgmt.Doctor{}, "id = ?", ID).Error
	return err
}

// DeleteDoctorByIds 批量删除医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) DeleteDoctorByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]doctormgmt.Doctor{}, "id in ?", IDs).Error
	return err
}

// UpdateDoctor 更新医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) UpdateDoctor(doctor doctormgmt.Doctor) (err error) {
	err = global.GVA_DB.Model(&doctormgmt.Doctor{}).Where("id = ?", doctor.ID).Updates(&doctor).Error
	return err
}

// GetDoctor 根据ID获取医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) GetDoctor(ID string) (doctor doctormgmt.Doctor, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&doctor).Error
	return
}

// GetDoctorInfoList 分页获取医生表记录
// Author [piexlmax](https://github.com/piexlmax)
func (doctorService *DoctorService) GetDoctorInfoList(info doctormgmtReq.DoctorSearch) (list []doctormgmt.Doctor, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&doctormgmt.Doctor{})
	var doctors []doctormgmt.Doctor
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
