package mysql

import "gorm.io/gorm"

// todo: 用户表,头像未设置，待加入
type User struct {
	ID           int    `gorm:"type:int(11);primarykey;not null" json:"id"`
	Username     string `gorm:"type:varchar(18);not null" json:"username"`
	Password     string `gorm:"type:char(32);not null" json:"password"`
	FullName     string `gorm:"type:varchar(50);" json:"full_name"`
	Sex          string `gorm:"type:type:enum('男', '女');not null" json:"sex"`
	Age          int    `gorm:"type:int(3);not null" json:"age"`
	Mobile       string `gorm:"type:varchar(20);not null;unique" json:"mobile"`
	HealthScore  int    `gorm:"type:int(11)" json:"health_score"`
	Email        string `gorm:"type:varchar(50)" json:"email"`
	Subscription int8   `gorm:"type:tinyint(1)" json:"subscription"`
	Status       int8   `gorm:"type:tinyint(1)" json:"status"`
	gorm.Model
}

// todo: 用户认证表
type UserAuth struct {
	ID       int    `gorm:"type:int(11);primarykey;not null" json:"id"`
	UserID   int    `gorm:"type:int(11);not null" json:"user_id"`
	RealName string `gorm:"type:varchar(30);not null" json:"real_name"`
	IdNumber string `gorm:"type:char(18);not null" json:"id_number"`
	Status   int8   `gorm:"type:tinyint(1)" json:"status"`
	gorm.Model
}

func (User) TableName() string {
	return "user"
}

func (UserAuth) TableName() string {
	return "user_auth"
}

// 添加用户用于用户注册
func AddUser(user *User) (err error) {
	err = DB.Create(&user).Error
	return
}

func GetUserByUsername(username string) (user *User, err error) {
	err = DB.Where("username=?", username).First(&user).Error
	return
}

func GetUserById(userid int) (user *User, err error) {
	err = DB.Where("id=?", userid).First(&user).Error
	return
}

func GetUserByMobile(mobile string) (user *User, err error) {
	err = DB.Where("mobile = ?", mobile).First(&user).Error
	return
}

func UpdateUser(user *User) (err error) {
	err = DB.Save(&user).Error
	return
}

func DeleteUser(userid int) (err error) {
	err = DB.Where("id = ?", userid).Delete(&User{}).Error
	return
}
