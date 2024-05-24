package service

import (
	"context"
	proto "demo/rpc/caseSrv/caseclient"
	"demo/rpc/model/esticsearch"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"strings"
)

type Case struct {
	proto.UnimplementedCaseServer
}

// 病历记录
func (c Case) CaseRecordList(ctx context.Context, in *proto.CaseRecordListReq) (*proto.CaseRecordListRes, error) {
	//接收参数
	index := in.Index
	query := elastic.NewMatchAllQuery()
	page := in.Page
	//调用查询所有的方法
	_, _, status, err := esticsearch.QueryEs(index, query, int(page))
	if err != nil {
		zap.S().Info("查询失败的原因:::", err.Error())
		return nil, err
	}
	var casesList []*proto.CaseRecord
	for _, v := range status {
		cases := &proto.CaseRecord{
			Content: v["content"].(string),
			Result:  v["result"].(string),
		}
		casesList = append(casesList, cases)
	}

	return &proto.CaseRecordListRes{
		Result: casesList,
	}, nil
}

// 搜索病历记录
func (c Case) SearchCaseRecord(ctx context.Context, in *proto.SearchCaseRecordReq) (*proto.SearchCaseRecordRes, error) {
	index := in.Index
	content := in.Content

	// 调用单条件查询的方法
	status, err := esticsearch.FilterSearch(index, content)
	if err != nil {
		zap.S().Info("查询失败的原因:::", err.Error())
		return nil, err
	}

	// 初始化结果字符串
	var res strings.Builder

	// 遍历查询结果，将每个结果连接成一个字符串
	for _, v := range status {
		result := v["result"].(string)
		res.WriteString(result)
		res.WriteString(" ") // 可能需要分隔符，这里假设用空格分隔每个结果
	}

	// 返回结果对象
	return &proto.SearchCaseRecordRes{
		Result: res.String(),
	}, nil
}
