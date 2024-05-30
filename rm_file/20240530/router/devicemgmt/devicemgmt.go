package devicemgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct {
}

// InitDeviceRouter 初始化 设备表 路由信息
func (s *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	deviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	deviceRouterWithoutRecord := Router.Group("device")
	deviceRouterWithoutAuth := PublicRouter.Group("device")

	var deviceApi = v1.ApiGroupApp.DevicemgmtApiGroup.DeviceApi
	{
		deviceRouter.POST("createDevice", deviceApi.CreateDevice)             // 新建设备表
		deviceRouter.DELETE("deleteDevice", deviceApi.DeleteDevice)           // 删除设备表
		deviceRouter.DELETE("deleteDeviceByIds", deviceApi.DeleteDeviceByIds) // 批量删除设备表
		deviceRouter.PUT("updateDevice", deviceApi.UpdateDevice)              // 更新设备表
	}
	{
		deviceRouterWithoutRecord.GET("findDevice", deviceApi.FindDevice)       // 根据ID获取设备表
		deviceRouterWithoutRecord.GET("getDeviceList", deviceApi.GetDeviceList) // 获取设备表列表
	}
	{
		deviceRouterWithoutAuth.GET("getDevicePublic", deviceApi.GetDevicePublic) // 获取设备表列表
	}
}
