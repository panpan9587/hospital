package advisory

import (
	"demo/api/common"
	"demo/api/etc"
	"demo/api/model"
	proto "demo/rpc/advisorySrv/advisoryclient"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// 开始进行AI对话
func Consult(ctx *gin.Context) {
	contents := ctx.PostForm("content")

	//调用第三方AI接口

	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	//握手并建立websocket 连接
	conn, resp, err := d.Dial(common.AssembleAuthUrl1(etc.ApiConfig.AISrv.HostUrl, etc.ApiConfig.AISrv.ApiKey, etc.ApiConfig.AISrv.ApiSecret), nil)
	if err != nil {
		panic(common.ReadResp(resp) + err.Error())
		return
	} else if resp.StatusCode != 101 {
		panic(common.ReadResp(resp) + err.Error())
	}

	go func() {

		data := common.GenParams1(etc.ApiConfig.AISrv.Appid, contents)
		conn.WriteJSON(data)

	}()

	var answer = ""
	//获取返回的数据
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var data map[string]interface{}
		err1 := json.Unmarshal(msg, &data)
		if err1 != nil {
			return
		}
		//解析数据
		payload := data["payload"].(map[string]interface{})
		choices := payload["choices"].(map[string]interface{})
		header := data["header"].(map[string]interface{})
		code := header["code"].(float64)

		if code != 0 {
			return
		}
		status := choices["status"].(float64)
		text := choices["text"].([]interface{})
		content := text[0].(map[string]interface{})["content"].(string)
		if status != 2 {
			answer += content
		} else {
			answer += content
			usage := payload["usage"].(map[string]interface{})
			temp := usage["text"].(map[string]interface{})
			totalTokens := temp["total_tokens"].(float64)
			fmt.Println("total_tokens:", totalTokens)
			conn.Close()
			break
		}

	}
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    answer,
	})

	time.Sleep(1 * time.Second)

}

// 记录历史咨询消息（包含敏感词过滤）
func UserInformation(ctx *gin.Context) {
	// 获取参数
	contents := ctx.PostForm("content")

	// 建立 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(common.AssembleAuthUrl1(etc.ApiConfig.AISrv.HostUrl, etc.ApiConfig.AISrv.ApiKey, etc.ApiConfig.AISrv.ApiSecret), nil)
	if err != nil {
		// 处理连接失败情况
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("日志含义建立WebSocket连接失败的原因: %s", err.Error()),
		})
		return
	}
	defer conn.Close() // 在函数退出时关闭连接

	// 异步发送数据
	go func() {
		data := common.GenParams1(etc.ApiConfig.AISrv.Appid, contents)
		conn.WriteJSON(data)
	}()

	var answer string

	// 接收返回数据
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// 处理读取消息错误
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("从WebSocket连接读取消息失败的原因: %s", err.Error()),
			})
			return
		}

		var data map[string]interface{}
		if err := json.Unmarshal(msg, &data); err != nil {
			// 处理 JSON 解析错误
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("日志含义解析JSON数据失败的原因: %s", err.Error()),
			})
			return
		}

		// 解析数据
		payload := data["payload"].(map[string]interface{})
		choices := payload["choices"].(map[string]interface{})
		header := data["header"].(map[string]interface{})
		code := header["code"].(float64)

		if code != 0 {
			// 处理错误码不为0的情况
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("收到非零错误码: %v", code),
			})
			return
		}

		status := choices["status"].(float64)
		text := choices["text"].([]interface{})
		content := text[0].(map[string]interface{})["content"].(string)
		answer += content

		if status == 2 {
			// 处理 status 为2的情况
			usage := payload["usage"].(map[string]interface{})
			temp := usage["text"].(map[string]interface{})
			totalTokens := temp["total_tokens"].(float64)
			fmt.Println("total_tokens:", totalTokens)
			break
		}
	}

	// 调用 AdvisoriesSrv.RecordsUserConsultationInformation
	res, err := AdvisoriesSrv.RecordsUserConsultationInformation(ctx, &proto.RecordsUserConsultationInformationReq{
		Content: contents,
		Result:  answer,
	})
	if err != nil {
		// 处理调用失败情况
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("调用RecordsUserConsultationInformation失败的原因: %s", err.Error()),
		})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res.DocumentId,
	})
}
