package registermgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/registermgmt"
	registermgmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/registermgmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AttendingphysicianApi struct {
}

var attendingphysicianService = service.ServiceGroupApp.RegistermgmtServiceGroup.AttendingphysicianService

// CreateAttendingphysician 创建挂号管理表
// @Tags Attendingphysician
// @Summary 创建挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body registermgmt.Attendingphysician true "创建挂号管理表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /attendingphysician/createAttendingphysician [post]
func (attendingphysicianApi *AttendingphysicianApi) CreateAttendingphysician(c *gin.Context) {
	var attendingphysician registermgmt.Attendingphysician
	err := c.ShouldBindJSON(&attendingphysician)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := attendingphysicianService.CreateAttendingphysician(&attendingphysician); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAttendingphysician 删除挂号管理表
// @Tags Attendingphysician
// @Summary 删除挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body registermgmt.Attendingphysician true "删除挂号管理表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /attendingphysician/deleteAttendingphysician [delete]
func (attendingphysicianApi *AttendingphysicianApi) DeleteAttendingphysician(c *gin.Context) {
	ID := c.Query("ID")
	if err := attendingphysicianService.DeleteAttendingphysician(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAttendingphysicianByIds 批量删除挂号管理表
// @Tags Attendingphysician
// @Summary 批量删除挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /attendingphysician/deleteAttendingphysicianByIds [delete]
func (attendingphysicianApi *AttendingphysicianApi) DeleteAttendingphysicianByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := attendingphysicianService.DeleteAttendingphysicianByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAttendingphysician 更新挂号管理表
// @Tags Attendingphysician
// @Summary 更新挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body registermgmt.Attendingphysician true "更新挂号管理表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /attendingphysician/updateAttendingphysician [put]
func (attendingphysicianApi *AttendingphysicianApi) UpdateAttendingphysician(c *gin.Context) {
	var attendingphysician registermgmt.Attendingphysician
	err := c.ShouldBindJSON(&attendingphysician)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := attendingphysicianService.UpdateAttendingphysician(attendingphysician); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAttendingphysician 用id查询挂号管理表
// @Tags Attendingphysician
// @Summary 用id查询挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query registermgmt.Attendingphysician true "用id查询挂号管理表"
// @Success 200 {object} response.Response{data=object{reattendingphysician=registermgmt.Attendingphysician},msg=string} "查询成功"
// @Router /attendingphysician/findAttendingphysician [get]
func (attendingphysicianApi *AttendingphysicianApi) FindAttendingphysician(c *gin.Context) {
	ID := c.Query("ID")
	if reattendingphysician, err := attendingphysicianService.GetAttendingphysician(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reattendingphysician": reattendingphysician}, c)
	}
}

// GetAttendingphysicianList 分页获取挂号管理表列表
// @Tags Attendingphysician
// @Summary 分页获取挂号管理表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query registermgmtReq.AttendingphysicianSearch true "分页获取挂号管理表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /attendingphysician/getAttendingphysicianList [get]
func (attendingphysicianApi *AttendingphysicianApi) GetAttendingphysicianList(c *gin.Context) {
	var pageInfo registermgmtReq.AttendingphysicianSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := attendingphysicianService.GetAttendingphysicianInfoList(pageInfo); err != nil {
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

// GetAttendingphysicianDataSource 获取Attendingphysician的数据源
// @Tags Attendingphysician
// @Summary 获取Attendingphysician的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /attendingphysician/getAttendingphysicianDataSource [get]
func (attendingphysicianApi *AttendingphysicianApi) GetAttendingphysicianDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := attendingphysicianService.GetAttendingphysicianDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetAttendingphysicianPublic 不需要鉴权的挂号管理表接口
// @Tags Attendingphysician
// @Summary 不需要鉴权的挂号管理表接口
// @accept application/json
// @Produce application/json
// @Param data query registermgmtReq.AttendingphysicianSearch true "分页获取挂号管理表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /attendingphysician/getAttendingphysicianPublic [get]
func (attendingphysicianApi *AttendingphysicianApi) GetAttendingphysicianPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的挂号管理表接口信息",
	}, "获取成功", c)
}
