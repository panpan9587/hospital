// 自动生成模板Shift
package doctorment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 排班表 结构体  Shift
type Shift struct {
	global.GVA_MODEL
	DoctorId *int       `json:"doctorId" form:"doctorId" gorm:"column:doctor_id;comment:医生id;size:10;"` //医生
	Date     *time.Time `json:"date" form:"date" gorm:"column:date;comment:日期;"`                        //日期
	Count    *int       `json:"count" form:"count" gorm:"column:count;comment:总票数;size:10;"`            //总票数
	OfficeId *int       `json:"officeId" form:"officeId" gorm:"column:office_id;comment:科室id;size:10;"` //科室
	Status   *int       `json:"status" form:"status" gorm:"column:status;comment:1：有票  2：无票;size:10;"`  //状态
}

// TableName 排班表 Shift自定义表名 shift
func (Shift) TableName() string {
	return "shift"
}
