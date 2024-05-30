package registermgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AttendingphysicianRouter struct {
}

// InitAttendingphysicianRouter 初始化 挂号管理表 路由信息
func (s *AttendingphysicianRouter) InitAttendingphysicianRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	attendingphysicianRouter := Router.Group("attendingphysician").Use(middleware.OperationRecord())
	attendingphysicianRouterWithoutRecord := Router.Group("attendingphysician")
	attendingphysicianRouterWithoutAuth := PublicRouter.Group("attendingphysician")

	var attendingphysicianApi = v1.ApiGroupApp.RegistermgmtApiGroup.AttendingphysicianApi
	{
		attendingphysicianRouter.POST("createAttendingphysician", attendingphysicianApi.CreateAttendingphysician)             // 新建挂号管理表
		attendingphysicianRouter.DELETE("deleteAttendingphysician", attendingphysicianApi.DeleteAttendingphysician)           // 删除挂号管理表
		attendingphysicianRouter.DELETE("deleteAttendingphysicianByIds", attendingphysicianApi.DeleteAttendingphysicianByIds) // 批量删除挂号管理表
		attendingphysicianRouter.PUT("updateAttendingphysician", attendingphysicianApi.UpdateAttendingphysician)              // 更新挂号管理表
	}
	{
		attendingphysicianRouterWithoutRecord.GET("findAttendingphysician", attendingphysicianApi.FindAttendingphysician)       // 根据ID获取挂号管理表
		attendingphysicianRouterWithoutRecord.GET("getAttendingphysicianList", attendingphysicianApi.GetAttendingphysicianList) // 获取挂号管理表列表
	}
	{
		attendingphysicianRouterWithoutAuth.GET("getAttendingphysicianDataSource", attendingphysicianApi.GetAttendingphysicianDataSource) // 获取挂号管理表数据源
		attendingphysicianRouterWithoutAuth.GET("getAttendingphysicianPublic", attendingphysicianApi.GetAttendingphysicianPublic)         // 获取挂号管理表列表
	}
}
