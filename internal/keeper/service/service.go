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
	GetLoginPass(ctx context.Context, req *GetLoginPassReq) (resp *GetLoginPassResp, err error)

	SaveTextData(ctx context.Context, req *SaveTextDataReq) (resp *SaveLoginPassResp, err error)
	ReadTextDataByID(ctx context.Context, req *ReadTextDataByIDreq) (resp *ReadTextDataByIDresp, err error)
	ReadTextDataByKeyword(ctx context.Context, req *ReadTextDataByKeywordReq) (resp *ReadTextDataByKeywordResp, err error)
	ListTextDataKeywords(ctx context.Context, req *ListTextDataKeywordsReq) (resp *ListTextDataKeywordsResp, err error)

	SaveCardData(ctx context.Context, req *SaveCardDataReq) (resp *SaveCardDataResp, err error)
	GetCardData(ctx context.Context, req *GetCardDataReq) (resp *GetCardDataResp, err error)
	ListCardDataKeywords(ctx context.Context, req *ListCardDataKeywordsReq) (resp *ListCardDataKeywordsResp, err error)
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
