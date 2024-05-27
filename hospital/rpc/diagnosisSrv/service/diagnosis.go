package service

import (
	"context"
	proto "demo/rpc/diagnosisSrv/diagnosisclient"
)

type Diagnosis struct {
	proto.UnimplementedDiagnosisServer
}

func (d Diagnosis) Register(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (d Diagnosis) mustEmbedUnimplementedDiagnosisServer() {
	//TODO implement me
	panic("implement me")
}
