// 自动生成模板Diseases
package disease

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 病历表 结构体  Diseases
type Diseases struct {
	global.GVA_MODEL
	Name         string     `json:"name" form:"name" gorm:"column:name;comment:就诊人员;size:50;"`                           //就诊人员
	IdNumber     string     `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:身份证号;size:20;"`              //身份证号
	VisitDate    *time.Time `json:"visitDate" form:"visitDate" gorm:"column:visit_date;comment:就诊日期;"`                   //就诊日期
	Diagnosis    string     `json:"diagnosis" form:"diagnosis" gorm:"column:diagnosis;comment:诊断;size:100;"`             //诊断
	Medications  string     `json:"medications" form:"medications" gorm:"column:medications;comment:药物;size:255;"`       //药物
	Allergies    string     `json:"allergies" form:"allergies" gorm:"column:allergies;comment:过敏史;size:50;"`             //过敏史
	MedicalTests string     `json:"medicalTests" form:"medicalTests" gorm:"column:medical_tests;comment:医学检查;size:255;"` //医学检查
	Status       *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:19;"`                       //状态
	IdCard       string     `json:"idCard" form:"idCard" gorm:"column:id_card;comment:身份证号;size:18;"`                    //身份证号
}

// TableName 病历表 Diseases自定义表名 disease
func (Diseases) TableName() string {
	return "disease"
}
