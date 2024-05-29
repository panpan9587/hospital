package patient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// InitUserRouter 初始化 用户表 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	userRouterWithoutAuth := PublicRouter.Group("user")

	var userApi = v1.ApiGroupApp.PatientApiGroup.UserApi
	{
		userRouter.POST("createPatient", userApi.CreateUser)             // 新建用户表
		userRouter.DELETE("deletePatient", userApi.DeleteUser)           // 删除患者表
		userRouter.DELETE("deletePatientByIds", userApi.DeleteUserByIds) // 批量删除用户表
		userRouter.PUT("updatePatient", userApi.UpdateUser)              // 更新用户表
	}
	{
		userRouterWithoutRecord.GET("findUser", userApi.FindUser)       // 根据ID获取用户表
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList) // 获取用户表列表
	}
	{
		userRouterWithoutAuth.GET("getUserPublic", userApi.GetUserPublic) // 获取用户表列表
	}
}
