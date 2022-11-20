package grcp

import (
	"context"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (ag *agentGRPC) SaveCardData(ctx context.Context, req *pb.SaveCardDataRequest) (resp *pb.SaveCardDataResponse, err error) {
	resp, err = ag.cardDataClient.SaveCardData(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) GetCardData(ctx context.Context, req *pb.GetCardDataRequest) (resp *pb.GetCardDataResponse, err error) {
	resp, err = ag.cardDataClient.GetCardData(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) ListCardDataKeywords(ctx context.Context, req *pb.ListCardDataKeywordsRequest) (resp *pb.ListCardDataKeywordsResponse, err error) {
	resp, err = ag.cardDataClient.ListCardDataKeywords(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
