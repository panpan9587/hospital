package devicemgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/devicemgmt"
	devicemgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/devicemgmt/request"
)

type DeviceService struct {
}

// CreateDevice 创建设备表记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService) CreateDevice(device *devicemgmt.Device) (err error) {
	err = global.GVA_DB.Create(device).Error
	return err
}

// DeleteDevice 删除设备表记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService) DeleteDevice(ID string) (err error) {
	err = global.GVA_DB.Delete(&devicemgmt.Device{}, "id = ?", ID).Error
	return err
}

// DeleteDeviceByIds 批量删除设备表记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService) DeleteDeviceByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]devicemgmt.Device{}, "id in ?", IDs).Error
	return err
}

// UpdateDevice 更新设备表记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService) UpdateDevice(device devicemgmt.Device) (err error) {
	err = global.GVA_DB.Model(&devicemgmt.Device{}).Where("id = ?", device.ID).Updates(&device).Error
	return err
}

// GetDevice 根据ID获取设备表记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService) GetDevice(ID string) (device devicemgmt.Device, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&device).Error
	return
}

// GetDeviceInfoList 分页获取设备表记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService) GetDeviceInfoList(info devicemgmtReq.DeviceSearch) (list []devicemgmt.Device, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&devicemgmt.Device{})
	var devices []devicemgmt.Device
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&devices).Error
	return devices, total, err
}
