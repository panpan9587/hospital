package mysql

import (
	"gorm.io/gorm"
)

// Doctor 医生表结构体
type Doctor struct {
	DoctorName  string `gorm:"type:varchar(30);not null" json:"doctor_name"`
	Age         int    `gorm:"type:int(11);not null" json:"age"`
	Sex         int    `gorm:"type:tinyint(1);not null;check:status IN (1,2);comment:1,男 2,女" json:"sex"`
	Position    string `gorm:"type:varchar(200);not null" json:"position"`
	Tag         string `gorm:"type:varchar(200);not null" json:"tag"`
	Description string `gorm:"type:TEXT;not null" json:"description"`
	WorkAge     int    `gorm:"type:int(11);not null" json:"work_age"`
	WorkTime    string `gorm:"type:varchar(50);not null" json:"work_time"`
	OfficeID    int    `gorm:"type:int(11);not null" json:"office_id"`
	Status      int    `gorm:"type:tinyint(1);not null;check:status IN (1,2);comment:状态(1:在职 2:离职)" json:"status"`
	gorm.Model
}

func (Doctor) TableName() string {
	return "doctor"
}
func GetDoctorById(id int) (*Doctor, error) {
	doctor := new(Doctor)
	return doctor, DB.Where("id = ?", id).First(doctor).Error
}
