package health

import (
	"context"
	"demo/api/model"
	proto "demo/rpc/healthSrv/healthclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAppointment 预约表记录
func GetAppointment(c *gin.Context) {
	var appointment model.GetAppointmentReq
	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	res, err := HealthSrv.GetAppointment(context.Background(), &proto.GetAppointmentReq{
		UserID:          appointment.UserID,
		AppointmentType: appointment.AppointmentType,
		Mobile:          appointment.Mobile,
		AppointmentData: appointment.AppointmentData,
		AppointmentTime: appointment.AppointmentTime,
		Status:          appointment.Status,
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
		Message: "查询预约记录成功",
		Data:    res,
	})
}

// GetHealth 体检表/体检项目记录
func GetHealth(c *gin.Context) {
	var health model.GetHealthReq
	var project model.GetHealthInfoReq
	err := c.ShouldBindJSON(&health)
	err = c.ShouldBindJSON(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	res, err := HealthSrv.GetHealth(context.Background(), &proto.GetHealthReq{
		ID:            health.UserID,
		AppointmentId: health.AppointmentId,
		UserID:        health.UserID,
		RealName:      health.RealName,
		IdNumber:      health.IdNumber,
		ProjectInfo: &proto.HealthProjectInfo{
			ID:           project.Id,
			HealthId:     project.HealthId,
			UserID:       project.UserID,
			DoctorId:     project.DoctorId,
			Height:       uint64(project.Height),
			Weight:       uint64(project.Weight),
			HeartRate:    project.HeartRate,
			Hearing:      project.Hearing,
			BloodPressur: project.BloodPressur,
			BloodSugar:   project.BloodSugar,
			Urine:        project.Urine,
			Ct:           project.Ct,
		},
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
		Message: "体检记录成功",
		Data:    res,
	})
}

// GetHealthId 根据AppointmentId查询体检详情列表
func GetHealthId(c *gin.Context) {
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	res, err := HealthSrv.GetHealthId(context.Background(), &proto.GetHealthIdReq{
		AppointmentId: int64(i),
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
		Message: "查询体检记录详情列表成功",
		Data:    res,
	})
}

// HealthProjectId 查看体检项目详情
func HealthProjectId(c *gin.Context) {
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	res, err := HealthSrv.HealthProjectId(context.Background(), &proto.HealthProjectIdReq{UserID: int64(i)})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "查询体检功能详情成功",
		Data:    res,
	})
}

func GetHealthInfo(c *gin.Context) {
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	res, err := HealthSrv.GetDoctorOffice(context.Background(), &proto.GetDoctorOfficeReq{
		ID: int64(i),
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
		Message: "查询体检功能详情成功",
		Data:    res,
	})
}

func GetPackage(c *gin.Context) {
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	res, err := HealthSrv.GetPackage(context.Background(), &proto.GetPackageReq{
		ID: int64(i),
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
		Message: "查询套餐详情成功",
		Data:    res,
	})
}

////////////////////
