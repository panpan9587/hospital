package model

type AddRegistration struct {
	UserId          int    `form:"user_id" json:"user_id" binding:"required"`
	DoctorID        int    `form:"doctor_id" json:"doctor_id" binding:"required"`
	RealName        string `json:"real_name" form:"real_name"`
	IdNumber        string `json:"id_number" form:"id_number"`
	Mobile          string `json:"mobile" form:"mobile"`
	AppointmentData string `form:"appointment_data" json:"appointment_data" binding:"required"`
	AppointmentType int    `form:"appointment_type" json:"appointment_type" binding:"required"`
	OfficeId        int    `form:"office_id" json:"office_id" binding:"required"`
	Status          int    `form:"status" json:"status"`
}
type UpdateRegistration struct {
	AppointmentId   int    `form:"appointment_id" json:"appointment_id" binding:"required"`
	AppointmentData string `form:"appointment_data" json:"appointment_data" binding:"required"`
	Mobile          string `json:"mobile" form:"mobile"`
	RealName        string `json:"real_name" form:"real_name"`
	IdNumber        string `json:"id_number" form:"id_number"`
	DoctorID        int    `form:"doctor_id" json:"doctor_id" binding:"required"`
	OfficeID        int    `form:"office_id" json:"office_id" binding:"required"`
	AppointmentTime string `form:"appointment_time" json:"appointment_time" binding:"required"`
	Status          int    `form:"status" json:"status"`
}
