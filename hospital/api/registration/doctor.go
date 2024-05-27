package registration

import (
	"demo/api/model"
	proto "demo/rpc/registerationSrv/registerationclient"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func OfficeList(ctx *gin.Context) {
	res, err := RegistrationSrv.OfficeList(ctx, &proto.Empty{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "科室列表如下：",
		Data:    res.Data,
	})
	return
}

// 医生排班列表   根据科室id  日期
func OfficeDoctorListByIdTime(ctx *gin.Context) {
	officeId, _ := ctx.GetQuery("id")
	shiftTime, _ := ctx.GetQuery("shiftTime")
	id, _ := strconv.Atoi(officeId)
	fmt.Println(officeId, shiftTime, "///")
	res, err := RegistrationSrv.GetOfficeDoctorListByIdTime(ctx, &proto.OfficeDoctorListByIdTimeReq{
		OfficeId:  int32(id),
		ShiftTime: shiftTime,
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
		Message: fmt.Sprintf("科室医生列表如下："),
		Data:    res.Data,
	})
	return
}

// 医生详情
func DoctorDetails(ctx *gin.Context) {
	doctor_id := ctx.PostForm("id")
	id, _ := strconv.Atoi(doctor_id)
	res, err := RegistrationSrv.DoctorDetails(ctx, &proto.DoctorDetailsReq{
		DoctorId: int32(id),
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
		Message: fmt.Sprintf("%s医生信息如下：", res.Data.DoctorName),
		Data:    res.Data,
	})
	return
}
