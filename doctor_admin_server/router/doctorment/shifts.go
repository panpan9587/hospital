package doctorment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShiftRouter struct {
}

// InitShiftRouter 初始化 排班表 路由信息
func (s *ShiftRouter) InitShiftRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	shiftRouter := Router.Group("shift").Use(middleware.OperationRecord())
	shiftRouterWithoutRecord := Router.Group("shift")
	shiftRouterWithoutAuth := PublicRouter.Group("shift")

	var shiftApi = v1.ApiGroupApp.DoctormentApiGroup.ShiftApi
	{
		shiftRouter.POST("createShift", shiftApi.CreateShift)             // 新建排班表
		shiftRouter.DELETE("deleteShift", shiftApi.DeleteShift)           // 删除排班表
		shiftRouter.DELETE("deleteShiftByIds", shiftApi.DeleteShiftByIds) // 批量删除排班表
		shiftRouter.PUT("updateShift", shiftApi.UpdateShift)              // 更新排班表
	}
	{
		shiftRouterWithoutRecord.GET("findShift", shiftApi.FindShift)       // 根据ID获取排班表
		shiftRouterWithoutRecord.GET("getShiftList", shiftApi.GetShiftList) // 获取排班表列表
	}
	{
		shiftRouterWithoutAuth.GET("getShiftDataSource", shiftApi.GetShiftDataSource) // 获取排班表数据源
		shiftRouterWithoutAuth.GET("getShiftPublic", shiftApi.GetShiftPublic)         // 获取排班表列表
	}
}
