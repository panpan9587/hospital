// 自动生成模板Attendingphysician
package registermgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 挂号管理表 结构体  Attendingphysician
type Attendingphysician struct {
	global.GVA_MODEL
	UserId        *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:19;"`                      //用户名称
	AppointmentId *int   `json:"appointmentId" form:"appointmentId" gorm:"column:appointment_id;comment:;size:19;"` //appointmentId字段
	DoctorId      *int   `json:"doctorId" form:"doctorId" gorm:"column:doctor_id;comment:;size:19;"`                //doctorId字段
	OfficeId      *int   `json:"officeId" form:"officeId" gorm:"column:office_id;comment:;size:19;"`                //officeId字段
	ShiftId       *int   `json:"shiftId" form:"shiftId" gorm:"column:shift_id;comment:;size:19;"`                   //shiftId字段
	RealName      string `json:"realName" form:"realName" gorm:"column:real_name;comment:;size:50;"`                //realName字段
	Mobile        string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:;size:20;"`                       //mobile字段
	IdNumber      string `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:;size:191;"`               //idNumber字段
	Symptoms      string `json:"symptoms" form:"symptoms" gorm:"column:symptoms;comment:;size:191;"`                //symptoms字段
	Diagnosis     string `json:"diagnosis" form:"diagnosis" gorm:"column:diagnosis;comment:;size:191;"`             //diagnosis字段
	Prescription  string `json:"prescription" form:"prescription" gorm:"column:prescription;comment:;size:191;"`    //prescription字段
}

// TableName 挂号管理表 Attendingphysician自定义表名 attendingphysician
func (Attendingphysician) TableName() string {
	return "attendingphysician"
}
