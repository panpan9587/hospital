package registration

import (
	"demo/api/model"
	proto "demo/rpc/registerationSrv/registerationclient"
	"fmt"
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
		fmt.Println(err, ";;;")
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	_, err = RegistrationSrv.AppointmentAttendingPhysician(ctx, &proto.AppointmentAttendingPhysicianReq{
		UserId:          int32(req.UserId),
		DoctorId:        int32(req.DoctorID),
		RealName:        req.RealName,
		IdNumber:        req.IdNumber,
		Mobile:          req.Mobile,
		AppointmentData: req.AppointmentData,
		AppointmentType: int32(req.AppointmentType),
		OfficeId:        int32(req.OfficeId),
		Status:          int32(req.Status),
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
	res, err := RegistrationSrv.GetAppointmentRegistrationById(ctx, &proto.GetAppointmentRegistrationByIdReq{AttendingPhysicianId: int32(id)})
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
	_, err = RegistrationSrv.UpdateAppointmentRegistration(ctx, &proto.UpdateAppointmentRegistrationReq{
		AppData: &proto.AppointmentRegistration{
			Id:              int32(req.AppointmentId),
			Mobile:          req.Mobile,
			AppointmentDate: req.AppointmentData,
			Status:          int32(req.Status),
		},
		AttData: &proto.AttendingPhysician{
			AppointmentId: int32(req.AppointmentId),
			DoctorID:      int32(req.DoctorID),
			OfficeID:      int32(req.OfficeID),
			RealName:      req.RealName,
			Mobile:        req.Mobile,
			IDNumber:      req.IdNumber,
		},
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
		Message: "修改成功",
	})
	return
}
