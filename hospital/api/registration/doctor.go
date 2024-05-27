package registration

import (
	"crypto/hmac"
	"crypto/sha256"
	"demo/api/model"
	proto "demo/rpc/registerationSrv/registerationclient"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
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
func OfficeDoctorList(ctx *gin.Context) {
	officeId := ctx.PostForm("id")
	id, _ := strconv.Atoi(officeId)
	res, err := RegistrationSrv.OfficeDoctorList(ctx, &proto.OfficeDoctorListReq{
		OfficeId: int32(id),
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
		Message: fmt.Sprintf("%s科室医生列表如下：", res.Data.OfficeName),
		Data:    res.Data,
	})
	return
}
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

type Messages struct {
	FriendId int    `json:"friend"`
	Content  string `json:"content"`
}
type Node struct {
	Conn *websocket.Conn
	Data chan []byte
}

var (
	ClientMap = map[int]*Node{}
	wg        sync.WaitGroup
)

func Demo(ctx *gin.Context) {
	userid := ctx.PostForm("userid")
	id, _ := strconv.Atoi(userid)
	upgrader := websocket.Upgrader{
		HandshakeTimeout: 0,
		ReadBufferSize:   0,
		WriteBufferSize:  0,
		WriteBufferPool:  nil,
		Subprotocols:     nil,
		Error:            nil,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: false,
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		panic(err)
	}
	node := &Node{
		Conn: conn,
		Data: make(chan []byte, 50),
	}
	log.Println(id, "./././")
	ClientMap[int(id)] = node
	wg.Add(2)
	go SendMsg(node)
	go ReadMsg(node)
	wg.Wait()
}
func SendMsg(node *Node) {
	defer wg.Done()
	for {
		_, message, err := node.Conn.ReadMessage()
		if err != nil {
			return
		}
		var msg Messages
		json.Unmarshal(message, &msg)
		FriendMap, ok := ClientMap[msg.FriendId]
		if ok {
			FriendMap.Data <- []byte(msg.Content)
		}
	}
}
func ReadMsg(node *Node) {
	defer wg.Done()
	for {
		select {
		case data := <-node.Data:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return
			}
		}
	}
}

// 生成参数
func genParams1(appid, question string) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名

	messages := []Message{
		{Role: "user", Content: question},
	}

	data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
		"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"app_id": appid, // 根据实际情况修改返回的数据结构和字段名
		},
		"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"domain":      "general",    // 根据实际情况修改返回的数据结构和字段名
				"temperature": float64(0.8), // 根据实际情况修改返回的数据结构和字段名
				"top_k":       int64(6),     // 根据实际情况修改返回的数据结构和字段名
				"max_tokens":  int64(2048),  // 根据实际情况修改返回的数据结构和字段名
				"auditing":    "default",    // 根据实际情况修改返回的数据结构和字段名
			},
		},
		"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"text": messages, // 根据实际情况修改返回的数据结构和字段名
			},
		},
	}
	return data // 根据实际情况修改返回的数据结构和字段名
}

// 创建鉴权url  apikey 即 hmac username
func assembleAuthUrl1(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	// fmt.Println(sgin)
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
