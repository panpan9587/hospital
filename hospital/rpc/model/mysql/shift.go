package mysql

import "gorm.io/gorm"

type Shift struct {
	OfficeID int    `gorm:"type:int(11);not null" json:"office_id"`
	DoctorId int64  `gorm:"type:int(11);not null;comment:医生id"`
	Date     string `gorm:"type:varchar(30);not null;comment:排班日期"`
	Time     int    `gorm:"type:tinyint(1);not null;check:time IN (1,2);comment:状态(1:上午 2:下午)" json:"time"`
	Count    int    `gorm:"type:int(11);not null;comment:预约数"`
	Status   int    `gorm:"type:tinyint(1);not null;check:status IN (1,2);comment:状态(1:未预约 2:已预约)" json:"status"`
	gorm.Model
}

func (Shift) TableName() string {
	return "shift"
}
func GetOfficeDoctorListByIdTime(id int, time int) ([]*Shift, error) {
	var list []*Shift
	return list, DB.Where("office_id = ?", id).Where("time = ?", time).Find(&list).Error
}
func GetShiftDoctorById(id int) (*Shift, error) {
	var res = new(Shift)
	return res, DB.Where("id = ?", id).First(&res).Error
}
func UpdateShift(db *gorm.DB, shift *Shift) error {
	return db.Save(shift).Error
}
