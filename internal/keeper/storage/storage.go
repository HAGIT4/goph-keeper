package storage

import "context"

type KeeperStorageInterface interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (resp *CreateUserResp, err error)
}

var _ KeeperStorageInterface = (*keeperPostgresStorage)(nil)

type CreateUserReq struct{}
type CreateUserResp struct{}
