package health

import (
	"context"
	"demo/api/model"
	proto "demo/rpc/healthSrv/healthclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddHealth 预约体检
func AddHealth(c *gin.Context) {
	var req model.HealthAddReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: "输入信息",
		})
		return
	}
	res, err := HealthSrv.AddBodyInspect(context.Background(), &proto.BodyInspectRequest{
		PeopleMsgId: int64(req.PeopleMsgId),
		Height:      int64(req.Height),
		Weight:      int64(req.Weight),
		Inheritance: req.Inheritance,
		DoctorId:    int64(req.DoctorId),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "注册成功",
		Data:    res,
	})
}

// GetMedicalItems 根据id查询体检项目的详情
func GetMedicalItems(c *gin.Context) {
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	res, err := HealthSrv.GetMedicalItems(context.Background(), &proto.MedicalItemsRequest{
		Id: int64(i),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "查询成功",
		Data:    res,
	})
}

// GetBodyInspect 预约体检信息详情
func GetBodyInspect(c *gin.Context) {
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	res, err := HealthSrv.GetBodyInspect(context.Background(), &proto.GetBodyInspectRequest{
		Id: int64(i),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "查询成功",
		Data:    res,
	})
}

// GetSignIn 签到记录
func GetSignIn(c *gin.Context) {
	var SignIn model.GetSignInReq
	err := c.ShouldBindJSON(&SignIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: "输入信息",
		})
		return
	}
	res, err := HealthSrv.GetSignIn(context.Background(), &proto.GetSignInRequest{
		UserId:       int64(SignIn.UserId),
		SignInMethod: SignIn.SignInMethod,
		Status:       int64(SignIn.Status),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "签到成功",
		Data:    res,
	})
}
