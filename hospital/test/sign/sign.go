package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserSign struct {
	rdb *redis.Client
}

func NewUserSign(opt *redis.Options) *UserSign {
	return &UserSign{rdb: redis.NewClient(opt)}
}

func (us *UserSign) Sign(userID int64, date time.Time) error {
	key := fmt.Sprintf("user:%d:sign:%d-%02d-%02d", userID, date.Year(), date.Month(), date.Day())
	_, err := us.rdb.SetNX(context.Background(), key, 1, 0).Result()
	return err
}

func (us *UserSign) GetMonthSignCount(userID int64, year int, month time.Month) (int, error) {
	signCount := 0
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	nextMonth := startOfMonth.AddDate(0, 1, 0)
	for currentDate := startOfMonth; currentDate.Before(nextMonth); currentDate = currentDate.AddDate(0, 0, 1) {
		key := fmt.Sprintf("user:%d:sign:%d-%02d-%02d", userID, currentDate.Year(), currentDate.Month(), currentDate.Day())
		exists, err := us.rdb.Exists(context.Background(), key).Result()
		if err != nil {
			return 0, err
		}
		if exists > 0 {
			signCount++
		}
	}
	return signCount, nil
}

// GetUserMonthlySignDetails 获取用户在指定年份内每个月的签到次数及详细情况
func (us *UserSign) GetUserMonthlySignDetails(userID int64, year int) (map[string][]string, map[string]int, error) {
	details := make(map[string][]string)
	counts := make(map[string]int)
	for month := time.January; month <= time.December; month++ {
		mo := int(month)
		monthStr := fmt.Sprintf("%d-%02d", year, mo)
		daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
		signDates := []string{}
		for day := 1; day <= daysInMonth; day++ {
			date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			key := fmt.Sprintf("user:%d:sign:%d-%02d-%02d", userID, date.Year(), date.Month(), date.Day())
			exists, err := us.rdb.Exists(context.Background(), key).Result()
			if err != nil {
				return nil, nil, err
			}
			if exists > 0 {
				signDates = append(signDates, date.Format("2006-01-02"))
			}
		}
		details[monthStr] = signDates
		counts[monthStr] = len(signDates)
	}
	return details, counts, nil
}

// GetSignDetailsInRange 获取用户在指定日期范围内的签到详细情况
func (us *UserSign) GetSignDetailsInRange(userID int64, startDate, endDate time.Time) ([]string, error) {
	var signDates []string
	currentDate := startDate
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		key := fmt.Sprintf("user:%d:sign:%d-%02d-%02d", userID, currentDate.Year(), currentDate.Month(), currentDate.Day())
		exists, err := us.rdb.Exists(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		if exists > 0 {
			signDates = append(signDates, currentDate.Format("2006-01-02"))
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return signDates, nil
}

func (us *UserSign) CompensateSign(userID int64, date time.Time) error {
	key := fmt.Sprintf("user:%d:sign:%d-%02d-%02d", userID, date.Year(), date.Month(), date.Day())
	_, err := us.rdb.SetNX(context.Background(), key, 1, 0).Result()
	return err
}

func main() {
	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	userSign := NewUserSign(opt)

	// 示例：用户1在今天的日期签到
	now := time.Now()
	err = userSign.Sign(1, now)
	if err != nil {
		fmt.Println("签到失败:", err)
	} else {
		fmt.Println("签到成功")
	}

	// 示例：获取用户1本月签到次数
	count, err := userSign.GetMonthSignCount(1, now.Year(), now.Month())
	if err != nil {
		fmt.Println("获取签到次数失败:", err)
	} else {
		fmt.Printf("本月签到次数: %d\n", count)
	}

	// 示例：为用户1在指定日期（例如昨天）补签
	yesterday := now.AddDate(0, 0, -1)
	err = userSign.CompensateSign(1, yesterday)
	if err != nil {
		fmt.Println("补签失败:", err)
	} else {
		fmt.Println("补签成功")
	}

	yearToCheck := time.Now().Year() // 查询本年度的签到情况，可以根据需要替换为其他年份

	// 示例：获取用户1本年度每个月的签到详细情况和次数
	details, counts, err := userSign.GetUserMonthlySignDetails(1, yearToCheck)
	if err != nil {
		fmt.Println("获取每月签到详情失败:", err)
	} else {
		for month, dates := range details {
			fmt.Printf("月份: %s, 签到次数: %d, 详细情况: %v\n", month, counts[month], dates)
		}
	}

	startDate := time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 5, 31, 23, 59, 59, 0, time.UTC) // 注意结束时间通常设置为月末的最后一秒
	signDetails, err := userSign.GetSignDetailsInRange(1, startDate, endDate)
	if err != nil {
		fmt.Println("获取签到详情失败:", err)
	} else {
		fmt.Printf("指定日期范围内的签到详细情况: %v\n", signDetails)
	}
}
