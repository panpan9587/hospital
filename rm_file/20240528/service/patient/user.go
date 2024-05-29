package patient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/patient"
	patientReq "github.com/flipped-aurora/gin-vue-admin/server/model/patient/request"
	"gorm.io/gorm"
)

type UserService struct {
}

// CreateUser 创建用户表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user *patient.User) (err error) {
	err = global.GVA_DB.Create(user).Error
	return err
}

// DeleteUser 删除用户表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUser(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patient.User{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&patient.User{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteUserByIds 批量删除用户表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUserByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patient.User{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&patient.User{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateUser 更新用户表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) UpdateUser(user patient.User) (err error) {
	err = global.GVA_DB.Model(&patient.User{}).Where("id = ?", user.ID).Updates(&user).Error
	return err
}

// GetUser 根据ID获取用户表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUser(ID string) (user patient.User, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&user).Error
	return
}

// GetUserInfoList 分页获取用户表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUserInfoList(info patientReq.UserSearch) (list []patient.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&patient.User{})
	var users []patient.User
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

	err = db.Find(&users).Error
	return users, total, err
}
