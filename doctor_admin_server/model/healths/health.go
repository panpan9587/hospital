// 自动生成模板Health
package healths

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 体检表 结构体  Health
type Health struct {
	global.GVA_MODEL
	AppointmentId *int   `json:"appointmentId" form:"appointmentId" gorm:"column:appointment_id;comment:预约id;size:10;"` //预约id
	IdNumber      string `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:身份证号;"`                        //身份证号
	RealName      string `json:"realName" form:"realName" gorm:"column:real_name;comment:患者姓名;size:30;"`                //患者姓名
	UserId        *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:10;"`                      //用户id
}

// TableName 体检表 Health自定义表名 health
func (Health) TableName() string {
	return "health"
}
