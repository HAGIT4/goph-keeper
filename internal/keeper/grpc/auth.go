package grpc

import (
	"context"
	"errors"

	"github.com/hagit4/goph-keeper/internal/keeper/service"
	sv "github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sg *serviceGRPC) Register(ctx context.Context, req *pb.RegisterRequest) (resp *pb.RegisterResponse, err error) {
	svReq := &sv.RegisterUserReq{
		Username: req.Username,
		Password: req.Password,
	}
	svResp, err := sg.service.RegisterUser(ctx, svReq)
	if err != nil {
		if errors.Is(err, &sv.ErrorUserNotRegistered{}) {
			err = status.Error(codes.Unknown, "User not created")
			return nil, err
		}
		err = status.Error(codes.Internal, "Internal error")
		return nil, err
	}
	resp = &pb.RegisterResponse{
		Username: svResp.Username,
	}
	return resp, nil
}

func (sg *serviceGRPC) Login(ctx context.Context, req *pb.LoginRequest) (resp *pb.LoginResponse, err error) {
	svReq := &service.LoginUserReq{
		Username: req.Username,
		Password: req.Password,
	}
	svResp, err := sg.service.Login(ctx, svReq)
	if err != nil {
		if errors.Is(err, &sv.ErrorUnauthenticated{}) {
			err = status.Error(codes.Unauthenticated, "Access denied")
			return nil, err
		}
		err = status.Error(codes.Internal, "Internal error")
		return nil, err
	}
	resp = &pb.LoginResponse{
		AccessToken: svResp.AccessToken,
	}
	return resp, nil
}
