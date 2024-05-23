package router

import (
	"demo/api/advisory"
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

		//在线咨询
		online := v1.Group("/advisory")
		{
			online.POST("/consult", advisory.Consult)             //AI在线咨询
			online.POST("/information", advisory.UserInformation) //记录历史咨询消息

		}

	}

}
