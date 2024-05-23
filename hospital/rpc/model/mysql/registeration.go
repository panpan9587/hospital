package mysql

import (
	"gorm.io/gorm"
	"time"
)

// 预约表
type PatientAppointment struct {
	ID              int            `gorm:"type:int(11);not null" json:"id"`
	PatientID       int            `gorm:"type:int(11)" json:"patient_id" db:"patient_id"`
	DoctorID        int            `gorm:"type:int(11)" json:"doctor_id" db:"doctor_id"`
	RealName        string         `gorm:"type:varchar(30)" json:"real_name" db:"real_name"`
	IdNumber        string         `gorm:"type:char(18)" json:"id_number" db:"id_number"`
	Mobile          string         `gorm:"type:char(11)" json:"mobile" db:"mobile"`
	AppointmentTime time.Time      `gorm:"type:varchar(30)" json:"appointment_time" db:"appointment_time"`
	Status          int            `gorm:"type:tinyint(1);check:status IN (1,2)" json:"status"`
	CreatedAt       *time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at" db:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at" db:"deleted_at"`
}

func (PatientAppointment) TableName() string {
	return "patient_appointment"
}
