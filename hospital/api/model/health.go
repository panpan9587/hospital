package model

type HealthAddReq struct {
	PeopleMsgId int    `form:"peopleMsgId" json:"peopleMsgId" binding:"required"`
	Height      int    `form:"height" json:"height" binding:"required"`
	Weight      int    `form:"weight" json:"weight" binding:"required"`
	Inheritance string `form:"inheritance" json:"inheritance" binding:"`
	DoctorId    int    `form:"DoctorId" json:"DoctorId" binding:"required"`
}

type GetSignInReq struct {
	UserId       int    `form:"userId" json:"userId" binding:"required"`
	Status       int    `form:"status" json:"status" binding:"required"`
	SignInMethod string `form:"signInMethod" json:"signInMethod" binding:"`
}
