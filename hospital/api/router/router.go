package router

import (
	"demo/api/advisory"
	_case "demo/api/case"
	"demo/api/health"
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
			users.POST("/code", user.GetMobileCode) //获取验证码
			users.POST("/update/password", user.UpdatePassword)
			users.Use(user.ParseToken)

			users.GET("/info", user.GetUserInfo) //详情不展示实名信息
			// todo：中间件校验手机号是否一致
			users.POST("/update", user.UpdateUser)
			users.POST("/auth", user.AddUserAuth)                //用户实名认证
			users.POST("/update/auth", user.UpdateUserAuth)      //修改用户实名信息
			users.POST("/get/auth", user.GetUserAuth)            //查看用户实名信息
			users.POST("/delete", user.DeleteUser)               //注销用户
			users.GET("/registration", user.GetRegistrationList) //查看个人的挂号纪录
			users.GET("/appointments", user.GetAppointments)     // 查询个人的预约
			users.GET("/health", user.GetHealthList)             //查看个人的体检纪录
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
		online := v1.Group("/advisory")
		{
			online.POST("/consult", advisory.Consult)             //AI在线咨询
			online.POST("/information", advisory.UserInformation) //记录历史咨询消息

		}
		healths := v1.Group("/health")
		{
			//预约体检信息记录
			healths.GET("/AddHealth", health.AddHealth)
			//获取体检项目信息
			healths.GET("/GetMedicalItems", health.GetMedicalItems)
			//预约体检信息详情
			healths.GET("/GetBodyInspect", health.GetBodyInspect)
			//签到记录
			healths.GET("/GetSignIn", health.GetSignIn)
		}
		cases := v1.Group("/cases")
		{
			cases.POST("/list", _case.CaseRecordList) //病历记录
			cases.POST("/search", _case.SearchCase)   //搜索病历
		}

	}

}
