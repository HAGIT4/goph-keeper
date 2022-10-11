package service

import (
	"context"

	"github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type KeeperServiceInterface interface {
	RegisterUser(ctx context.Context, req *RegisterUserReq) (resp *RegisterUserResp, err error)
}

type keeperService struct {
	storage storage.KeeperStorageInterface
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
