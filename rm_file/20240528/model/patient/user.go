// 自动生成模板User
package patient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户表 结构体  User
type User struct {
	global.GVA_MODEL
	Username     string `json:"username" form:"username" gorm:"column:username;comment:用户名（手机号或邮箱）;size:18;" binding:"required"` //用户名（手机号或邮箱）
	Password     string `json:"password" form:"password" gorm:"column:password;comment:密码;size:64;" binding:"required"`          //密码
	FullName     string `json:"fullName" form:"fullName" gorm:"column:full_name;comment:姓名;size:50;"`                            //姓名
	Sex          string `json:"sex" form:"sex" gorm:"column:sex;comment:性别 男 女;size:2;" binding:"required"`                      //性别 男 女
	Age          *int   `json:"age" form:"age" gorm:"column:age;comment:年龄;size:10;" binding:"required"`                         //年龄
	Mobile       string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;size:20;" binding:"required"`               //手机号
	HealthScore  *int   `json:"healthScore" form:"healthScore" gorm:"column:health_score;comment:用户健康分,用于活动抽奖;size:10;"`         //用户健康分,用于活动抽奖
	Email        string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:50;"`                                      //邮箱
	Subscription *int   `json:"subscription" form:"subscription" gorm:"column:subscription;comment:是否订阅健康资讯 0:否 1:是;size:1;"`    //是否订阅健康资讯 0:否 1:是
	Status       *int   `json:"status" form:"status" gorm:"column:status;comment:用户状态，3，表示已经实名;size:1;"`                         //用户状态，3，表示已经实名
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 用户表 User自定义表名 user
func (User) TableName() string {
	return "user"
}
