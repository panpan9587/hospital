package common

/*
	info(
	desc: "验证码的封装"
	author: "panpan"
	email: "918716975@qq.com"
)
*/
import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// GenerateFourDigitRandomNumber 生成一个四位随机数
func GenerateFourDigitRandomNumber() (string, error) {
	// 定义四位数的范围，从1000到9999
	Min := int64(1000)
	Max := int64(9999)

	// 生成范围内的随机数
	number, err := rand.Int(rand.Reader, big.NewInt(Max-Min+1))
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(number.Int64() + Min)), nil

}
