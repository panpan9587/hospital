package redis

import (
	"context"
	"time"
)

// 存入验证码
func StorageCode(code, mobile string) {
	Client.Set(context.Background(), mobile, code, time.Second*60)
}

// 获取验证码
func GetCode(mobile string) string {
	code, _ := Client.Get(context.Background(), mobile).Result()
	return code
}

// 删除验证码
