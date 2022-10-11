package service

import (
	"context"

	"github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type RegisterUserReq struct{}
type RegisterUserResp struct{}

func (ks *keeperService) RegisterUser(ctx context.Context, req *RegisterUserReq) (resp *RegisterUserResp, err error) {
	stReq := &storage.CreateUserReq{}
	ks.storage.CreateUser(ctx, stReq)
	resp = &RegisterUserResp{}
	return resp, nil
}
