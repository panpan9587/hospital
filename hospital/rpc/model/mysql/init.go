package mysql

import (
	"demo/config"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	fmt.Println(config.GlobalConfig, "尚鑫平")
	// todo:初始化连接配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.Mysql.Username,
		config.GlobalConfig.Mysql.Password,
		config.GlobalConfig.Mysql.Host,
		config.GlobalConfig.Mysql.Port,
		config.GlobalConfig.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Info("数据库连接失败", err)
	}
	DB = db
	sql, _ := db.DB()
	sql.SetMaxIdleConns(config.GlobalConfig.Mysql.MaxIdleConn)
	sql.SetMaxOpenConns(config.GlobalConfig.Mysql.MaxOpenConn)
}
