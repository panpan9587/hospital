package appointmentmgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/appointmentmgmt"
	appointmentmgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/appointmentmgmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppointmentApi struct {
}

var appointmentService = service.ServiceGroupApp.AppointmentmgmtServiceGroup.AppointmentService

// CreateAppointment 创建预约记录表
// @Tags Appointment
// @Summary 创建预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appointmentmgmt.Appointment true "创建预约记录表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /appointment/createAppointment [post]
func (appointmentApi *AppointmentApi) CreateAppointment(c *gin.Context) {
	var appointment appointmentmgmt.Appointment
	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appointmentService.CreateAppointment(&appointment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppointment 删除预约记录表
// @Tags Appointment
// @Summary 删除预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appointmentmgmt.Appointment true "删除预约记录表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /appointment/deleteAppointment [delete]
func (appointmentApi *AppointmentApi) DeleteAppointment(c *gin.Context) {
	ID := c.Query("ID")
	if err := appointmentService.DeleteAppointment(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppointmentByIds 批量删除预约记录表
// @Tags Appointment
// @Summary 批量删除预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /appointment/deleteAppointmentByIds [delete]
func (appointmentApi *AppointmentApi) DeleteAppointmentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := appointmentService.DeleteAppointmentByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppointment 更新预约记录表
// @Tags Appointment
// @Summary 更新预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appointmentmgmt.Appointment true "更新预约记录表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /appointment/updateAppointment [put]
func (appointmentApi *AppointmentApi) UpdateAppointment(c *gin.Context) {
	var appointment appointmentmgmt.Appointment
	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appointmentService.UpdateAppointment(appointment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppointment 用id查询预约记录表
// @Tags Appointment
// @Summary 用id查询预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appointmentmgmt.Appointment true "用id查询预约记录表"
// @Success 200 {object} response.Response{data=object{reappointment=appointmentmgmt.Appointment},msg=string} "查询成功"
// @Router /appointment/findAppointment [get]
func (appointmentApi *AppointmentApi) FindAppointment(c *gin.Context) {
	ID := c.Query("ID")
	if reappointment, err := appointmentService.GetAppointment(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappointment": reappointment}, c)
	}
}

// GetAppointmentList 分页获取预约记录表列表
// @Tags Appointment
// @Summary 分页获取预约记录表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appointmentmgmtReq.AppointmentSearch true "分页获取预约记录表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /appointment/getAppointmentList [get]
func (appointmentApi *AppointmentApi) GetAppointmentList(c *gin.Context) {
	var pageInfo appointmentmgmtReq.AppointmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appointmentService.GetAppointmentInfoList(pageInfo); err != nil {
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

// GetAppointmentDataSource 获取Appointment的数据源
// @Tags Appointment
// @Summary 获取Appointment的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /appointment/getAppointmentDataSource [get]
func (appointmentApi *AppointmentApi) GetAppointmentDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := appointmentService.GetAppointmentDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetAppointmentPublic 不需要鉴权的预约记录表接口
// @Tags Appointment
// @Summary 不需要鉴权的预约记录表接口
// @accept application/json
// @Produce application/json
// @Param data query appointmentmgmtReq.AppointmentSearch true "分页获取预约记录表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appointment/getAppointmentPublic [get]
func (appointmentApi *AppointmentApi) GetAppointmentPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的预约记录表接口信息",
	}, "获取成功", c)
}
