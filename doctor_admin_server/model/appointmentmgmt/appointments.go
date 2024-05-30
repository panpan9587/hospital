// 自动生成模板Appointment
package appointmentmgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 预约记录表 结构体  Appointment
type Appointment struct {
	global.GVA_MODEL
	UserId          *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:19;"`                    //用户id
	AppointmentType *int       `json:"appointmentType" form:"appointmentType" gorm:"column:appointment_type;comment:预约类型;"` //预约类型
	AppointmentData *time.Time `json:"appointmentData" form:"appointmentData" gorm:"column:appointment_data;comment:预约时间;"` //预约时间
	AppointmentTime *time.Time `json:"appointmentTime" form:"appointmentTime" gorm:"column:appointment_time;comment:预约时间;"` //预约时间
	Status          *bool      `json:"status" form:"status" gorm:"column:status;comment:状态;"`                               //状态
	PackageId       *int       `json:"packageId" form:"packageId" gorm:"column:package_id;comment:上级id;size:19;"`           //上级id
}

// TableName 预约记录表 Appointment自定义表名 appointment
func (Appointment) TableName() string {
	return "appointment"
}
