package doctorment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/doctorment"
	doctormentReq "github.com/flipped-aurora/gin-vue-admin/server/model/doctorment/request"
)

type ShiftService struct {
}

// CreateShift 创建排班表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftService *ShiftService) CreateShift(shift *doctorment.Shift) (err error) {
	err = global.GVA_DB.Create(shift).Error
	return err
}

// DeleteShift 删除排班表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftService *ShiftService) DeleteShift(ID string) (err error) {
	err = global.GVA_DB.Delete(&doctorment.Shift{}, "id = ?", ID).Error
	return err
}

// DeleteShiftByIds 批量删除排班表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftService *ShiftService) DeleteShiftByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]doctorment.Shift{}, "id in ?", IDs).Error
	return err
}

// UpdateShift 更新排班表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftService *ShiftService) UpdateShift(shift doctorment.Shift) (err error) {
	err = global.GVA_DB.Model(&doctorment.Shift{}).Where("id = ?", shift.ID).Updates(&shift).Error
	return err
}

// GetShift 根据ID获取排班表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftService *ShiftService) GetShift(ID string) (shift doctorment.Shift, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&shift).Error
	return
}

// GetShiftInfoList 分页获取排班表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftService *ShiftService) GetShiftInfoList(info doctormentReq.ShiftSearch) (list []doctorment.Shift, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&doctorment.Shift{})
	var shifts []doctorment.Shift
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

	err = db.Find(&shifts).Error
	return shifts, total, err
}
func (shiftService *ShiftService) GetShiftDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	doctorId := make([]map[string]any, 0)
	global.GVA_DB.Table("doctor").Select("doctor_name as label,id as value").Scan(&doctorId)
	res["doctorId"] = doctorId
	officeId := make([]map[string]any, 0)
	global.GVA_DB.Table("doctor_office").Select("office_name as label,id as value").Scan(&officeId)
	res["officeId"] = officeId
	return
}
