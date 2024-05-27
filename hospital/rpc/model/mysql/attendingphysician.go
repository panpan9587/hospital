package mysql

import (
	"gorm.io/gorm"
)

// AttendingPhysician 挂号记录表结构体
type AttendingPhysician struct {
	UserID        int    `gorm:"type:int(11);not null" json:"user_id"`
	AppointmentId int    `gorm:"type:int(11);not null" json:"appointment_id" db:"appointment_id"`
	DoctorID      int    `gorm:"type:int(11);not null" json:"doctor_id,omitempty"`
	OfficeID      int    `gorm:"type:int(11);not null" json:"office_id,omitempty"`
	ShiftId       int    `gorm:"type:int(11);not null" json:"shift_id,omitempty"`
	RealName      string `gorm:"type:varchar(50);not null" json:"real_name,omitempty"`
	IDNumber      string `gorm:"type:char(18);not null" json:"id_number,omitempty"`
	Mobile        string `gorm:"type:varchar(20);not null" json:"mobile"`
	Symptoms      string `gorm:"type:TEXT" json:"symptoms,omitempty"`
	Diagnosis     string `gorm:"type:TEXT" json:"diagnosis,omitempty"`
	Prescription  string `gorm:"type:TEXT" json:"prescription,omitempty"`
	gorm.Model
}

func (AttendingPhysician) TableName() string {
	return "attendingphysician"
}
func CreateAttendingPhysician(db *gorm.DB, physician *AttendingPhysician) error {
	return db.Create(physician).Error
}
func DeleteAttendingPhysician(db *gorm.DB, id int) error {
	return db.Where("appointment_id = ?", id).Delete(&AttendingPhysician{}).Error
}
func GetAttendingPhysicianById(db *gorm.DB, id int) (*AttendingPhysician, error) {
	var (
		attending_physician = new(AttendingPhysician)
	)
	return attending_physician, db.Where("appointment_id = ?", id).First(attending_physician).Error
}

func GetAttendingPhysicianByIdNumber(db *gorm.DB, idNumber string) ([]*AttendingPhysician, error) {
	var (
		attending_physician []*AttendingPhysician
	)
	return attending_physician, db.Where("id_number = ?", idNumber).Find(&attending_physician).Error
}
