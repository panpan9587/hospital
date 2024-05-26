package mysql

import "gorm.io/gorm"

type Shift struct {
	DoctorId  int64  `gorm:"type:int(11);not null;comment:医生id"`
	Date      string `gorm:"type:varchar(30);not null;comment:排班日期"`
	StartTime string `gorm:"type:varchar(30);not null;comment:开始时间"`
	EndTime   string `gorm:"type:varchar(30not null;comment:结束时间"`
	Status    int    `gorm:"type:tinyint(1);not null;check:status IN (1,2);comment:状态(1:未预约 2:已预约)" json:"status"`
	gorm.Model
}

func (Shift) TableName() string {
	return "shift"
}
