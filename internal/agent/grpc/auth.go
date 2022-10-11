package grcp

import (
	"context"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (ag *agentGRPC) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (resp *pb.RegisterResponse, err error) {
	resp, err = ag.authClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) Login(ctx context.Context, req *pb.LoginRequest) (resp *pb.LoginResponse, err error) {
	resp, err = ag.authClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
