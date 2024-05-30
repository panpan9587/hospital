package healths

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/healths"
	healthsReq "github.com/flipped-aurora/gin-vue-admin/server/model/healths/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HealthApi struct {
}

var healthService = service.ServiceGroupApp.HealthsServiceGroup.HealthService

// CreateHealth 创建体检表
// @Tags Health
// @Summary 创建体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body healths.Health true "创建体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /health/createHealth [post]
func (healthApi *HealthApi) CreateHealth(c *gin.Context) {
	var health healths.Health
	err := c.ShouldBindJSON(&health)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := healthService.CreateHealth(&health); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHealth 删除体检表
// @Tags Health
// @Summary 删除体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body healths.Health true "删除体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /health/deleteHealth [delete]
func (healthApi *HealthApi) DeleteHealth(c *gin.Context) {
	ID := c.Query("ID")
	if err := healthService.DeleteHealth(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHealthByIds 批量删除体检表
// @Tags Health
// @Summary 批量删除体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /health/deleteHealthByIds [delete]
func (healthApi *HealthApi) DeleteHealthByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := healthService.DeleteHealthByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHealth 更新体检表
// @Tags Health
// @Summary 更新体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body healths.Health true "更新体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /health/updateHealth [put]
func (healthApi *HealthApi) UpdateHealth(c *gin.Context) {
	var health healths.Health
	err := c.ShouldBindJSON(&health)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := healthService.UpdateHealth(health); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHealth 用id查询体检表
// @Tags Health
// @Summary 用id查询体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query healths.Health true "用id查询体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /health/findHealth [get]
func (healthApi *HealthApi) FindHealth(c *gin.Context) {
	ID := c.Query("ID")
	if rehealth, err := healthService.GetHealth(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehealth": rehealth}, c)
	}
}

// GetHealthList 分页获取体检表列表
// @Tags Health
// @Summary 分页获取体检表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query healthsReq.HealthSearch true "分页获取体检表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /health/getHealthList [get]
func (healthApi *HealthApi) GetHealthList(c *gin.Context) {
	var pageInfo healthsReq.HealthSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := healthService.GetHealthInfoList(pageInfo); err != nil {
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

// GetHealthPublic 不需要鉴权的体检表接口
// @Tags Health
// @Summary 不需要鉴权的体检表接口
// @accept application/json
// @Produce application/json
// @Param data query healthsReq.HealthSearch true "分页获取体检表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /health/getHealthList [get]
func (healthApi *HealthApi) GetHealthPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的体检表接口信息",
	}, "获取成功", c)
}
