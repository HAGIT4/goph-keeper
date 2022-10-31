package storage

import (
	"context"

	"github.com/google/uuid"
)

type KeeperStorageInterface interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (resp *CreateUserResp, err error)
	ReadUser(ctx context.Context, req *ReadUserReq) (resp *ReadUserResp, err error)
	ReadUserByUsername(ctx context.Context, req *ReadUserByUsernameReq) (resp *ReadUserByUsernameResp, err error)

	CreateLoginPass(ctx context.Context, req *CreateLoginPassReq) (resp *CreateLoginPassResp, err error)
	// GetLoginPass(ctx context.Context, req *GetLoginPassRequest) (resp *GetLoginPassResponse, err error)
}

type User struct {
	UserID   uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Passhash []byte    `db:"passhash"`
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

type LoginPass struct {
	ID       uuid.UUID `db:"uuid"`
	UserID   uuid.UUID `db:"user_id"`
	Keyword  string    `db:"keyword"`
	Login    string    `db:"login"`
	Password []byte    `db:"password"`
	Meta     string    `db:"meta"`
}

type CreateLoginPassReq struct {
	LoginPass
}

type CreateLoginPassResp struct{}

type ReadLoginPassByIDreq struct {
	UserID uuid.UUID `db:"user_id"`
	ID     uuid.UUID `db:"uuid"`
}

type ReadLoginPassByIDresp struct {
	LoginPass
}

type ReadLoginPassByKeywordReq struct {
	UserID  uuid.UUID `db:"user_id"`
	Keyword string    `db:"keyword"`
}

type ReadLoginPassByKeywordResp struct {
	LoginPass
}

type UpdateLoginByIDreq struct {
	UserID uuid.UUID `db:"user_id"`
	ID     uuid.UUID `db:"id"`
	Login  string    `db:"login"`
}

type UpdateLoginByIDresp struct{}
