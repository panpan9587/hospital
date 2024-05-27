package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var Router *gin.Engine
var wg sync.WaitGroup

func main() {
	Router = gin.New()
	//Router.LoadHTMLGlob("./views/*") //todo:视图层
	//Router.GET("index", func(c *gin.Context) {
	//	c.HTML(200, "index.html", gin.H{
	//		"data": "",
	//	})
	//})
	chars := Router.Group("/chars")
	{
		chars.GET("/chat", Chat)
	}
	Router.Run(":9999")
}

var ClientMap = map[string]*Client{}

type Client struct {
	Conn []*websocket.Conn
	Send chan []byte
}

func Chat(c *gin.Context) {
	//userid := "1"
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 建立用户连接
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
	//握手并建立websocket 连接
	// 建立ai的连接
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	aiconn, _, err := d.Dial((assembleAuthUrl1("wss://spark-api.xf-yun.com/v1.1/chat", "37d7a17ca08d0da8ec27777a19cc3a92", "NTBmNjZlMmZiOTYyNTk4YzE2NjQ5YmJl")), nil)
	if err != nil {
		fmt.Println(err)
	}
	//ClientMap[userid].Conn = append(ClientMap[userid].Conn, conn)
	var connes []*websocket.Conn
	connes = append(connes, conn, aiconn)
	node := &Client{
		Conn: connes,
		Send: make(chan []byte, 1024),
	}
	wg.Add(3)
	go ReadMessage(node)
	go ReadAiMessage(node)
	go SendMessage(node)
	wg.Wait()
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

// 加密函数
func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func ReadMessage(node *Client) {
	defer wg.Done()
	for {
		_, message, err := node.Conn[0].ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(string(message), "读取到用户发送的信息")
		//ClientMap[userid].Send <- message
		data := GenParams1("b9ba97ee", string(message))
		node.Conn[1].WriteJSON(data)
		//todo:校验消息的格式
	}
}

// todo:向客户端发送消息，接收者不需要调用
func SendMessage(node *Client) {
	defer wg.Done()
	for {
		select {
		//todo: 带写入返回值
		// 写入用户发送的消息
		case context := <-node.Send:
			fmt.Println(context, "接收到返回的消息")
			err := node.Conn[0].WriteMessage(websocket.TextMessage, context)
			if err != nil {
				log.Println("消息接收失败")
			}
		}
	}
}

var answer = ""

func ReadAiMessage(node *Client) {
	defer wg.Done()
	for {
		_, message, err := node.Conn[1].ReadMessage()
		if err != nil {
			log.Println("接收返回错误", err) // 修改这里，应该打印错误信息
			break                      // 退出循环
		}
		// 对结果返回处理
		log.Println("接收到的返回结果", string(message))
		var data map[string]interface{}
		err1 := json.Unmarshal(message, &data)
		if err1 != nil {
			fmt.Println("Error parsing JSON:", err1) // 修改这里，应该打印错误信息
			break                                    // 退出循环
		}

		// 检查是否存在预期的字段
		payload, ok := data["payload"].(map[string]interface{})
		if !ok {
			log.Println("Missing payload field")
			break // 退出循环
		}
		choices, ok := payload["choices"].(map[string]interface{})
		if !ok {
			log.Println("Missing choices field")
			break // 退出循环
		}
		header, ok := data["header"].(map[string]interface{})
		if !ok {
			log.Println("Missing header field")
			break // 退出循环
		}

		code, ok := header["code"].(float64)
		if !ok {
			log.Println("Invalid code field")
			break // 退出循环
		}

		if code != 0 {
			break // 退出循环
		}
		status, ok := choices["status"].(float64)
		if !ok {
			log.Println("Invalid status field")
			break // 退出循环
		}
		text, ok := choices["text"].([]interface{})
		if !ok {
			log.Println("Invalid text field")
			break // 退出循环
		}
		if len(text) == 0 {
			log.Println("Text field is empty")
			break // 退出循环
		}
		content, ok := text[0].(map[string]interface{})["content"].(string)
		if !ok {
			log.Println("Invalid content field")
			break // 退出循环
		}
		answer += content

		if status == 2 {
			usage, ok := payload["usage"].(map[string]interface{})
			if !ok {
				log.Println("Missing usage field")
				break // 退出循环
			}
			temp, ok := usage["text"].(map[string]interface{})
			if !ok {
				log.Println("Missing text field in usage")
				break // 退出循环
			}
			totalTokens, ok := temp["total_tokens"].(float64)
			if !ok {
				log.Println("Invalid total_tokens field")
				break // 退出循环
			}
			fmt.Println("total_tokens:", totalTokens)
			node.Conn[1].Close()
			break // 退出循环
		}

		//node.Send <- []byte(answer)
	}
	node.Send <- []byte(answer)
}

// 拼接的发送消息
func GenParams1(appid, question string) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名

	messages := []Message{
		{Role: "user", Content: question},
	}

	data := map[string]interface{}{
		"header": map[string]interface{}{
			"app_id": appid,
		},
		"parameter": map[string]interface{}{
			"chat": map[string]interface{}{
				"domain":      "general",
				"temperature": float64(0.8),
				"top_k":       int64(6),
				"max_tokens":  int64(2048),
				"auditing":    "default",
			},
		},
		"payload": map[string]interface{}{
			"message": map[string]interface{}{
				"text": messages,
			},
		},
	}
	return data
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

//func SendMessage(node *Client) {
//	defer wg.Done()
//	for {
//		select {
//		case context := <-node.Send:
//			err := node.Conn.WriteMessage(websocket.TextMessage, context)
//			if err != nil {
//				zap.S().Info("消息写入失败")
//			}
//		}
//	}
//
//}
