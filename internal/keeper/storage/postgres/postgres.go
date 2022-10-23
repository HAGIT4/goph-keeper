package postgres

import (
	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
	"github.com/jmoiron/sqlx"
)

type keeperPostgresStorage struct {
	db *sqlx.DB
}

var _ st.KeeperStorageInterface = (*keeperPostgresStorage)(nil)

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
