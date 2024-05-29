package patientmgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// InitUserRouter 初始化 患者表 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	userRouterWithoutAuth := PublicRouter.Group("user")

	var userApi = v1.ApiGroupApp.PatientmgmtApiGroup.UserApi
	{
		userRouter.POST("createPatient", userApi.CreateUser)             // 新建患者表
		userRouter.DELETE("deletePatient", userApi.DeleteUser)           // 删除患者表
		userRouter.DELETE("deletePatientByIds", userApi.DeleteUserByIds) // 批量删除患者表
		userRouter.PUT("updatePatient", userApi.UpdateUser)              // 更新患者表
	}
	{
		userRouterWithoutRecord.GET("findUser", userApi.FindUser)       // 根据ID获取患者表
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList) // 获取患者表列表
	}
	{
		userRouterWithoutAuth.GET("getUserPublic", userApi.GetUserPublic) // 获取患者表列表
	}
}
