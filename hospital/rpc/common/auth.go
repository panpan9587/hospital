package common

/*
	info(
	desc: "实名认证第三方"
	author: "panpan"
	email: "918716975@qq.com"
)
*/
import (
	"crypto/hmac"
	"crypto/sha1"
	"demo/api/etc"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	gourl "net/url"
	"strings"
	"time"
)

type Response struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
	Result    struct {
		Realname    string `json:"realname"`
		Idcard      string `json:"idcard"`
		Isok        bool   `json:"isok"`
		IdCardInfor struct {
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
			Area     string `json:"area"`
			Sex      string `json:"sex"`
			Birthday string `json:"birthday"`
		} `json:"IdCardInfor"`
	} `json:"result"`
}

func calcAuthorization(source string, secretId string, secretKey string) (auth string, datetime string, err error) {
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	datetime = time.Now().In(timeLocation).Format("Mon, 02 Jan 2006 15:04:05 GMT")
	signStr := fmt.Sprintf("x-date: %s\nx-source: %s", datetime, source)

	// hmac-sha1
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(signStr))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	auth = fmt.Sprintf("hmac id=\"%s\", algorithm=\"hmac-sha1\", headers=\"x-date x-source\", signature=\"%s\"",
		secretId, sign)

	return auth, datetime, nil
}

func urlencode(params map[string]string) string {
	var p = gourl.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	return p.Encode()
}

func AuthUser(cardNo, realName string) (res *Response, err error) {
	// 云市场分配的密钥Id
	secretId := etc.ApiConfig.Auth.SecretID
	// 云市场分配的密钥Key
	secretKey := etc.ApiConfig.Auth.SecretKey
	source := "market"

	// 签名
	auth, datetime, _ := calcAuthorization(source, secretId, secretKey)

	// 请求方法
	method := "POST"
	// 请求头
	headers := map[string]string{"X-Source": source, "X-Date": datetime, "Authorization": auth}

	// 查询参数
	queryParams := make(map[string]string)

	// body参数
	bodyParams := make(map[string]string)
	bodyParams["cardNo"] = cardNo
	bodyParams["realName"] = realName
	// url参数拼接
	url := "https://service-hsyjtg0f-1301232119.bj.apigw.tencentcs.com/release/id/check"
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s?%s", url, urlencode(queryParams))
	}

	bodyMethods := map[string]bool{"POST": true, "PUT": true, "PATCH": true}
	var body io.Reader = nil
	if bodyMethods[method] {
		body = strings.NewReader(urlencode(bodyParams))
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var resp *Response
	json.Unmarshal(bodyBytes, resp)
	return resp, err
}
