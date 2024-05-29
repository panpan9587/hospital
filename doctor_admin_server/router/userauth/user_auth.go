package userauth

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserAuthRouter struct {
}

// InitUserAuthRouter 初始化 实名认证表 路由信息
func (s *UserAuthRouter) InitUserAuthRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userAuthRouter := Router.Group("userAuth").Use(middleware.OperationRecord())
	userAuthRouterWithoutRecord := Router.Group("userAuth")
	userAuthRouterWithoutAuth := PublicRouter.Group("userAuth")

	var userAuthApi = v1.ApiGroupApp.UserauthApiGroup.UserAuthApi
	{
		userAuthRouter.POST("createUserAuth", userAuthApi.CreateUserAuth)             // 新建实名认证表
		userAuthRouter.DELETE("deleteUserAuth", userAuthApi.DeleteUserAuth)           // 删除实名认证表
		userAuthRouter.DELETE("deleteUserAuthByIds", userAuthApi.DeleteUserAuthByIds) // 批量删除实名认证表
		userAuthRouter.PUT("updateUserAuth", userAuthApi.UpdateUserAuth)              // 更新实名认证表
	}
	{
		userAuthRouterWithoutRecord.GET("findUserAuth", userAuthApi.FindUserAuth)       // 根据ID获取实名认证表
		userAuthRouterWithoutRecord.GET("getUserAuthList", userAuthApi.GetUserAuthList) // 获取实名认证表列表
	}
	{
		userAuthRouterWithoutAuth.GET("getUserAuthDataSource", userAuthApi.GetUserAuthDataSource) // 获取实名认证表数据源
		userAuthRouterWithoutAuth.GET("getUserAuthPublic", userAuthApi.GetUserAuthPublic)         // 获取实名认证表列表
	}
}
