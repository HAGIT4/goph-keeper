package service

import (
	"context"

	"github.com/hagit4/goph-keeper/internal/keeper/storage"
	"golang.org/x/crypto/bcrypt"
)

var (
	bcryptCost = bcrypt.DefaultCost
)

type RegisterUserReq struct {
	Username string
	Password string
}
type RegisterUserResp struct {
	Username string
}

func (ks *keeperService) RegisterUser(ctx context.Context, req *RegisterUserReq) (resp *RegisterUserResp, err error) {
	passHash, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	stReq := &storage.CreateUserReq{
		User: storage.User{
			Username: req.Username,
			Passhash: passHash,
		},
	}
	stResp, err := ks.storage.CreateUser(ctx, stReq)
	if err != nil {
		return nil, &ErrorUserNotRegistered{}
	}
	resp = &RegisterUserResp{
		Username: stResp.Username,
	}
	return resp, nil
}

type LoginUserReq struct {
	Username string
	Password string
}
type LoginUserResp struct {
	AccessToken string
}

func (ks *keeperService) Login(ctx context.Context, req *LoginUserReq) (resp *LoginUserResp, err error) {
	stReq := &storage.ReadUserByUsernameReq{
		Username: req.Username,
	}
	stResp, err := ks.storage.ReadUserByUsername(ctx, stReq)
	if err != nil {
		return nil, &ErrorUnauthenticated{}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(stResp.User.Passhash), []byte(req.Password)); err != nil {
		return nil, &ErrorUnauthenticated{}
	}
	token, err := ks.tokenMaker.CreateAuthToken(req.Username)
	if err != nil {
		return nil, &ErrorInternal{}
	}
	resp = &LoginUserResp{
		AccessToken: token,
	}
	return resp, nil
}

func hashPassword(password string) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}
