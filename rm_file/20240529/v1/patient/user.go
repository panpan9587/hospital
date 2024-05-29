package patient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/patient"
	patientReq "github.com/flipped-aurora/gin-vue-admin/server/model/patient/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

var userService = service.ServiceGroupApp.PatientServiceGroup.UserService

// CreateUser 创建user表
// @Tags User
// @Summary 创建user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patient.User true "创建user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user/createUser [post]
func (userApi *UserApi) CreateUser(c *gin.Context) {
	var user patient.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.CreateUser(&user); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUser 删除user表
// @Tags User
// @Summary 删除user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patient.User true "删除user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (userApi *UserApi) DeleteUser(c *gin.Context) {
	ID := c.Query("ID")
	if err := userService.DeleteUser(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserByIds 批量删除user表
// @Tags User
// @Summary 批量删除user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /user/deleteUserByIds [delete]
func (userApi *UserApi) DeleteUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := userService.DeleteUserByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUser 更新user表
// @Tags User
// @Summary 更新user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patient.User true "更新user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateUser [put]
func (userApi *UserApi) UpdateUser(c *gin.Context) {
	var user patient.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.UpdateUser(user); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUser 用id查询user表
// @Tags User
// @Summary 用id查询user表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patient.User true "用id查询user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findUser [get]
func (userApi *UserApi) FindUser(c *gin.Context) {
	ID := c.Query("ID")
	if reuser, err := userService.GetUser(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuser": reuser}, c)
	}
}

// GetUserList 分页获取user表列表
// @Tags User
// @Summary 分页获取user表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patientReq.UserSearch true "分页获取user表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo patientReq.UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userService.GetUserInfoList(pageInfo); err != nil {
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

// GetUserPublic 不需要鉴权的user表接口
// @Tags User
// @Summary 不需要鉴权的user表接口
// @accept application/json
// @Produce application/json
// @Param data query patientReq.UserSearch true "分页获取user表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
func (userApi *UserApi) GetUserPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的user表接口信息",
	}, "获取成功", c)
}
