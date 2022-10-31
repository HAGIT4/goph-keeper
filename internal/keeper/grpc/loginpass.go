package grpc

import (
	"context"

	sv "github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc/metadata"
)

func (sg *serviceGRPC) SaveLoginPass(ctx context.Context, req *pb.SaveLoginPassRequest) (resp *pb.SaveLoginPassResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.SaveLoginPassReq{
		UserID:      payload.UserID,
		EncKeyword:  req.EncKeyword,
		EncLogin:    req.EncLogin,
		EncPassword: req.EncPassword,
		EncMeta:     req.EncMeta,
	}
	svResp, err := sg.service.SaveLoginPass(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.SaveLoginPassResponse{
		EncKeyword: svResp.EncKeyword,
	}
	return resp, nil
}

func (sg *serviceGRPC) GetLoginPass(ctx context.Context, req *pb.GetLoginPassRequest) (resp *pb.GetLoginPassResponse, err error) {
	return nil, nil
}
