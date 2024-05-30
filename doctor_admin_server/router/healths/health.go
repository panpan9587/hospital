package healths

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HealthRouter struct {
}

// InitHealthRouter 初始化 体检表 路由信息
func (s *HealthRouter) InitHealthRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	healthRouter := Router.Group("health").Use(middleware.OperationRecord())
	healthRouterWithoutRecord := Router.Group("health")
	healthRouterWithoutAuth := PublicRouter.Group("health")

	var healthApi = v1.ApiGroupApp.HealthsApiGroup.HealthApi
	{
		healthRouter.POST("createHealth", healthApi.CreateHealth)             // 新建体检表
		healthRouter.DELETE("deleteHealth", healthApi.DeleteHealth)           // 删除体检表
		healthRouter.DELETE("deleteHealthByIds", healthApi.DeleteHealthByIds) // 批量删除体检表
		healthRouter.PUT("updateHealth", healthApi.UpdateHealth)              // 更新体检表
	}
	{
		healthRouterWithoutRecord.GET("findHealth", healthApi.FindHealth)       // 根据ID获取体检表
		healthRouterWithoutRecord.GET("getHealthList", healthApi.GetHealthList) // 获取体检表列表
	}
	{
		healthRouterWithoutAuth.GET("getHealthPublic", healthApi.GetHealthPublic) // 获取体检表列表
	}
}
