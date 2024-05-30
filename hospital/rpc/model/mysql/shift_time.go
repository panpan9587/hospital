package mysql

import (
	"gorm.io/gorm"
)

type ShiftTime struct {
	Title   string
	Count   int
	ShiftId string
	Status  int
	gorm.Model
}

func (ShiftTime) TableName() string {
	return "shift_time"
}
func GetShiftTimeById(id int) ([]*ShiftTime, error) {
	var list []*ShiftTime
	return list, DB.Where("shift_id = ?", id).Find(&list).Error
}
func GetShiftTimeBId(id int) (*ShiftTime, error) {
	var list = new(ShiftTime)
	return list, DB.Where("id = ?", id).First(&list).Error
}
func UpdateShiftTime(db *gorm.DB, shift *ShiftTime) error {
	return db.Save(&shift).Error
}
