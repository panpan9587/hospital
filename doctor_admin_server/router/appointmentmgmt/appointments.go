package appointmentmgmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppointmentRouter struct {
}

// InitAppointmentRouter 初始化 预约记录表 路由信息
func (s *AppointmentRouter) InitAppointmentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	appointmentRouter := Router.Group("appointment").Use(middleware.OperationRecord())
	appointmentRouterWithoutRecord := Router.Group("appointment")
	appointmentRouterWithoutAuth := PublicRouter.Group("appointment")

	var appointmentApi = v1.ApiGroupApp.AppointmentmgmtApiGroup.AppointmentApi
	{
		appointmentRouter.POST("createAppointment", appointmentApi.CreateAppointment)             // 新建预约记录表
		appointmentRouter.DELETE("deleteAppointment", appointmentApi.DeleteAppointment)           // 删除预约记录表
		appointmentRouter.DELETE("deleteAppointmentByIds", appointmentApi.DeleteAppointmentByIds) // 批量删除预约记录表
		appointmentRouter.PUT("updateAppointment", appointmentApi.UpdateAppointment)              // 更新预约记录表
	}
	{
		appointmentRouterWithoutRecord.GET("findAppointment", appointmentApi.FindAppointment)       // 根据ID获取预约记录表
		appointmentRouterWithoutRecord.GET("getAppointmentList", appointmentApi.GetAppointmentList) // 获取预约记录表列表
	}
	{
		appointmentRouterWithoutAuth.GET("getAppointmentDataSource", appointmentApi.GetAppointmentDataSource) // 获取预约记录表数据源
		appointmentRouterWithoutAuth.GET("getAppointmentPublic", appointmentApi.GetAppointmentPublic)         // 获取预约记录表列表
	}
}
