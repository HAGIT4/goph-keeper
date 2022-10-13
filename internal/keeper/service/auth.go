package service

import (
	"context"

	"github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type RegisterUserReq struct {
	Username string
	PassHash string
}
type RegisterUserResp struct {
	Username string
}

func (ks *keeperService) RegisterUser(ctx context.Context, req *RegisterUserReq) (resp *RegisterUserResp, err error) {
	stReq := &storage.CreateUserReq{
		Username: req.Username,
		Passhash: req.PassHash,
	}
	dbResp, err := ks.storage.CreateUser(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &RegisterUserResp{
		Username: dbResp.Username,
	}
	return resp, nil
}
