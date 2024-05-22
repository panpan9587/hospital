package model

type UserRegisterReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Avatar   string `form:"avatar" json:"avatar" binding:"required"`
}

type UserLoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
