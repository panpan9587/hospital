package mdeical

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MedicalRouter struct {
}

// InitMedicalRouter 初始化 病历表 路由信息
func (s *MedicalRouter) InitMedicalRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	medicalRouter := Router.Group("medical").Use(middleware.OperationRecord())
	medicalRouterWithoutRecord := Router.Group("medical")
	medicalRouterWithoutAuth := PublicRouter.Group("medical")

	var medicalApi = v1.ApiGroupApp.MdeicalApiGroup.MedicalApi
	{
		medicalRouter.POST("createMedical", medicalApi.CreateMedical)             // 新建病历表
		medicalRouter.DELETE("deleteMedical", medicalApi.DeleteMedical)           // 删除病历表
		medicalRouter.DELETE("deleteMedicalByIds", medicalApi.DeleteMedicalByIds) // 批量删除病历表
		medicalRouter.PUT("updateMedical", medicalApi.UpdateMedical)              // 更新病历表
	}
	{
		medicalRouterWithoutRecord.GET("findMedical", medicalApi.FindMedical)       // 根据ID获取病历表
		medicalRouterWithoutRecord.GET("getMedicalList", medicalApi.GetMedicalList) // 获取病历表列表
	}
	{
		medicalRouterWithoutAuth.GET("getMedicalPublic", medicalApi.GetMedicalPublic) // 获取病历表列表
	}
}
