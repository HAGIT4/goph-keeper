package service

import (
	"context"

	"github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type key int

const (
	KeyUsername key = iota
)

type KeeperServiceInterface interface {
	RegisterUser(ctx context.Context, req *RegisterUserReq) (resp *RegisterUserResp, err error)
	Login(ctx context.Context, req *LoginUserReq) (resp *LoginUserResp, err error)
	VerifyAuthToken(token string) (payload *AuthTokenPayload, err error)

	SaveLoginPass(ctx context.Context, req *SaveLoginPassReq) (resp *SaveLoginPassResp, err error)
	ListLoginPassKeywords(ctx context.Context, req *ListLoginPassKeywordsReq) (resp *ListLoginPassKeywordsResp, err error)
}

type keeperService struct {
	storage    storage.KeeperStorageInterface
	tokenMaker TokenMakerInterface
}

var _ KeeperServiceInterface = (*keeperService)(nil)

type keeperServiceOption func(*keeperService)

func NewKeeperService(opts ...keeperServiceOption) (service *keeperService) {

	service = &keeperService{}
	for _, opt := range opts {
		opt(service)
	}
	return service
}

func WithStorage(st storage.KeeperStorageInterface) keeperServiceOption {
	return func(ks *keeperService) {
		ks.storage = st
	}
}

func WithTokenMaker(tm TokenMakerInterface) keeperServiceOption {
	return func(ks *keeperService) {
		ks.tokenMaker = tm
	}
}
