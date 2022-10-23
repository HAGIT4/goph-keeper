package storage

import (
	"context"

	"github.com/google/uuid"
)

type KeeperStorageInterface interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (resp *CreateUserResp, err error)
	ReadUser(ctx context.Context, req *ReadUserReq) (resp *ReadUserResp, err error)
	ReadUserByUsername(ctx context.Context, req *ReadUserByUsernameReq) (resp *ReadUserByUsernameResp, err error)
}

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Passhash string    `db:"passhash"`
}

type CreateUserReq struct {
	User
}
type CreateUserResp struct {
	Username string `db:"username"`
}

type ReadUserReq struct {
	ID uuid.UUID
}
type ReadUserResp struct {
	User
}

type ReadUserByUsernameReq struct {
	Username string
}

type ReadUserByUsernameResp struct {
	User
}
