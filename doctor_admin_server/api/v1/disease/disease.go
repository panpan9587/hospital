package disease

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/disease"
	diseaseReq "github.com/flipped-aurora/gin-vue-admin/server/model/disease/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DiseasesApi struct {
}

var diseasesService = service.ServiceGroupApp.DiseaseServiceGroup.DiseasesService

// CreateDiseases 创建病历表
// @Tags Diseases
// @Summary 创建病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body disease.Diseases true "创建病历表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /diseases/createDiseases [post]
func (diseasesApi *DiseasesApi) CreateDiseases(c *gin.Context) {
	var diseases disease.Diseases
	err := c.ShouldBindJSON(&diseases)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := diseasesService.CreateDiseases(&diseases); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDiseases 删除病历表
// @Tags Diseases
// @Summary 删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body disease.Diseases true "删除病历表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /diseases/deleteDiseases [delete]
func (diseasesApi *DiseasesApi) DeleteDiseases(c *gin.Context) {
	ID := c.Query("ID")
	if err := diseasesService.DeleteDiseases(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDiseasesByIds 批量删除病历表
// @Tags Diseases
// @Summary 批量删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /diseases/deleteDiseasesByIds [delete]
func (diseasesApi *DiseasesApi) DeleteDiseasesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := diseasesService.DeleteDiseasesByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDiseases 更新病历表
// @Tags Diseases
// @Summary 更新病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body disease.Diseases true "更新病历表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /diseases/updateDiseases [put]
func (diseasesApi *DiseasesApi) UpdateDiseases(c *gin.Context) {
	var diseases disease.Diseases
	err := c.ShouldBindJSON(&diseases)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := diseasesService.UpdateDiseases(diseases); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDiseases 用id查询病历表
// @Tags Diseases
// @Summary 用id查询病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query disease.Diseases true "用id查询病历表"
// @Success 200 {object} response.Response{data=object{rediseases=disease.Diseases},msg=string} "查询成功"
// @Router /diseases/findDiseases [get]
func (diseasesApi *DiseasesApi) FindDiseases(c *gin.Context) {
	ID := c.Query("ID")
	if rediseases, err := diseasesService.GetDiseases(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rediseases": rediseases}, c)
	}
}

// GetDiseasesList 分页获取病历表列表
// @Tags Diseases
// @Summary 分页获取病历表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query diseaseReq.DiseasesSearch true "分页获取病历表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /diseases/getDiseasesList [get]
func (diseasesApi *DiseasesApi) GetDiseasesList(c *gin.Context) {
	var pageInfo diseaseReq.DiseasesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := diseasesService.GetDiseasesInfoList(pageInfo); err != nil {
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

// GetDiseasesPublic 不需要鉴权的病历表接口
// @Tags Diseases
// @Summary 不需要鉴权的病历表接口
// @accept application/json
// @Produce application/json
// @Param data query diseaseReq.DiseasesSearch true "分页获取病历表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /diseases/getDiseasesPublic [get]
func (diseasesApi *DiseasesApi) GetDiseasesPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的病历表接口信息",
	}, "获取成功", c)
}
