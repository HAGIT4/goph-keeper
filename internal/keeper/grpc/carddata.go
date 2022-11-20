package grpc

import (
	"context"

	sv "github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (sg *serviceGRPC) SaveCardData(ctx context.Context, req *pb.SaveCardDataRequest) (resp *pb.SaveCardDataResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.SaveCardDataReq{
		UserID:        payload.UserID,
		EncKeyword:    req.EncKeyword,
		EncCardNumber: req.EncCardNumber,
		EncCardHolder: req.EncCardHolder,
		EncMeta:       req.EncMeta,
	}
	svResp, err := sg.service.SaveCardData(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.SaveCardDataResponse{
		EncKeyword: svResp.EncKeyword,
	}
	return resp, nil
}

func (sg *serviceGRPC) GetCardData(ctx context.Context, req *pb.GetCardDataRequest) (resp *pb.GetCardDataResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.GetCardDataReq{
		UserID:  payload.UserID,
		Keyword: req.GetEncKeyword(),
	}
	svResp, err := sg.service.GetCardData(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.GetCardDataResponse{
		EncKeyword:    svResp.EncKeyword,
		EncCardNumber: svResp.EncCardNumber,
		EncCardHolder: svResp.EncCardHolder,
		EncCardCode:   svResp.EncCardCode,
		EncMeta:       svResp.EncMeta,
	}
	return resp, nil
}

func (sg *serviceGRPC) ListCardDataKeywords(ctx context.Context, req *pb.ListCardDataKeywordsRequest) (resp *pb.ListCardDataKeywordsResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["token"][0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, err
	}
	svReq := &sv.ListCardDataKeywordsReq{
		UserID: payload.UserID,
	}
	svResp, err := sg.service.ListCardDataKeywords(ctx, svReq)
	if err != nil {
		err = status.Error(codes.Internal, "Internal error")
		return nil, err
	}
	resp = &pb.ListCardDataKeywordsResponse{
		Keywords: svResp.Keywords,
	}
	return resp, nil
}
