// 自动生成模板Medical
package mdeical

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 病历表 结构体  Medical
type Medical struct {
	global.GVA_MODEL
	Allergies    string     `json:"allergies" form:"allergies" gorm:"column:allergies;comment:过敏史;size:50;"`             //过敏史
	Diagnosis    string     `json:"diagnosis" form:"diagnosis" gorm:"column:diagnosis;comment:诊断;size:100;"`             //诊断
	IdCard       string     `json:"idCard" form:"idCard" gorm:"column:id_card;comment:身份证号;size:18;" binding:"required"` //身份证号
	MedicalTests string     `json:"medicalTests" form:"medicalTests" gorm:"column:medical_tests;comment:医学检查;size:255;"` //医学检查
	Medications  string     `json:"medications" form:"medications" gorm:"column:medications;comment:药物;size:255;"`       //药物
	Name         string     `json:"name" form:"name" gorm:"column:name;comment:就诊人员;size:50;"`                           //就诊人员
	Status       *int       `json:"status" form:"status" gorm:"column:status;comment:状态;"`                               //状态
	VisitDate    *time.Time `json:"visitDate" form:"visitDate" gorm:"column:visit_date;comment:就诊日期;"`                   //就诊日期
}

// TableName 病历表 Medical自定义表名 medical
func (Medical) TableName() string {
	return "medical"
}
