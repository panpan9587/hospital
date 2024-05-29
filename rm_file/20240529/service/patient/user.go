package patient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/patient"
	patientReq "github.com/flipped-aurora/gin-vue-admin/server/model/patient/request"
)

type UserService struct {
}

// CreateUser 创建user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user *patient.User) (err error) {
	err = global.GVA_DB.Create(user).Error
	return err
}

// DeleteUser 删除user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&patient.User{}, "id = ?", ID).Error
	return err
}

// DeleteUserByIds 批量删除user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]patient.User{}, "id in ?", IDs).Error
	return err
}

// UpdateUser 更新user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) UpdateUser(user patient.User) (err error) {
	err = global.GVA_DB.Model(&patient.User{}).Where("id = ?", user.ID).Updates(&user).Error
	return err
}

// GetUser 根据ID获取user表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUser(ID string) (user patient.User, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&user).Error
	return
}

// GetUserInfoList 分页获取user表记录
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
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["age"] = true
	orderMap["health_score"] = true
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
