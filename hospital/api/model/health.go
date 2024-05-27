package model

// GetAppointmentReq 预约表记录
type GetAppointmentReq struct {
	UserID          int64  `form:"UserID" json:"user_id" binding:"required"`
	AppointmentType int64  `form:"AppointmentType" json:"appointment_type" binding:"required"`
	Mobile          string `form:"Mobile" json:"mobile" binding:"required"`
	AppointmentData string `form:"AppointmentData" json:"appointment_data" binding:"required"`
	AppointmentTime string `form:"AppointmentTime" json:"appointment_time" binding:"required"`
	Status          int64  `form:"Status" json:"status" binding:"required"`
}

// GetHealthReq 体检表记录
type GetHealthReq struct {
	AppointmentId int64  `form:"AppointmentId" json:"appointment_id" binding:"required"`
	UserID        int64  `form:"status" json:"user_id" binding:"required"`
	RealName      string `form:"real_name" json:"real_name" binding:""`
	IdNumber      string `form:"id_number" json:"id_number" binding:""`
}

// GetHealthInfoReq 体检项目表记录
type GetHealthInfoReq struct {
	Id           int64   `form:"Id" json:"id" binding:"required"`
	HealthId     int64   `form:"HealthId" json:"health_id" binding:"required"`
	UserID       int64   `form:"status" json:"user_id" binding:"required"`
	DoctorId     int64   `form:"DoctorId" json:"doctor_id" binding:"required"`
	Height       float64 `form:"Height" json:"height" binding:"required"`
	Weight       float64 `form:"Weight" json:"weight" binding:"required"`
	HeartRate    int64   `form:"HeartRate" json:"heart_rate" binding:"required"`
	Hearing      string  `form:"Hearing" json:"hearing" binding:"required"`
	BloodPressur string  `form:"BloodPressure" json:"blood_pressur" binding:"required"`
	BloodSugar   string  `form:"BloodSugar" json:"blood_sugar" binding:"required"`
	Urine        string  `form:"Urine" json:"urine" binding:"required"`
	Ct           int64   `form:"Ct" json:"ct" binding:"required"`
}
