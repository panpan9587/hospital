package userauth

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userauth"
	userauthReq "github.com/flipped-aurora/gin-vue-admin/server/model/userauth/request"
)

type UserAuthService struct {
}

// CreateUserAuth 创建userAuth表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAuthService *UserAuthService) CreateUserAuth(userAuth *userauth.UserAuth) (err error) {
	err = global.GVA_DB.Create(userAuth).Error
	return err
}

// DeleteUserAuth 删除userAuth表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAuthService *UserAuthService) DeleteUserAuth(ID string) (err error) {
	err = global.GVA_DB.Delete(&userauth.UserAuth{}, "id = ?", ID).Error
	return err
}

// DeleteUserAuthByIds 批量删除userAuth表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAuthService *UserAuthService) DeleteUserAuthByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]userauth.UserAuth{}, "id in ?", IDs).Error
	return err
}

// UpdateUserAuth 更新userAuth表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAuthService *UserAuthService) UpdateUserAuth(userAuth userauth.UserAuth) (err error) {
	err = global.GVA_DB.Model(&userauth.UserAuth{}).Where("id = ?", userAuth.ID).Updates(&userAuth).Error
	return err
}

// GetUserAuth 根据ID获取userAuth表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAuthService *UserAuthService) GetUserAuth(ID string) (userAuth userauth.UserAuth, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&userAuth).Error
	return
}

// GetUserAuthInfoList 分页获取userAuth表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAuthService *UserAuthService) GetUserAuthInfoList(info userauthReq.UserAuthSearch) (list []userauth.UserAuth, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&userauth.UserAuth{})
	var userAuths []userauth.UserAuth
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

	err = db.Find(&userAuths).Error
	return userAuths, total, err
}
func (userAuthService *UserAuthService) GetUserAuthDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("user").Select("full_name as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}
