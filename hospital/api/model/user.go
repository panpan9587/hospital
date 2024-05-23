package model

type UserRegisterReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	FullName string `form:"full_name" json:"full_name" binding:"required"`
	Sex      string `form:"sex" json:"sex" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
	Mobiles  string `form:"mobile" json:"mobile" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type UserLoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 修改用户信息
type UpdateUserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	FullName string `form:"full_name" json:"full_name" binding:"required"`
	Sex      string `form:"sex" json:"sex" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

// 修改密码

type UpdatePassword struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Code   string `form:"code" json:"code" binding:"required"`
	NewPwd string `form:"new_pwd" json:"new_pwd" binding:"required"`
}
