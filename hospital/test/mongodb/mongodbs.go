package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	// 设置 MongoDB 连接信息，包括用户名和密码
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@124.223.203.56:27017")

	// 连接到 MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 确保在程序结束时关闭数据库连接
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// 选择要操作的数据库和集合
	collection := client.Database("information").Collection("user")

	// 创建要插入的文档
	document := bson.M{
		"content": "你上班了吗？",
		"result":  "多喝水休息",
	}

	// 在集合中插入文档
	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功插入文档")
}
