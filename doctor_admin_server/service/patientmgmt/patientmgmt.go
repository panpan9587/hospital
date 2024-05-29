package patientmgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/patientmgmt"
	patientmgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/patientmgmt/request"
)

type UserService struct {
}

// CreateUser 创建患者表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user *patientmgmt.User) (err error) {
	err = global.GVA_DB.Create(user).Error
	return err
}

// DeleteUser 删除患者表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&patientmgmt.User{}, "id = ?", ID).Error
	return err
}

// DeleteUserByIds 批量删除患者表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]patientmgmt.User{}, "id in ?", IDs).Error
	return err
}

// UpdateUser 更新患者表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) UpdateUser(user patientmgmt.User) (err error) {
	err = global.GVA_DB.Model(&patientmgmt.User{}).Where("id = ?", user.ID).Updates(&user).Error
	return err
}

// GetUser 根据ID获取患者表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUser(ID string) (user patientmgmt.User, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&user).Error
	return
}

// GetUserInfoList 分页获取患者表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUserInfoList(info patientmgmtReq.UserSearch) (list []patientmgmt.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&patientmgmt.User{})
	var users []patientmgmt.User
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StartAge != nil && info.EndAge != nil {
		db = db.Where("age BETWEEN ? AND ? ", info.StartAge, info.EndAge)
	}
	if info.Mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+info.Mobile+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["age"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&users).Error
	return users, total, err
}
