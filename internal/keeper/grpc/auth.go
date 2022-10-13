package grpc

import (
	"context"

	"github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (sg *serviceGRPC) Register(ctx context.Context, req *pb.RegisterRequest) (resp *pb.RegisterResponse, err error) {
	svReq := &service.RegisterUserReq{
		Username: req.Username,
		PassHash: req.PassHash,
	}
	svResp, err := sg.service.RegisterUser(ctx, svReq)
	if err != nil {
		return nil, err
	}
	resp = &pb.RegisterResponse{
		Username: svResp.Username,
	}
	return resp, nil
}
