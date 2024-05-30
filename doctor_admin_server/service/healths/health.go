package healths

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/healths"
	healthsReq "github.com/flipped-aurora/gin-vue-admin/server/model/healths/request"
)

type HealthService struct {
}

// CreateHealth 创建体检表记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthService *HealthService) CreateHealth(health *healths.Health) (err error) {
	err = global.GVA_DB.Create(health).Error
	return err
}

// DeleteHealth 删除体检表记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthService *HealthService) DeleteHealth(ID string) (err error) {
	err = global.GVA_DB.Delete(&healths.Health{}, "id = ?", ID).Error
	return err
}

// DeleteHealthByIds 批量删除体检表记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthService *HealthService) DeleteHealthByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]healths.Health{}, "id in ?", IDs).Error
	return err
}

// UpdateHealth 更新体检表记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthService *HealthService) UpdateHealth(health healths.Health) (err error) {
	err = global.GVA_DB.Model(&healths.Health{}).Where("id = ?", health.ID).Updates(&health).Error
	return err
}

// GetHealth 根据ID获取体检表记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthService *HealthService) GetHealth(ID string) (health healths.Health, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&health).Error
	return
}

// GetHealthInfoList 分页获取体检表记录
// Author [piexlmax](https://github.com/piexlmax)
func (healthService *HealthService) GetHealthInfoList(info healthsReq.HealthSearch) (list []healths.Health, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&healths.Health{})
	var healths []healths.Health
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

	err = db.Find(&healths).Error
	return healths, total, err
}
