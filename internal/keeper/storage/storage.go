package storage

import (
	"context"

	"github.com/google/uuid"
)

type KeeperStorageInterface interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (resp *CreateUserResp, err error)
	ReadUser(ctx context.Context, req *ReadUserReq) (resp *ReadUserResp, err error)
	ReadUserByUsername(ctx context.Context, req *ReadUserByUsernameReq) (resp *ReadUserByUsernameResp, err error)
	//LoginPass
	CreateLoginPass(ctx context.Context, req *CreateLoginPassReq) (resp *CreateLoginPassResp, err error)
	ReadLoginPassByID(ctx context.Context, req *ReadLoginPassByIDreq) (resp *ReadLoginPassByIDresp, err error)
	ReadLoginPassByKeyword(ctx context.Context, req *ReadLoginPassByKeywordReq) (resp *ReadLoginPassByKeywordResp, err error)
	ListLoginPassKeywords(ctx context.Context, req *ListLoginPassKeywordsReq) (resp *ListLoginPassKeywordsResp, err error)
	//TextData
	CreateTextData(ctx context.Context, req *CreateTextDataReq) (resp *CreateTextDataResp, err error)
	ReadTextData(ctx context.Context, req *ReadTextDataReq) (resp *ReadTextDataResp, err error)
	ReadTextDataByKeyword(ctx context.Context, req *ReadTextDataByKeywordReq) (resp *ReadTextDataByKeywordResp, err error)
	UpdateTextData(ctx context.Context, req *UpdateTextDataReq) (resp *UpdateTextDataResp, err error)
	DeleteTextData(ctx context.Context, req *DeleteTextDataReq) (resp *DeleteTextDataResp, err error)
	ListTextDataKeywords(ctx context.Context, req *ListTextDataKeywordsReq) (resp *ListTextDataKeywordsResp, err error)
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

type TextData struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	Keyword  string    `db:"keyword"`
	TextData string    `db:"textdata"`
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

//Text structures

//Create TextData
type CreateTextDataReq struct {
	TextData TextData
}

type CreateTextDataResp struct{}

//Read TextData
type ReadTextDataReq struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

type ReadTextDataResp struct {
	TextData TextData
}

type ReadTextDataByKeywordReq struct {
	UserId  uuid.UUID
	Keyword string
}

type ReadTextDataByKeywordResp struct {
	TextData TextData
}

//Update TextData
type UpdateTextDataReq struct {
	TextData TextData
}

type UpdateTextDataResp struct{}

//Delete TextData
type DeleteTextDataReq struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

type DeleteTextDataResp struct{}

//List TextData
type ListTextDataKeywordsReq struct {
	UserID uuid.UUID
}

type ListTextDataKeywordsResp struct {
	Keywords []string
}
