package storage

import (
	"context"

	"github.com/google/uuid"
)

type KeeperStorageInterface interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (resp *CreateUserResp, err error)
	ReadUser(ctx context.Context, Req *ReadUserReq) (resp *ReadUserResp, err error)
}

var _ KeeperStorageInterface = (*keeperPostgresStorage)(nil)

type CreateUserReq struct {
	Username string
	Passhash string
}
type CreateUserResp struct {
	Username string `db:"username"`
}

type ReadUserReq struct {
	id uuid.UUID
}
type ReadUserResp struct {
	Username string `db:"username"`
	PassHash string `db:"passhash"`
}
