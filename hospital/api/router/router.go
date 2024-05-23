package router

import (
	"demo/api/registration"
	"demo/api/user"
	"github.com/gin-gonic/gin"
)

func initRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"data": "success",
			})
		})
		users := v1.Group("/user")
		{
			users.POST("/register", user.Register)
			users.POST("/login", user.Login)
		}
		registrations := v1.Group("/registration")
		{
			//预约
			registrations.POST("/add", registration.AddRegistration)
			//取消预约
			registrations.POST("/cancel", registration.CancelRegistration)
			//获取预约信息
			registrations.GET("/get/id", registration.GetRegistrationById)
			//修改预约信息
			registrations.POST("/update", registration.UpdateRegistrationMsg)
		}
		//searches := v1.Group("/search")
		//{
		//	//todo: 全文搜索
		//	searches.POST("/search")
		//	//todo: 分类查询
		//	searches.GET("/list")
		//}
	}

}
