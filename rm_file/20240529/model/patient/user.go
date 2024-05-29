// 自动生成模板User
package patient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// user表 结构体  User
type User struct {
	global.GVA_MODEL
	Age          *int   `json:"age" form:"age" gorm:"column:age;comment:年龄;size:3;" binding:"required"` //年龄
	Email        string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:50;"`             //邮箱
	FullName     string `json:"fullName" form:"fullName" gorm:"column:full_name;comment:姓名;size:50;"`
	HealthScore  *int   `json:"healthScore" form:"healthScore" gorm:"column:health_score;comment:用户健康分,用于活动抽奖;size:11;"`         //用户健康分,用于活动抽奖
	Mobile       string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;size:20;" binding:"required"`               //手机号
	Password     string `json:"password" form:"password" gorm:"column:password;comment:密码;size:64;" binding:"required"`          //密码
	Sex          string `json:"sex" form:"sex" gorm:"column:sex;comment:性别 男 女;size:2;"`                                         //性别 男 女
	Status       *int   `json:"status" form:"status" gorm:"column:status;comment:用户状态，3，表示已经实名;size:1;"`                         //用户状态，3，表示已经实名
	Subscription *int   `json:"subscription" form:"subscription" gorm:"column:subscription;comment:是否订阅健康资讯 0:否 1:是;size:1;"`    //是否订阅健康资讯 0:否 1:是
	Username     string `json:"username" form:"username" gorm:"column:username;comment:用户名（手机号或邮箱）;size:18;" binding:"required"` //用户名（手机号或邮箱）
}

// TableName user表 User自定义表名 user
func (User) TableName() string {
	return "user"
}
