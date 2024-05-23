package registration

import (
	"demo/api/model"
	proto "demo/rpc/registerationSrv/registerationclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// todo: 对应的接口api方法
func AddRegistration(ctx *gin.Context) {
	var (
		req model.AddRegistration
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	_, err = RegistrationSrv.AddRegisteration(ctx, &proto.AddAppointmentRegisterReq{
		PatientId:       int32(req.PatientID),
		DoctorId:        int32(req.DoctorID),
		RealName:        req.RealName,
		IdNumber:        req.IdNumber,
		Mobile:          req.Mobile,
		AppointmentTime: req.AppointmentTime,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "预约成功",
	})
}

// 取消预约
func CancelRegistration(ctx *gin.Context) {
	rId := ctx.PostForm("id")
	id, _ := strconv.Atoi(rId)
	_, err := RegistrationSrv.CancelAppointmentRegistration(ctx, &proto.CancelAppointmentRegistrationReq{
		AppointmentId: int32(id),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "取消预约成功",
	})
}

// 获取预约信息
func GetRegistrationById(ctx *gin.Context) {
	appointmentId := ctx.PostForm("appointmentId")
	id, _ := strconv.Atoi(appointmentId)
	res, err := RegistrationSrv.GetAppointmentRegistrationById(ctx, &proto.GetAppointmentRegistrationByIdReq{AppointmentId: int32(id)})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "获取预约信息成功",
		Data:    res.Data,
	})
}

// 修改预约信息
func UpdateRegistrationMsg(ctx *gin.Context) {
	var (
		req model.UpdateRegistration
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	_, err = RegistrationSrv.UpdateAppointmentRegistration(ctx, &proto.UpdateAppointmentRegistrationReq{Data: &proto.AppointmentRegistration{
		Id:              int32(req.AppointmentId),
		PatientId:       int32(req.PatientID),
		DoctorId:        int32(req.DoctorID),
		AppointmentTime: req.AppointmentTime,
		Status:          int32(req.Status),
	}})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "修改成功",
	})
	return
}
