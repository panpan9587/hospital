package mongoDB

import (
	"context"
	"demo/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Collection *mongo.Collection

func init() {
	// 从配置文件中获取 MongoDB 连接信息
	username := config.GlobalConfig.Mongodb.Username
	password := config.GlobalConfig.Mongodb.Password
	host := config.GlobalConfig.Mongodb.Host
	port := config.GlobalConfig.Mongodb.Port
	database := config.GlobalConfig.Mongodb.Database
	collectionName := config.GlobalConfig.Mongodb.Collection

	// 构建 MongoDB 连接字符串
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)

	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI(dsn)

	// 设置连接池大小
	clientOptions.SetMaxPoolSize(100) // 设置连接池大小为100

	// 设置超时设置
	clientOptions.SetConnectTimeout(5 * time.Second) // 连接超时时间为5秒
	clientOptions.SetSocketTimeout(30 * time.Second) // Socket超时时间为30秒

	// 连接到 MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("连接MongoDB失败的原因: %v", err)
	}

	// 选择要操作的数据库和集合
	Collection = client.Database(database).Collection(collectionName)

}
