package user

import (
	"demo/api/common"
	"demo/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 验证用户登录的中间件
func ParseToken(ctx *gin.Context) {

	// 解析出指定的token
	token := ctx.Request.Header.Get("token")
	user, err := common.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Code:    404,
			Message: "登录状态已过期",
		})
		ctx.Abort() // 阻止后续的处理器执行
		return
	} else {
		ctx.Set("user", user)
		ctx.Next()
	}
}
