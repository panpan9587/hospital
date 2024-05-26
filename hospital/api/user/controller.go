package user

/*
	info(
	desc: "用户网关中心逻辑"
	author: "panpan"
	email: "918716975@qq.com"
)
*/

import (
	"demo/api/common"
	"demo/api/model"
	"demo/api/model/redis"
	proto "demo/rpc/userSrv/userclient"
	"fmt"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	res, err := UserSrv.Register(ctx, &proto.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Sex:      req.Sex,
		Age:      int32(req.Age),
		Mobiles:  req.Mobiles,
		Email:    req.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	if res.UserId == 0 {
		ctx.JSON(http.StatusBadGateway, model.Response{
			Code:    502,
			Message: "注册失败",
		})
		return
	}
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
	res, err := UserSrv.Login(ctx, &proto.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	data, err := json.Marshal(res.User)
	if err != nil {
		zap.S().Info(err)
	}
	token, err := common.GetJwtToken(string(data))
	if err != nil {
		zap.S().Info(err)
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "登录成功",
		Data:    token,
	})
}

// 获取用户详情信息
func GetUserInfo(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: "登录状态已过期",
		})
		return
	}
	userinfo := user.(*proto.UserInfo)
	res, err := UserSrv.GetUserInfo(ctx, &proto.GetUserInfoRequest{
		UserId: userinfo.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	if res.User == nil {
		ctx.JSON(http.StatusBadGateway, model.Response{
			Code:    http.StatusBadGateway,
			Message: "用户不存在",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    res.User,
	})
}

// 修改用户信息
func UpdateUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: "登录状态已过期",
		})
		return
	}
	userinfo := user.(*proto.UserInfo)
	var req *model.UpdateUserInfo
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	// todo：修改手机号接收验证码，单独接口，验证完成才能调用此接口
	res, err := UserSrv.UpdateUser(ctx, &proto.UpdateUserRequest{
		ID:       userinfo.ID,
		Username: req.Username,
		FullName: req.FullName,
		Sex:      req.Sex,
		Age:      int32(req.Age),
		Mobile:   req.Mobile,
		Email:    req.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	if res.User == nil {
		ctx.JSON(http.StatusBadGateway, model.Response{
			Code:    http.StatusBadGateway,
			Message: "用户不存在",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "修改成功",
		Data:    res.User,
	})
}

// 获取手机验证码
func GetMobileCode(ctx *gin.Context) {
	//todo：校验手机号验证码
	mobile := ctx.PostForm("mobile")
	code, err := common.GenerateFourDigitRandomNumber()
	if err != nil {
		ctx.JSON(http.StatusForbidden, model.Response{
			Code:    http.StatusForbidden,
			Message: "验证码获取失败",
		})
		return
	}
	err = common.Sms(code, mobile)
	if err != nil {
		ctx.JSON(http.StatusForbidden, model.Response{
			Code:    http.StatusForbidden,
			Message: "验证码获取失败",
		})
		return
	}
	redis.StorageCode(code, mobile)
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "验证码获取成功",
	})
}

// 修改用户密码
func UpdatePassword(ctx *gin.Context) {
	//todo：校验手机号验证码
	var req *model.UpdatePassword
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	code := redis.GetCode(req.Mobile)
	if req.Code != code {
		ctx.JSON(http.StatusForbidden, model.Response{
			Code:    http.StatusForbidden,
			Message: "验证码错误",
		})
		return
	}
	res, err := UserSrv.UpdatePassword(ctx, &proto.UpdatePasswordRequest{
		Mobile:   req.Mobile,
		Password: req.NewPwd,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	if !res.IsOk {
		ctx.JSON(http.StatusBadGateway, model.Response{
			Code:    http.StatusBadGateway,
			Message: "密码修改失败，请重试",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "密码修改成功",
	})
}

// 注销用户
func DeleteUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: "登录状态已过期",
		})
		return
	}
	userinfo := user.(*proto.UserInfo)
	res, err := UserSrv.DeleteUser(ctx, &proto.DeleteUserRequest{
		UserId: userinfo.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	if !res.IsOk {
		ctx.JSON(http.StatusBadGateway, model.Response{
			Code:    http.StatusBadGateway,
			Message: "用户注销失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "用户注销成功",
	})
	return
}

/////////////////////////////////////////////////////////todo:待实现//////////////////////////////

func AddUserAuth(ctx *gin.Context) {

}

func UpdateUserAuth(ctx *gin.Context) {}

func GetUserAuth(ctx *gin.Context) {}

// 查看用户的体检记录
func GetHealthList(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: "登录状态已过期",
		})
		return
	}
	userinfo := user.(*proto.UserInfo)
	res, err := UserSrv.GetHealthList(ctx, &proto.GetHealthListRequest{
		UserId: userinfo.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	fmt.Println(res)
	return
}

// 查看用户的挂号记录
func GetRegistrationList(ctx *gin.Context) {

}

// 查看个人的预约
func GetAppointments(ctx *gin.Context) {

}
