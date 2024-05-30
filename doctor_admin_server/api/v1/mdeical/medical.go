package mdeical

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mdeical"
	mdeicalReq "github.com/flipped-aurora/gin-vue-admin/server/model/mdeical/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MedicalApi struct {
}

var medicalService = service.ServiceGroupApp.MdeicalServiceGroup.MedicalService

// CreateMedical 创建病历表
// @Tags Medical
// @Summary 创建病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mdeical.Medical true "创建病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /medical/createMedical [post]
func (medicalApi *MedicalApi) CreateMedical(c *gin.Context) {
	var medical mdeical.Medical
	err := c.ShouldBindJSON(&medical)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := medicalService.CreateMedical(&medical); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMedical 删除病历表
// @Tags Medical
// @Summary 删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mdeical.Medical true "删除病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /medical/deleteMedical [delete]
func (medicalApi *MedicalApi) DeleteMedical(c *gin.Context) {
	ID := c.Query("ID")
	if err := medicalService.DeleteMedical(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMedicalByIds 批量删除病历表
// @Tags Medical
// @Summary 批量删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /medical/deleteMedicalByIds [delete]
func (medicalApi *MedicalApi) DeleteMedicalByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := medicalService.DeleteMedicalByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMedical 更新病历表
// @Tags Medical
// @Summary 更新病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mdeical.Medical true "更新病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /medical/updateMedical [put]
func (medicalApi *MedicalApi) UpdateMedical(c *gin.Context) {
	var medical mdeical.Medical
	err := c.ShouldBindJSON(&medical)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := medicalService.UpdateMedical(medical); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMedical 用id查询病历表
// @Tags Medical
// @Summary 用id查询病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mdeical.Medical true "用id查询病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /medical/findMedical [get]
func (medicalApi *MedicalApi) FindMedical(c *gin.Context) {
	ID := c.Query("ID")
	if remedical, err := medicalService.GetMedical(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remedical": remedical}, c)
	}
}

// GetMedicalList 分页获取病历表列表
// @Tags Medical
// @Summary 分页获取病历表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mdeicalReq.MedicalSearch true "分页获取病历表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /medical/getMedicalList [get]
func (medicalApi *MedicalApi) GetMedicalList(c *gin.Context) {
	var pageInfo mdeicalReq.MedicalSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := medicalService.GetMedicalInfoList(pageInfo); err != nil {
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

// GetMedicalPublic 不需要鉴权的病历表接口
// @Tags Medical
// @Summary 不需要鉴权的病历表接口
// @accept application/json
// @Produce application/json
// @Param data query mdeicalReq.MedicalSearch true "分页获取病历表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /medical/getMedicalList [get]
func (medicalApi *MedicalApi) GetMedicalPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的病历表接口信息",
	}, "获取成功", c)
}
