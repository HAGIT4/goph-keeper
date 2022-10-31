package grcp

import (
	"context"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (ag *agentGRPC) SaveLoginPass(ctx context.Context, req *pb.SaveLoginPassRequest) (resp *pb.SaveLoginPassResponse, err error) {
	resp, err = ag.loginPassClient.SaveLoginPass(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) GetLoginPass(ctx context.Context, req *pb.GetLoginPassRequest) (resp *pb.GetLoginPassResponse, err error) {
	resp, err = ag.loginPassClient.GetLoginPass(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) ListLoginPassKeywords(ctx context.Context, req *pb.ListLoginPassKeywordsRequest) (resp *pb.ListLoginPassKeywordsResponse, err error) {
	resp, err = ag.loginPassClient.ListLoginPassKeywords(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
