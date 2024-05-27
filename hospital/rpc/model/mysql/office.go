package mysql

import "gorm.io/gorm"

// DoctorOffice 科室表结构体
type DoctorOffice struct {
	OfficeName string `gorm:"type:varchar(200);not null" json:"office_name"`
	Pid        int    `gorm:"type:int(11);not null" json:"pid"`
	Status     int    `gorm:"type:tinyint(1);not null;check:status IN (1,2)" json:"status"`
	gorm.Model
	Doctor []*Doctor `gorm:"foreignKey:office_id"`
}

func (DoctorOffice) TableName() string {
	return "doctor_office"
}
func GetOfficeList() ([]*DoctorOffice, error) {
	var list []*DoctorOffice
	return list, DB.Find(&list).Error
}
func GetOfficeDoctorList(id int) (*DoctorOffice, error) {
	var list = new(DoctorOffice)
	return list, DB.Preload("Doctor").Where("id = ?", id).First(&list).Error
}
