package devicemgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/devicemgmt"
	devicemgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/devicemgmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeviceApi struct {
}

var deviceService = service.ServiceGroupApp.DevicemgmtServiceGroup.DeviceService

// CreateDevice 创建设备表
// @Tags Device
// @Summary 创建设备表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body devicemgmt.Device true "创建设备表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /device/createDevice [post]
func (deviceApi *DeviceApi) CreateDevice(c *gin.Context) {
	var device devicemgmt.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := deviceService.CreateDevice(&device); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDevice 删除设备表
// @Tags Device
// @Summary 删除设备表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body devicemgmt.Device true "删除设备表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteDevice [delete]
func (deviceApi *DeviceApi) DeleteDevice(c *gin.Context) {
	ID := c.Query("ID")
	if err := deviceService.DeleteDevice(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceByIds 批量删除设备表
// @Tags Device
// @Summary 批量删除设备表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /device/deleteDeviceByIds [delete]
func (deviceApi *DeviceApi) DeleteDeviceByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := deviceService.DeleteDeviceByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDevice 更新设备表
// @Tags Device
// @Summary 更新设备表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body devicemgmt.Device true "更新设备表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /device/updateDevice [put]
func (deviceApi *DeviceApi) UpdateDevice(c *gin.Context) {
	var device devicemgmt.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := deviceService.UpdateDevice(device); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDevice 用id查询设备表
// @Tags Device
// @Summary 用id查询设备表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query devicemgmt.Device true "用id查询设备表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /device/findDevice [get]
func (deviceApi *DeviceApi) FindDevice(c *gin.Context) {
	ID := c.Query("ID")
	if redevice, err := deviceService.GetDevice(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redevice": redevice}, c)
	}
}

// GetDeviceList 分页获取设备表列表
// @Tags Device
// @Summary 分页获取设备表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query devicemgmtReq.DeviceSearch true "分页获取设备表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/getDeviceList [get]
func (deviceApi *DeviceApi) GetDeviceList(c *gin.Context) {
	var pageInfo devicemgmtReq.DeviceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := deviceService.GetDeviceInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetDevicePublic 不需要鉴权的设备表接口
// @Tags Device
// @Summary 不需要鉴权的设备表接口
// @accept application/json
// @Produce application/json
// @Param data query devicemgmtReq.DeviceSearch true "分页获取设备表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/getDeviceList [get]
func (deviceApi *DeviceApi) GetDevicePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的设备表接口信息",
	}, "获取成功", c)
}
