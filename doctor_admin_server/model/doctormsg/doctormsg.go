// 自动生成模板Doctor
package doctormsg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 医生表 结构体  Doctor
type Doctor struct {
	global.GVA_MODEL
	DoctorName string     `json:"doctorName" form:"doctorName" gorm:"column:doctor_name;comment:医生姓名;size:30;" binding:"required"` //医生姓名
	Age        *int       `json:"age" form:"age" gorm:"column:age;comment:医生年龄;size:11;" binding:"required"`                       //医生年龄
	Sex        string     `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:4;" binding:"required"`                          //性别
	Position   string     `json:"position" form:"position" gorm:"column:position;comment:医生职位;size:200;" binding:"required"`       //医生职位
	Tag        string     `json:"tag" form:"tag" gorm:"column:tag;comment:医生标签;size:200;" binding:"required"`                      //医生标签
	Desc       string     `json:"desc" form:"desc" gorm:"column:desc;comment:医生简介;size:191;type:text;" binding:"required"`         //医生简介
	WorkAge    *int       `json:"workAge" form:"workAge" gorm:"column:work_age;comment:医生工龄;size:11;" binding:"required"`          //医生工龄
	WorkTime   *time.Time `json:"workTime" form:"workTime" gorm:"column:work_time;comment:医生工作时间;" binding:"required"`             //医生工作时间
	OfficeId   *int       `json:"officeId" form:"officeId" gorm:"column:office_id;comment:科室id;size:11;" binding:"required"`       //科室
	Status     *int       `json:"status" form:"status" gorm:"column:status;comment:状态(1:在职 2:离职);size:1;" binding:"required"`      //状态
}

// TableName 医生表 Doctor自定义表名 doctor
func (Doctor) TableName() string {
	return "doctor"
}
