package diagnosis

import (
	"demo/api/model"
	proto "demo/rpc/userSrv/userclient"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

/*
	info(
	desc: "在线问诊模块逻辑"
	author: "panpan"
	email: "918716975@qq.com"
)
*/
// todo: 建立websocket连接

var wg sync.WaitGroup

var ClientMap = make(map[int]*model.Client)

// 连线医生
func AddDoctorChat(ctx *gin.Context) {
	// 逻辑添加对应的问诊，对医生端推送连线请求

	// 将请求发送至消息队，医生端订阅消息队列，进行提示

}

// todo:逻辑优化，等待连接
func Chat(ctx *gin.Context) {
	//实现一对一聊天
	//从token中获取用户登录信息
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: "登录状态已过期",
		})
		return
	}
	doctor_id := ctx.Query("doctor_id")
	userinfo := user.(*proto.UserInfo)
	//upgrader := websocket.Upgrader{
	//	CheckOrigin: func(r *http.Request) bool {
	//		return true
	//	},
	//}
	// todo: 连接对应的socket指定地址，判断连接成功或失败
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("wss://doctor-server-url/ws?doctor_id=%s", doctor_id), nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: "医生连线失败",
		})
		return
	}
	node := &model.Client{
		Conn: conn,
		Send: make(chan []byte, 50),
	}
	ClientMap[int(userinfo.ID)] = node
	wg.Add(2)
	go ReadMessage(node)
	go SendMessage(node)
	wg.Wait()
}

func ReadMessage(node *model.Client) {
	defer wg.Done()
	for {
		_, message, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("消息解析失败")
			return
		}
		var msg model.Message
		//todo:校验消息的格式
		err = json.Unmarshal(message, &msg)
		if err != nil {
			ClientMap[msg.UserID].Send <- []byte("消息发送是败")
			return
		}
		ClientMap[msg.DisId].Send <- message
	}
}

func SendMessage(node *model.Client) {
	defer wg.Done()
	for {
		select {
		case context := <-node.Send:
			err := node.Conn.WriteMessage(websocket.TextMessage, context)
			if err != nil {
				zap.S().Info("消息写入失败")
			}
		}
	}
}
