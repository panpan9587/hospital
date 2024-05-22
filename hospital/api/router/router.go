package router

import (
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
		//searches := v1.Group("/search")
		//{
		//	//todo: 全文搜索
		//	searches.POST("/search")
		//	//todo: 分类查询
		//	searches.GET("/list")
		//}
	}

}
