package grpc

import (
	"context"

	sv "github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc/metadata"
)

func (sg *serviceGRPC) SaveBinary(ctx context.Context, req *pb.SaveBinaryRequest) (resp *pb.SaveBinaryResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.SaveBinaryReq{
		UserID:     payload.UserID,
		EncKeyword: req.Keyword,
		EncBin:     req.Bin,
		EncMeta:    req.Meta,
	}
	svResp, err := sg.service.SaveBinary(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.SaveBinaryResponse{
		Keyword: svResp.EncKeyword,
	}
	return resp, nil
}

func (sg *serviceGRPC) GetBinary(ctx context.Context, req *pb.GetBinaryRequest) (resp *pb.GetBinaryResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.GetBinaryReq{
		UserID:  payload.UserID,
		Keyword: req.GetKeyword(),
	}
	svResp, err := sg.service.GetBinary(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.GetBinaryResponse{
		Keyword: svResp.Keyword,
		Bin:     svResp.Bin,
		Meta:    svResp.Meta,
	}
	return resp, nil
}

func (sg *serviceGRPC) ListBinaryKeywords(ctx context.Context, req *pb.ListBinaryKeywordsRequest) (resp *pb.ListBinaryKeywordsResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.ListBinaryKeywordsReq{
		UserID: payload.UserID,
	}
	svResp, err := sg.service.ListBinaryKeywords(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.ListBinaryKeywordsResponse{
		Keywords: svResp.Keywords,
	}
	return resp, nil
}
