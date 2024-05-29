package userauth

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userauth"
	userauthReq "github.com/flipped-aurora/gin-vue-admin/server/model/userauth/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserAuthApi struct {
}

var userAuthService = service.ServiceGroupApp.UserauthServiceGroup.UserAuthService

// CreateUserAuth 创建实名认证表
// @Tags UserAuth
// @Summary 创建实名认证表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userauth.UserAuth true "创建实名认证表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userAuth/createUserAuth [post]
func (userAuthApi *UserAuthApi) CreateUserAuth(c *gin.Context) {
	var userAuth userauth.UserAuth
	err := c.ShouldBindJSON(&userAuth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userAuthService.CreateUserAuth(&userAuth); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserAuth 删除实名认证表
// @Tags UserAuth
// @Summary 删除实名认证表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userauth.UserAuth true "删除实名认证表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAuth/deleteUserAuth [delete]
func (userAuthApi *UserAuthApi) DeleteUserAuth(c *gin.Context) {
	ID := c.Query("ID")
	if err := userAuthService.DeleteUserAuth(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserAuthByIds 批量删除实名认证表
// @Tags UserAuth
// @Summary 批量删除实名认证表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userAuth/deleteUserAuthByIds [delete]
func (userAuthApi *UserAuthApi) DeleteUserAuthByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := userAuthService.DeleteUserAuthByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserAuth 更新实名认证表
// @Tags UserAuth
// @Summary 更新实名认证表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userauth.UserAuth true "更新实名认证表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAuth/updateUserAuth [put]
func (userAuthApi *UserAuthApi) UpdateUserAuth(c *gin.Context) {
	var userAuth userauth.UserAuth
	err := c.ShouldBindJSON(&userAuth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userAuthService.UpdateUserAuth(userAuth); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserAuth 用id查询实名认证表
// @Tags UserAuth
// @Summary 用id查询实名认证表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userauth.UserAuth true "用id查询实名认证表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAuth/findUserAuth [get]
func (userAuthApi *UserAuthApi) FindUserAuth(c *gin.Context) {
	ID := c.Query("ID")
	if reuserAuth, err := userAuthService.GetUserAuth(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserAuth": reuserAuth}, c)
	}
}

// GetUserAuthList 分页获取实名认证表列表
// @Tags UserAuth
// @Summary 分页获取实名认证表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userauthReq.UserAuthSearch true "分页获取实名认证表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAuth/getUserAuthList [get]
func (userAuthApi *UserAuthApi) GetUserAuthList(c *gin.Context) {
	var pageInfo userauthReq.UserAuthSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userAuthService.GetUserAuthInfoList(pageInfo); err != nil {
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

func (userAuthApi *UserAuthApi) GetUserAuthDataSource(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的实名认证表接口信息",
	}, "获取成功", c)
}

// GetUserAuthPublic 不需要鉴权的实名认证表接口
// @Tags UserAuth
// @Summary 不需要鉴权的实名认证表接口
// @accept application/json
// @Produce application/json
// @Param data query userauthReq.UserAuthSearch true "分页获取实名认证表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAuth/getUserAuthList [get]
func (userAuthApi *UserAuthApi) GetUserAuthPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的实名认证表接口信息",
	}, "获取成功", c)
}
