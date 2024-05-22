package user

import (
	"demo/api/model"
	proto "demo/rpc/userSrv/userclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo: 对应的接口api方法
func Register(ctx *gin.Context) {
	var req model.UserRegisterReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	res, err := UserSrv.Register(ctx, &proto.RegisterRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	//if res.UserId == 0 {
	//	ctx.JSON(http.StatusBadGateway, model.Response{
	//		Code:    502,
	//		Message: "注册失败",
	//	})
	//	return
	//}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "注册成功",
		Data:    res,
	})
}

func Login(ctx *gin.Context) {
	var req *model.UserLoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	res, err := UserSrv.Login(ctx, &proto.LoginRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	//data, err := json.Marshal(res.User)
	//if err != nil {
	//	zap.S().Info(err)
	//}
	//token, err := pkg.GetJwtToken(string(data))
	//if err != nil {
	//	zap.S().Info(err)
	//}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "登录成功",
		Data:    res,
	})
}
