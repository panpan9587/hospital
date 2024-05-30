package doctormsg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/doctormsg"
	doctormsgReq "github.com/flipped-aurora/gin-vue-admin/server/model/doctormsg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DoctorApi struct {
}

var doctorService = service.ServiceGroupApp.DoctormsgServiceGroup.DoctorService

// CreateDoctor 创建医生表
// @Tags Doctor
// @Summary 创建医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctormsg.Doctor true "创建医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /doctor/createDoctor [post]
func (doctorApi *DoctorApi) CreateDoctor(c *gin.Context) {
	var doctor doctormsg.Doctor
	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := doctorService.CreateDoctor(&doctor); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDoctor 删除医生表
// @Tags Doctor
// @Summary 删除医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctormsg.Doctor true "删除医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doctor/deleteDoctor [delete]
func (doctorApi *DoctorApi) DeleteDoctor(c *gin.Context) {
	ID := c.Query("ID")
	if err := doctorService.DeleteDoctor(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDoctorByIds 批量删除医生表
// @Tags Doctor
// @Summary 批量删除医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /doctor/deleteDoctorByIds [delete]
func (doctorApi *DoctorApi) DeleteDoctorByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := doctorService.DeleteDoctorByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDoctor 更新医生表
// @Tags Doctor
// @Summary 更新医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctormsg.Doctor true "更新医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /doctor/updateDoctor [put]
func (doctorApi *DoctorApi) UpdateDoctor(c *gin.Context) {
	var doctor doctormsg.Doctor
	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := doctorService.UpdateDoctor(doctor); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDoctor 用id查询医生表
// @Tags Doctor
// @Summary 用id查询医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query doctormsg.Doctor true "用id查询医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /doctor/findDoctor [get]
func (doctorApi *DoctorApi) FindDoctor(c *gin.Context) {
	ID := c.Query("ID")
	if redoctor, err := doctorService.GetDoctor(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redoctor": redoctor}, c)
	}
}

// GetDoctorList 分页获取医生表列表
// @Tags Doctor
// @Summary 分页获取医生表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query doctormsgReq.DoctorSearch true "分页获取医生表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /doctor/getDoctorList [get]
func (doctorApi *DoctorApi) GetDoctorList(c *gin.Context) {
	var pageInfo doctormsgReq.DoctorSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := doctorService.GetDoctorInfoList(pageInfo); err != nil {
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

// GetDoctorPublic 不需要鉴权的医生表接口
// @Tags Doctor
// @Summary 不需要鉴权的医生表接口
// @accept application/json
// @Produce application/json
// @Param data query doctormsgReq.DoctorSearch true "分页获取医生表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /doctor/getDoctorList [get]
func (doctorApi *DoctorApi) GetDoctorPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的医生表接口信息",
	}, "获取成功", c)
}

// GetDoctoressDataSource 获取Doctoress的数据源
// @Tags Doctoress
// @Summary 获取Doctoress的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /doctoress/getDoctoressDataSource [get]
func (doctorApi *DoctorApi) GetDoctorDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := doctorService.GetDoctorDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}
