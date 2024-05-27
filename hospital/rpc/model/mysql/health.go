package mysql

import "gorm.io/gorm"

type Health struct {
	UserID        int        `gorm:"type:int(11);not null" json:"user_id"`
	AppointmentId int        `gorm:"type:int(11);not null" json:"appointment_id" db:"appointment_id"`
	RealName      string     `gorm:"type:varchar(50);not null" json:"real_name,omitempty"`
	IDNumber      string     `gorm:"type:char(18);not null" json:"id_number,omitempty"`
	HealthInfo    HealthInfo `gorm:"foreignKey:HealthID"`
	gorm.Model
}

func (Health) TableName() string {
	return "health"
}

type HealthInfo struct {
	HealthID      int     `gorm:"index" json:"health_id"`
	UserID        int     `gorm:"type:int(11);not null" json:"user_id"`
	DoctorID      int     `gorm:"type:int(11);not null;comment:医生id"`
	Height        float64 `json:"height"`
	Weight        float64 `json:"weight"`
	HeartRate     int     `json:"heart_rate"`
	Hearing       int     `json:"hearing"`
	BloodPressure int     `json:"blood_pressure"`
	BloodSugar    int     `json:"blood_sugar"`
	Urine         int     `json:"urine"`
	CT            int     `json:"ct"`
	Status        int     `json:"status"`
	Description   string  `gorm:"size:255" json:"desc"`
	gorm.Model
}

func (HealthInfo) TableName() string {
	return "health_info"
}
