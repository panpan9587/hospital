package mysql

import "gorm.io/gorm"

// Appointment 用户预约表结构体
type Appointment struct {
	UserID          int    `gorm:"type:int(11);not null" json:"user_id"`
	AppointmentType int    `gorm:"type:tinyint(1);not null;check:status IN (1,2);comment:1：代表预约的是体检 2：代表预约挂号" json:"appointment_type"`
	AppointmentData string `gorm:"type:varchar(20);not null" json:"appointment_data"`
	AppointmentTime string `gorm:"type:varchar(20);not null" son:"appointment_time"`
	Status          int    `gorm:"type:tinyint(1);not null;check:status IN (1,2)comment:1:已预约 2:已处理" json:"status"`
	gorm.Model
}

func (Appointment) TableName() string {
	return "appointment"
}
func GetAppointmentById(db *gorm.DB, id int) (*Appointment, error) {
	appointment := new(Appointment)
	return appointment, db.Where("id = ?", id).First(appointment).Error
}
func CreateAppointment(db *gorm.DB, appointment *Appointment) error {
	return db.Create(appointment).Error
}
func DeleteAppointment(db *gorm.DB, id int) error {
	return db.Where("id = ?", id).Delete(&Appointment{}).Error
}
func UpdatesAppointment(db *gorm.DB, id int, appointment *Appointment) error {
	return db.Where("id = ?", id).Updates(appointment).Error
}
