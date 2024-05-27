package service

import (
	"context"
	proto "demo/rpc/advisorySrv/advisoryclient"
	"demo/rpc/model/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Advisory struct {
	proto.UnimplementedAdvisoryServer
}

// 记录历史咨询消息
func (a Advisory) RecordsUserConsultationInformation(ctx context.Context, in *proto.RecordsUserConsultationInformationReq) (*proto.RecordsUserConsultationInformationRes, error) {
	//创建要插入的文档
	document := bson.M{
		"content": in.Content,
		"result":  in.Result,
	}
	//在集合中插入文档
	result, err := mongoDB.Collection.InsertOne(ctx, document)
	if err != nil {
		log.Printf("在集合中插入文档失败的原因: %v", err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &proto.RecordsUserConsultationInformationRes{
		DocumentId: result.InsertedID.(primitive.ObjectID).Hex(), // 将插入的文档ID包含在响应中
	}, nil

}
