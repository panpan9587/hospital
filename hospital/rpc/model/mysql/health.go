package mysql

import "gorm.io/gorm"

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

// BodyInspect 体检表
type BodyInspect struct {
	ID          int64  `json:"id"`
	PeopleMsgId int64  `json:"people_msg_id"` //个人信息id
	Height      int    `json:"height"`        //身高
	Weight      int    `json:"weight"`        //体重
	Inheritance string `json:"inheritance"`   //遗传
	DoctorId    int64  `json:"doctor_id"`     //医生id
	gorm.Model
}

// MedicalItems 体检项目表
type MedicalItems struct {
	ID            int64  `json:"id"`
	Visual        string `json:"visual"`         //视力检查
	Hearing       string `json:"hearing"`        //听力检查
	BloodPressure string `json:"blood_pressure"` //血压
	BloodSugar    string `json:"blood_sugar"`    //血糖
	Urine         string `json:"urine"`          //尿常规
	Ct            int64  `json:"ct"`             //CT 1.胸部ct 2.脑部ct
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
