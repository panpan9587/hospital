package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 同步
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

	// 创建查询过滤器，这里查询所有文档
	filter := bson.M{} // 空 filter 表示查询所有文档
	//包含_id字段在投影中
	projection := bson.M{
		"_id":     1,
		"content": 1,
		"result":  1,
	}
	// 执行查询操作
	cursor, err := collection.Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// 初始化 Elasticsearch 客户端
	esClient, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9201"), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	// 遍历查询结果并索引到 Elasticsearch
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		// 获取文档的 _id 字段值，并将其转换为字符串类型
		id := result["_id"].(primitive.ObjectID).Hex()

		// 移除原始的 _id 字段
		delete(result, "_id")

		// 索引到 Elasticsearch，并指定字符串类型的 _id 字段
		_, err := esClient.Index().
			Index("users").
			Id(id). // 指定字符串类型的 _id 字段
			BodyJson(result).
			Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("文档内容:", result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功查询文档列表并将数据同步到 Elasticsearch")
}
