package grcp

import (
	"context"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (ag *agentGRPC) SaveTextData(ctx context.Context, req *pb.SaveTextDataRequest) (resp *pb.SaveTextDataResponse, err error) {
	resp, err = ag.textdataClient.SaveTextData(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) GetTextData(ctx context.Context, req *pb.GetTextDataRequest) (resp *pb.GetTextDataResponse, err error) {
	resp, err = ag.textdataClient.GetTextData(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) ListTextDataKeywords(ctx context.Context, req *pb.ListTextDataKeywordsRequest) (resp *pb.ListTextDataKeywordsResponse, err error) {
	resp, err = ag.textdataClient.ListTextDataKeywords(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
