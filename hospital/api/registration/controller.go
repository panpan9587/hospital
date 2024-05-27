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
	_, err = RegistrationSrv.CreateAppointment(ctx, &proto.CreateAppointmentReq{
		UserId:          int32(req.UserId),
		OfficeId:        int32(req.OfficeId),
		AppointmentType: int32(req.AppointmentType),
		AppointmentData: req.AppointmentData,
		Status:          int32(req.Status),
		DoctorId:        int32(req.DoctorID),
		RealName:        req.RealName,
		IdNumber:        req.IdNumber,
		Mobile:          req.Mobile,
		ShiftId:         int32(req.ShiftId),
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
	shiftId := ctx.PostForm("shift_id")
	id, _ := strconv.Atoi(rId)
	sid, _ := strconv.Atoi(shiftId)
	_, err := RegistrationSrv.CancelAppointmentRegistration(ctx, &proto.CancelAppointmentRegistrationReq{
		AppointmentId: int32(id),
		ShiftId:       int32(sid),
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
func GetAppointmentById(ctx *gin.Context) {
	var (
		req  model.Appointment
		data = map[string]interface{}{}
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: "请输入正确的用户信息",
		})
		return
	}
	appointment, err := RegistrationSrv.GetAppointmentById(ctx, &proto.GetAppointmentByIdReq{AppointmentId: int32(req.AppointmentId)})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	data["id"] = appointment.Data.Id
	data["userid"] = appointment.Data.UserId
	data["appointment_type"] = appointment.Data.AppointmentType
	data["appointment_data"] = appointment.Data.AppointmentData
	data["appointment_time"] = appointment.Data.AppointmentTime
	data["status"] = appointment.Data.Status
	if req.AppointmentType == 1 {
		//health.HealthSrv.
	} else {
		res, err := RegistrationSrv.GetAttendingPhysicianByAppointmentId(ctx, &proto.GetAttendingPhysicianByAppointmentIdReq{AppointmentId: int32(req.AppointmentId)})
		if err != nil {
			return
		}
		data["doctor_id"] = res.Data.DoctorID
		data["office_id"] = res.Data.OfficeID
		data["shift_id"] = res.Data.ShiftId
		data["real_name"] = res.Data.RealName
		data["id_number"] = res.Data.IDNumber
		data["mobile"] = res.Data.Mobile
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "获取预约信息成功",
		Data:    data,
	})
}

// 获取挂号信息  （根据身份证号）
func GetAttendingByIdNumber(ctx *gin.Context) {
	idNumber := ctx.PostForm("idNumber")
	res, err := RegistrationSrv.GetAttendingPhysicianByIdNumber(ctx, &proto.GetAttendingPhysicianByIdNumberReq{IdNumber: idNumber})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "获取挂号信息成功",
		Data: map[string]interface{}{
			"data": res.Data,
		},
	})
}
