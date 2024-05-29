package patient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// InitUserRouter 初始化 user表 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	userRouterWithoutAuth := PublicRouter.Group("user")

	var userApi = v1.ApiGroupApp.PatientApiGroup.UserApi
	{
		userRouter.POST("createUser", userApi.CreateUser)             // 新建user表
		userRouter.DELETE("deleteUser", userApi.DeleteUser)           // 删除user表
		userRouter.DELETE("deleteUserByIds", userApi.DeleteUserByIds) // 批量删除user表
		userRouter.PUT("updateUser", userApi.UpdateUser)              // 更新user表
	}
	{
		userRouterWithoutRecord.GET("findUser", userApi.FindUser)       // 根据ID获取user表
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList) // 获取user表列表
	}
	{
		userRouterWithoutAuth.GET("getUserPublic", userApi.GetUserPublic) // 获取user表列表
	}
}
