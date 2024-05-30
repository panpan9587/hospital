package mdeical

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mdeical"
	mdeicalReq "github.com/flipped-aurora/gin-vue-admin/server/model/mdeical/request"
)

type MedicalService struct {
}

// CreateMedical 创建病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (medicalService *MedicalService) CreateMedical(medical *mdeical.Medical) (err error) {
	err = global.GVA_DB.Create(medical).Error
	return err
}

// DeleteMedical 删除病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (medicalService *MedicalService) DeleteMedical(ID string) (err error) {
	err = global.GVA_DB.Delete(&mdeical.Medical{}, "id = ?", ID).Error
	return err
}

// DeleteMedicalByIds 批量删除病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (medicalService *MedicalService) DeleteMedicalByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]mdeical.Medical{}, "id in ?", IDs).Error
	return err
}

// UpdateMedical 更新病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (medicalService *MedicalService) UpdateMedical(medical mdeical.Medical) (err error) {
	err = global.GVA_DB.Model(&mdeical.Medical{}).Where("id = ?", medical.ID).Updates(&medical).Error
	return err
}

// GetMedical 根据ID获取病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (medicalService *MedicalService) GetMedical(ID string) (medical mdeical.Medical, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&medical).Error
	return
}

// GetMedicalInfoList 分页获取病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (medicalService *MedicalService) GetMedicalInfoList(info mdeicalReq.MedicalSearch) (list []mdeical.Medical, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mdeical.Medical{})
	var medicals []mdeical.Medical
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.IdCard != "" {
		db = db.Where("id_card LIKE ?", "%"+info.IdCard+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&medicals).Error
	return medicals, total, err
}
