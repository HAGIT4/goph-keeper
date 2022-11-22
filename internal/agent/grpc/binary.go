package grcp

import (
	"context"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (ag *agentGRPC) SaveBinary(ctx context.Context, req *pb.SaveBinaryRequest) (resp *pb.SaveBinaryResponse, err error) {
	resp, err = ag.binaryClient.SaveBinary(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) GetBinary(ctx context.Context, req *pb.GetBinaryRequest) (resp *pb.GetBinaryResponse, err error) {
	resp, err = ag.binaryClient.GetBinary(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) ListBinaryKeywords(ctx context.Context, req *pb.ListBinaryKeywordsRequest) (resp *pb.ListBinaryKeywordsResponse, err error) {
	resp, err = ag.binaryClient.ListBinaryKeywords(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
