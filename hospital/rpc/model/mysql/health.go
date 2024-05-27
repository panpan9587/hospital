package mysql

import (
	"gorm.io/gorm"
)

// Patient 患者表
type Patient struct {
	Id       int64
	UserName string //用户名
	Password string //密码
	Mobile   string //手机号
	NickName string //昵称
	Email    string //邮箱
	IsReal   int    //是否实名
	Status   int    //状态
	gorm.Model
}

// PeopleMsg 患者个人信息身份认证表
type PeopleMsg struct {
	Id       int64
	RealName string //用户真实姓名
	IdNumber string //身份证
	gorm.Model
}

// Doctor 医生表
type Doctor struct {
	Id         int64  `json:"id"`
	DoctorName string `json:"doctor_name"` //医生姓名
	Age        int    `json:"age"`         //医生年龄
	Sex        int    `json:"sex"`         //医生性别
	Position   string `json:"position"`    //医生职位
	Tag        string `json:"tag"`         //医生标签
	Desc       string `json:"desc"`        //医生简介
	WorkAge    int    `json:"work_age"`    //医生工作年限
	WorkTime   string `json:"work_time"`   //医生工作时间
	OfficeId   int64  `json:"office_id"`   //科室id
	Status     int    `json:"status"`      //状态
	gorm.Model
}

// DoctorOffice 医生科室表
type DoctorOffice struct {
	Id         int64
	OfficeName string `json:"office_name"` //科室名称
	Status     int    `json:"status"`      //状态
}

// Case 患者病例表
type Case struct {
	Id          int64  `json:"id"`
	PeopleMsgId int64  `json:"people_msg_id"` //个人信息id
	Mobile      string `json:"mobile"`        //手机号
	DiseaseName string `json:"disease_name"`  //疾病名称
	IllnessDesc string `json:"illness_desc"`  //病情描述
	Age         int    `json:"age"`           //发现是年龄
	DoctorId    int64  `json:"doctor_id"`     //医生id
	gorm.Model
}

// Health 体检表
type Health struct {
	ID            int64  `json:"id"`
	AppointmentId int64  `json:"appointment_id"` //预约id
	UserID        int64  `json:"user_id"`        //用户id
	DoctorId      int64  `json:"doctor_id"`      //医生id
	RealName      string `json:"real_name"`      //患者姓名
	IdNumber      string `json:"id_number"`      //身份证
	gorm.Model
}

// HealthInfo 体检项目详情表
type HealthInfo struct {
	ID           int64   `json:"id"`
	HealthID     int64   `json:"health_id"`      //体检id
	UserID       int64   `json:"user_id"`        //用户id
	DoctorId     int64   `json:"doctor_id"`      //医生id
	Height       float64 `json:"height"`         //身高
	Weight       float64 `json:"weight"`         //体重
	HeartRate    int64   `json:"heart_rate"`     //心率
	Hearing      string  `json:"hearing"`        //听力检查
	BloodPressur string  `json:"blood_pressure"` //血压
	BloodSugar   string  `json:"blood_sugar"`    //血糖
	Urine        string  `json:"urine"`          //尿常规
	Ct           int64   `json:"ct"`             //CT 1.胸部ct 2.脑部ct
	gorm.Model
}

// SignIn 签到表
type SignIn struct {
	Id           int64  `json:"id"`
	UserId       int    `json:"user_id"`        //用户id
	SignInMethod string `json:"sign_in_method"` //签到方式
	Status       int    `json:"status"`         //签到状态
	gorm.Model
}

// Appointment 用户预约表结构体
type Appointment struct {
	UserID          int    `gorm:"type:int(11);not null" json:"user_id"`
	AppointmentType int    `gorm:"type:tinyint(1);not null;check:status IN (1,2);comment:1：代表预约的是体检 2：代表预约挂号" json:"appointment_type"`
	Mobile          string `gorm:"type:varchar(20);not null" json:"mobile"`
	AppointmentData string `gorm:"type:varchar(20);not null" json:"appointment_data"`
	AppointmentTime string `gorm:"type:varchar(20);not null" son:"appointment_time"`
	Status          int    `gorm:"type:tinyint(1);not null;check:status IN (1,2)comment:1:已预约 2:已处理" json:"status"`
	gorm.Model
}

func (Appointment) TableName() string {
	return "appointment"
}

// AttendingPhysician 挂号记录表结构体
type AttendingPhysician struct {
	UserID        int    `gorm:"type:int(11);not null" json:"user_id"`
	AppointmentID int    `gorm:"type:int(11);not null" json:"appointment_id"`
	DoctorID      *int   `gorm:"type:int(11);not null" json:"doctor_id,omitempty"`
	OfficeID      *int   `gorm:"type:int(11);not null" json:"office_id,omitempty"`
	RealName      string `gorm:"type:varchar(50);not null" json:"real_name,omitempty"`
	Mobile        string `gorm:"type:varchar(20);not null" json:"mobile,omitempty"`
	IDNumber      string `gorm:"type:char(18);not null" json:"id_number,omitempty"`
	Symptoms      string `gorm:"type:TEXT" json:"symptoms,omitempty"`
	Diagnosis     string `gorm:"type:TEXT" json:"diagnosis,omitempty"`
	Prescription  string `gorm:"type:TEXT" json:"prescription,omitempty"`
	gorm.Model
}

func (AttendingPhysician) TableName() string {
	return "attending_physician"
}

type Package struct {
	Id            int64  `json:"Id"`
	PackageName   string `json:"package_name"`
	PackageAmount string `json:"package_amount"`
	gorm.Model
}
