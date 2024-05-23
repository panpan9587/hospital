package model

type AddRegistration struct {
	PatientID       int    `form:"patient_id" json:"patient_id" binding:"required"`
	DoctorID        int    `form:"doctor_id" json:"doctor_id" binding:"required"`
	RealName        string `json:"real_name" form:"real_name"`
	IdNumber        string `json:"id_number" form:"id_number"`
	Mobile          string `json:"mobile" form:"mobile"`
	AppointmentTime string `form:"appointment_time" json:"appointment_time" binding:"required"`
}
type UpdateRegistration struct {
	AppointmentId   int    `form:"appointment_id" json:"appointment_id" binding:"required"`
	PatientID       int    `form:"patient_id" json:"patient_id" binding:"required"`
	DoctorID        int    `form:"doctor_id" json:"doctor_id" binding:"required"`
	AppointmentTime string `form:"appointment_time" json:"appointment_time" binding:"required"`
	Status          int    `form:"status" json:"status"`
}
