// 自动生成模板UserAuth
package userauth

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 实名认证表 结构体  UserAuth
type UserAuth struct {
	global.GVA_MODEL
	IdNumber string `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:身份证号;size:18;" binding:"required"`   //身份证号
	RealName string `json:"realName" form:"realName" gorm:"column:real_name;comment:用户真实姓名;size:30;" binding:"required"` //用户真实姓名
	Status   *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:1;"`                                //状态
	UserId   *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户;size:11;" binding:"required"`           //用户
}

// TableName 实名认证表 UserAuth自定义表名 user_auth
func (UserAuth) TableName() string {
	return "user_auth"
}
