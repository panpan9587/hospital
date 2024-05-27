package user

import (
	"demo/api/common"
	"demo/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 验证用户登录的中间件
func ParseToken(ctx *gin.Context) {

	// 解析出指定的token
	//tokenString := ctx.Request.Header.Get("token")
	tokenString := ctx.Request.Header.Get("Authorization") //获取token信息
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: "请先登录",
		})
		ctx.Abort() // 阻止后续的处理器执行
		return
	}
	// 检查是否以Bearer开头
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing Bearer token"})
		ctx.Abort()
		return
	}
	// 提取token值
	token := parts[1]
	//解析token
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
