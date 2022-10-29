package grpc

import (
	"context"
	"fmt"

	sv "github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (sg *serviceGRPC) SaveLoginPass(ctx context.Context, req *pb.SaveLoginPassRequest) (resp *pb.SaveLoginPassResponse, err error) {
	username := ctx.Value(sv.KeyUsername).(string)
	if len(username) == 0 {
		return nil, fmt.Errorf("no username")
	}
	svReq := &sv.SaveLoginPassReq{
		Username:    "username",
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
