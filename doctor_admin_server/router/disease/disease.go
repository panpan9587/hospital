package disease

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DiseasesRouter struct {
}

// InitDiseasesRouter 初始化 病历表 路由信息
func (s *DiseasesRouter) InitDiseasesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	diseasesRouter := Router.Group("diseases").Use(middleware.OperationRecord())
	diseasesRouterWithoutRecord := Router.Group("diseases")
	diseasesRouterWithoutAuth := PublicRouter.Group("diseases")

	var diseasesApi = v1.ApiGroupApp.DiseaseApiGroup.DiseasesApi
	{
		diseasesRouter.POST("createDiseases", diseasesApi.CreateDiseases)             // 新建病历表
		diseasesRouter.DELETE("deleteDiseases", diseasesApi.DeleteDiseases)           // 删除病历表
		diseasesRouter.DELETE("deleteDiseasesByIds", diseasesApi.DeleteDiseasesByIds) // 批量删除病历表
		diseasesRouter.PUT("updateDiseases", diseasesApi.UpdateDiseases)              // 更新病历表
	}
	{
		diseasesRouterWithoutRecord.GET("findDiseases", diseasesApi.FindDiseases)       // 根据ID获取病历表
		diseasesRouterWithoutRecord.GET("getDiseasesList", diseasesApi.GetDiseasesList) // 获取病历表列表
	}
	{
		diseasesRouterWithoutAuth.GET("getDiseasesPublic", diseasesApi.GetDiseasesPublic) // 获取病历表列表
	}
}
