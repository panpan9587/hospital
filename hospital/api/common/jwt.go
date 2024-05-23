package common

/*
	info(
	desc: "封装的token调用"
	author: "panpan"
	email: "918716975@qq.com"
)
*/
import (
	"demo/api/etc"
	"demo/rpc/userSrv/userclient"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体

func GetJwtToken(payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + etc.ApiConfig.Token.AccessExpire
	claims["iat"] = time.Now().Unix()
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(etc.ApiConfig.Token.AccessSecret))
}

// 验证token,反解返回用户信息

func ParseToken(tokenString string) (user *proto.UserInfo, err error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		hmacSampleSecret := []byte(etc.ApiConfig.Token.AccessSecret)
		return hmacSampleSecret, nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return nil, errors.New("令牌过期")
	}
	data := []byte(claims["payload"].(string))
	err = json.Unmarshal(data, &user)
	return
}
