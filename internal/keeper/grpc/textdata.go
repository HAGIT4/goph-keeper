package grpc

import (
	"context"

	sv "github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (sg *serviceGRPC) SaveTextData(ctx context.Context, req *pb.SaveTextDataRequest) (resp *pb.SaveTextDataResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	tokens, ok := md["token"]
	if !ok || len(tokens) == 0 {
		return nil, status.Error(codes.PermissionDenied, "No token found")
	}
	token := tokens[0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	svReq := &sv.SaveTextDataReq{
		UserID:      payload.UserID,
		EncKeyword:  req.EncKeyword,
		EncTextData: req.EncTextData,
		EncMeta:     req.EncMeta,
	}
	svResp, err := sg.service.SaveTextData(ctx, svReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	resp = &pb.SaveTextDataResponse{
		EncKeyword: svResp.EncKeyword,
	}
	return resp, nil
}

func (sg *serviceGRPC) GetTextData(ctx context.Context, req *pb.GetTextDataRequest) (resp *pb.GetTextDataResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	tokens, ok := md["token"]
	if !ok || len(tokens) == 0 {
		return nil, status.Error(codes.PermissionDenied, "No token found")
	}
	token := tokens[0]
	payload, err := sg.service.VerifyAuthToken(token)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	svReq := &sv.ReadTextDataByKeywordReq{
		UserID:  payload.UserID,
		Keyword: req.EncKeyword,
	}
	svResp, err := sg.service.ReadTextDataByKeyword(ctx, svReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	resp = &pb.GetTextDataResponse{
		EncKeyword:  svResp.Keyword,
		EncTextData: svResp.TextData,
		EncMeta:     svResp.Meta,
	}
	return resp, nil
}
