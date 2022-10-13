package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type keeperPostgresStorage struct {
	db *sqlx.DB
}

var _ KeeperStorageInterface = (*keeperPostgresStorage)(nil)

type keeperPostgresStorageOption func(*keeperPostgresStorage)

func NewKeeperPostgresStorage(opts ...keeperPostgresStorageOption) (storage *keeperPostgresStorage) {
	storage = &keeperPostgresStorage{}
	for _, opt := range opts {
		opt(storage)
	}
	return storage
}

func WithConnection(db *sqlx.DB) keeperPostgresStorageOption {
	return func(kps *keeperPostgresStorage) {
		kps.db = db
	}
}

func (ps *keeperPostgresStorage) CreateUser(ctx context.Context, req *CreateUserReq) (resp *CreateUserResp, err error) {
	resp = &CreateUserResp{}
	userId := uuid.New()
	query := `INSERT INTO keeper.users (id, username, passhash) VALUES ($1, $2, $3)`
	_, err = ps.db.ExecContext(ctx, query, userId, req.Username, req.Passhash)
	if err != nil {
		return nil, err
	}
	resp.Username = req.Username
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadUser(ctx context.Context, req *ReadUserReq) (resp *ReadUserResp, err error) {
	resp = &ReadUserResp{}
	if err = ps.db.Get(resp, "SELECT * FROM keeper.users WHERE id=$1", req.id); err != nil {
		return nil, err
	}
	return resp, nil
}
