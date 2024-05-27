package router

import (
	"demo/api/advisory"
	_case "demo/api/case"
	"demo/api/diagnosis"
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
		}
		online := v1.Group("/advisory")
		{
			online.POST("/consult", advisory.Consult)             //AI在线咨询
			online.POST("/information", advisory.UserInformation) //记录历史咨询消息

		}
		healths := v1.Group("/health")
		{
			//记录预约
			healths.GET("/GetAppointment", health.GetAppointment)
			//体检表记录/体检项目/使用事务两表记录
			healths.GET("/GetHealth", health.GetHealth)
			//根据AppointmentId查询体检详情列表
			healths.POST("/GetHealthId", health.GetHealthId)
			//根据user_id查询体检项目详情
			healths.POST("/HealthProjectId", health.HealthProjectId)
			//科室详情
			healths.POST("/GetHealthInfo", health.GetHealthInfo)
			//套餐详情
			healths.POST("/GetPackage", health.GetPackage)
		}
		cases := v1.Group("/cases")
		{
			cases.POST("/list", _case.CaseRecordList) //病历记录
			cases.POST("/search", _case.SearchCase)   //搜索病历
		}
		doctors := v1.Group("/doctor")
		{
			//科室列表
			doctors.GET("/office/list", registration.OfficeList)
			//科室医生列表
			doctors.GET("/office/doctor/list", registration.OfficeDoctorList)
			//医生详情
			doctors.GET("/doctor/details", registration.DoctorDetails)
			//doctors.POST("/demo", registration.Demo)
		}
		//在线问诊
		chats := v1.Group("/diagnosis")
		{
			chats.GET("/chat", diagnosis.Chat) //在线聊天室
		}
	}

}
