package postgres

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbString = "postgresql://localhost:5432/keeper?user=keeper_user&password=password&sslmode=disable"
var postgres *keeperPostgresStorage

func TestMain(m *testing.M) {
	var err error
	postgres, err = initDB()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
}

func initDB() (st *keeperPostgresStorage, err error) {
	db, err := sqlx.Connect("postgres", dbString)
	if err != nil {
		return nil, err
	}
	st = NewKeeperPostgresStorage(
		WithConnection(db),
	)
	return st, nil
}

func TestGetNotExistingUser(t *testing.T) {
	ctx := context.Background()
	stReq := &st.ReadUserByUsernameReq{
		Username: "SomeGuy",
	}
	resp, err := postgres.ReadUserByUsername(ctx, stReq)
	if !errors.Is(err, &st.ErrorUsernameNotFound{}) {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGetExistingUser(t *testing.T) {
	ctx := context.Background()
	stCreateReq := &st.CreateUserReq{
		User: st.User{
			ID:       uuid.New(),
			Username: "NewUser2",
			Passhash: "pshs",
		},
	}
	_, err := postgres.CreateUser(ctx, stCreateReq)
	if err != nil {
		t.Fatal(err)
	}

	ctx = context.Background()
	stReadReq := &st.ReadUserByUsernameReq{
		Username: "NewUser2",
	}
	stReadResp, err := postgres.ReadUserByUsername(ctx, stReadReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(stReadResp)
}
