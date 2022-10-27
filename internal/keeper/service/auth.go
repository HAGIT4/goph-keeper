package service

import (
	"bytes"
	"context"

	"crypto/rand"

	"github.com/hagit4/goph-keeper/internal/keeper/storage"
	"golang.org/x/crypto/argon2"
)

type RegisterUserReq struct {
	Username string
	Password string
}
type RegisterUserResp struct {
	Username string
}

func (ks *keeperService) RegisterUser(ctx context.Context, req *RegisterUserReq) (resp *RegisterUserResp, err error) {
	salt := make([]byte, 8)
	if _, err = rand.Read(salt); err != nil {
		return nil, err
	}
	passHash := hashPassword(req.Password, salt)
	stReq := &storage.CreateUserReq{
		User: storage.User{
			Username: req.Username,
			Passhash: string(passHash),
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

	if !checkPassword([]byte(stResp.Passhash), req.Password) {
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

func hashPassword(plainPass string, salt []byte) (saltedHash []byte) {
	rand.Read(salt)
	hashedPass := argon2.IDKey([]byte(plainPass), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func checkPassword(storedSaltedPassHash []byte, plainPass string) bool {
	salt := storedSaltedPassHash[0:8]
	userSaltedPassHash := hashPassword(plainPass, salt)
	return bytes.Equal(userSaltedPassHash, storedSaltedPassHash)
}

func (ks *keeperService) VerifyAuthToken(token string) (payload *AuthTokenPayload, err error) {
	return ks.tokenMaker.VerifyAuthToken(token)
}
