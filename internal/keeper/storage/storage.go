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
	ReadLoginPassByID(ctx context.Context, req *ReadLoginPassByIDreq) (resp *ReadLoginPassByIDresp, err error)
	ReadLoginPassByKeyword(ctx context.Context, req *ReadLoginPassByKeywordReq) (resp *ReadLoginPassByKeywordResp, err error)
	ListLoginPassKeywords(ctx context.Context, req *ListLoginPassKeywordsReq) (resp *ListLoginPassKeywordsResp, err error)
}

type User struct {
	UserID   uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Passhash []byte    `db:"passhash"`
}

type LoginPass struct {
	ID       uuid.UUID `db:"uuid"`
	UserID   uuid.UUID `db:"user_id"`
	Keyword  string    `db:"keyword"`
	Login    string    `db:"login"`
	Password []byte    `db:"password"`
	Meta     string    `db:"meta"`
}

//Create User
type CreateUserReq struct {
	User
}
type CreateUserResp struct {
	Username string `db:"username"`
}

//Read User by UUID
type ReadUserReq struct {
	ID uuid.UUID
}
type ReadUserResp struct {
	User
}

//Read User By Username
type ReadUserByUsernameReq struct {
	Username string
}

type ReadUserByUsernameResp struct {
	User
}

//Create LoginPass
type CreateLoginPassReq struct {
	LoginPass
}

type CreateLoginPassResp struct{}

//Read LoginPass by ID
type ReadLoginPassByIDreq struct {
	UserID uuid.UUID `db:"user_id"`
	ID     uuid.UUID `db:"uuid"`
}

type ReadLoginPassByIDresp struct {
	LoginPass
}

//Read LoginPass by Keyword
type ReadLoginPassByKeywordReq struct {
	UserID  uuid.UUID `db:"user_id"`
	Keyword string    `db:"keyword"`
}

type ReadLoginPassByKeywordResp struct {
	LoginPass
}

//Update LoginPass
type UpdateLoginByIDreq struct {
	UserID uuid.UUID `db:"user_id"`
	ID     uuid.UUID `db:"id"`
	Login  string    `db:"login"`
}

type UpdateLoginByIDresp struct{}

//List LoginPass
type ListLoginPassKeywordsReq struct {
	UserID uuid.UUID
}

type ListLoginPassKeywordsResp struct {
	Keywords []string
}
