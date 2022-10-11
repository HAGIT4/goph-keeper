package storage

import (
	"context"
	"fmt"
)

type keeperPostgresStorage struct{}

var _ KeeperStorageInterface = (*keeperPostgresStorage)(nil)

type keeperPostgresStorageOption func(*keeperPostgresStorage)

func NewKeeperPostgresStorage(opts ...keeperPostgresStorageOption) (storage *keeperPostgresStorage) {
	storage = &keeperPostgresStorage{}
	for _, opt := range opts {
		opt(storage)
	}
	return storage
}

func (ps *keeperPostgresStorage) CreateUser(ctx context.Context, req *CreateUserReq) (resp *CreateUserResp, err error) {
	fmt.Println("User created in db")
	resp = &CreateUserResp{}
	return resp, nil
}
