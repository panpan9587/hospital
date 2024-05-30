package doctormsg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DoctorRouter struct {
}

// InitDoctorRouter 初始化 医生表 路由信息
func (s *DoctorRouter) InitDoctorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	doctorRouter := Router.Group("doctor").Use(middleware.OperationRecord())
	doctorRouterWithoutRecord := Router.Group("doctor")
	doctorRouterWithoutAuth := PublicRouter.Group("doctor")

	var doctorApi = v1.ApiGroupApp.DoctormsgApiGroup.DoctorApi
	{
		doctorRouter.POST("createDoctor", doctorApi.CreateDoctor)             // 新建医生表
		doctorRouter.DELETE("deleteDoctor", doctorApi.DeleteDoctor)           // 删除医生表
		doctorRouter.DELETE("deleteDoctorByIds", doctorApi.DeleteDoctorByIds) // 批量删除医生表
		doctorRouter.PUT("updateDoctor", doctorApi.UpdateDoctor)              // 更新医生表
	}
	{
		doctorRouterWithoutRecord.GET("findDoctor", doctorApi.FindDoctor)       // 根据ID获取医生表
		doctorRouterWithoutRecord.GET("getDoctorList", doctorApi.GetDoctorList) // 获取医生表列表
	}
	{
		doctorRouterWithoutAuth.GET("getDoctorDataSource", doctorApi.GetDoctorDataSource) // 获取医生表数据源
		doctorRouterWithoutAuth.GET("getDoctorPublic", doctorApi.GetDoctorPublic)         // 获取医生表列表
	}
}
