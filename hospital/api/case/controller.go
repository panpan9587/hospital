package _case

import (
	"demo/api/model"
	proto "demo/rpc/caseSrv/caseclient"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CaseRecordList 病历记录列表
func CaseRecordList(ctx *gin.Context) {
	index := ctx.PostForm("index")
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	res, err := CaseSrv.CaseRecordList(ctx, &proto.CaseRecordListReq{
		Index: index,
		Page:  int64(page),
	})
	if err != nil {
		// 处理调用失败情况
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("调用CaseRecordList失败的原因: %s", err.Error()),
		})
		return
	}
	// 返回成功响应
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res.Result,
	})
}

// SearchCase 搜索病历记录
func SearchCase(ctx *gin.Context) {
	content := ctx.PostForm("content")
	index := ctx.PostForm("index")
	res, err := CaseSrv.SearchCaseRecord(ctx, &proto.SearchCaseRecordReq{
		Content: content,
		Index:   index,
	})
	if err != nil {
		// 处理调用失败情况
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	//返回成功的响应
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res.Result,
	})
}
