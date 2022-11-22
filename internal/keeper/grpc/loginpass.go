package grpc

import (
	"context"

	sv "github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	// TODO: middleware for that
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.GetLoginPassReq{
		UserID:  payload.UserID,
		Keyword: req.GetEncKeyword(),
	}
	svResp, err := sg.service.GetLoginPass(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.GetLoginPassResponse{
		EncKeyword:  svResp.Keyword,
		EncLogin:    svResp.Login,
		EncPassword: svResp.Password,
		EncMeta:     svResp.Meta,
	}
	return resp, nil
}

func (sg *serviceGRPC) ListLoginPassKeywords(ctx context.Context, req *pb.ListLoginPassKeywordsRequest) (resp *pb.ListLoginPassKeywordsResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.ListLoginPassKeywordsReq{
		UserID: payload.UserID,
	}
	svResp, err := sg.service.ListLoginPassKeywords(ctx, svReq)
	if err != nil {
		err = status.Error(codes.Internal, "Internal error")
		return nil, err
	}
	resp = &pb.ListLoginPassKeywordsResponse{
		Keywords: svResp.Keywords,
	}
	return resp, nil
}
