package disease

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/disease"
	diseaseReq "github.com/flipped-aurora/gin-vue-admin/server/model/disease/request"
)

type DiseasesService struct {
}

// CreateDiseases 创建病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (diseasesService *DiseasesService) CreateDiseases(diseases *disease.Diseases) (err error) {
	err = global.GVA_DB.Create(diseases).Error
	return err
}

// DeleteDiseases 删除病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (diseasesService *DiseasesService) DeleteDiseases(ID string) (err error) {
	err = global.GVA_DB.Delete(&disease.Diseases{}, "id = ?", ID).Error
	return err
}

// DeleteDiseasesByIds 批量删除病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (diseasesService *DiseasesService) DeleteDiseasesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]disease.Diseases{}, "id in ?", IDs).Error
	return err
}

// UpdateDiseases 更新病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (diseasesService *DiseasesService) UpdateDiseases(diseases disease.Diseases) (err error) {
	err = global.GVA_DB.Model(&disease.Diseases{}).Where("id = ?", diseases.ID).Updates(&diseases).Error
	return err
}

// GetDiseases 根据ID获取病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (diseasesService *DiseasesService) GetDiseases(ID string) (diseases disease.Diseases, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&diseases).Error
	return
}

// GetDiseasesInfoList 分页获取病历表记录
// Author [piexlmax](https://github.com/piexlmax)
func (diseasesService *DiseasesService) GetDiseasesInfoList(info diseaseReq.DiseasesSearch) (list []disease.Diseases, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&disease.Diseases{})
	var diseasess []disease.Diseases
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.IdNumber != "" {
		db = db.Where("id_number LIKE ?", "%"+info.IdNumber+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&diseasess).Error
	return diseasess, total, err
}
