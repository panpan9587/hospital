package doctorment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/doctorment"
	doctormentReq "github.com/flipped-aurora/gin-vue-admin/server/model/doctorment/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShiftApi struct {
}

var shiftService = service.ServiceGroupApp.DoctormentServiceGroup.ShiftService

// CreateShift 创建排班表
// @Tags Shift
// @Summary 创建排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctorment.Shift true "创建排班表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /shift/createShift [post]
func (shiftApi *ShiftApi) CreateShift(c *gin.Context) {
	var shift doctorment.Shift
	err := c.ShouldBindJSON(&shift)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := shiftService.CreateShift(&shift); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShift 删除排班表
// @Tags Shift
// @Summary 删除排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctorment.Shift true "删除排班表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /shift/deleteShift [delete]
func (shiftApi *ShiftApi) DeleteShift(c *gin.Context) {
	ID := c.Query("ID")
	if err := shiftService.DeleteShift(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShiftByIds 批量删除排班表
// @Tags Shift
// @Summary 批量删除排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /shift/deleteShiftByIds [delete]
func (shiftApi *ShiftApi) DeleteShiftByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := shiftService.DeleteShiftByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShift 更新排班表
// @Tags Shift
// @Summary 更新排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctorment.Shift true "更新排班表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /shift/updateShift [put]
func (shiftApi *ShiftApi) UpdateShift(c *gin.Context) {
	var shift doctorment.Shift
	err := c.ShouldBindJSON(&shift)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := shiftService.UpdateShift(shift); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShift 用id查询排班表
// @Tags Shift
// @Summary 用id查询排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query doctorment.Shift true "用id查询排班表"
// @Success 200 {object} response.Response{data=object{reshift=doctorment.Shift},msg=string} "查询成功"
// @Router /shift/findShift [get]
func (shiftApi *ShiftApi) FindShift(c *gin.Context) {
	ID := c.Query("ID")
	if reshift, err := shiftService.GetShift(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshift": reshift}, c)
	}
}

// GetShiftList 分页获取排班表列表
// @Tags Shift
// @Summary 分页获取排班表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query doctormentReq.ShiftSearch true "分页获取排班表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /shift/getShiftList [get]
func (shiftApi *ShiftApi) GetShiftList(c *gin.Context) {
	var pageInfo doctormentReq.ShiftSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shiftService.GetShiftInfoList(pageInfo); err != nil {
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

// GetShiftDataSource 获取Shift的数据源
// @Tags Shift
// @Summary 获取Shift的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /shift/getShiftDataSource [get]
func (shiftApi *ShiftApi) GetShiftDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := shiftService.GetShiftDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetShiftPublic 不需要鉴权的排班表接口
// @Tags Shift
// @Summary 不需要鉴权的排班表接口
// @accept application/json
// @Produce application/json
// @Param data query doctormentReq.ShiftSearch true "分页获取排班表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /shift/getShiftPublic [get]
func (shiftApi *ShiftApi) GetShiftPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的排班表接口信息",
	}, "获取成功", c)
}
