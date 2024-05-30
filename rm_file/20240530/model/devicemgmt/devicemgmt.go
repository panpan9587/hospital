// 自动生成模板Device
package devicemgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 设备表 结构体  Device
type Device struct {
	global.GVA_MODEL
	Id *int `json:"id" form:"id" gorm:"column:id;comment:;"` //主键
}

// TableName 设备表 Device自定义表名 device
func (Device) TableName() string {
	return "device"
}
