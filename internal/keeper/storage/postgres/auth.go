package postgres

import (
	"context"

	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
)

func (ps *keeperPostgresStorage) CreateUser(ctx context.Context, req *st.CreateUserReq) (resp *st.CreateUserResp, err error) {
	resp = &st.CreateUserResp{}
	query := `INSERT INTO keeper.user (id, username, passhash) VALUES ($1, $2, $3)`
	_, err = ps.db.ExecContext(ctx, query, req.UserID, req.Username, req.Passhash)
	if err != nil {
		return nil, &st.ErrorUserNotCreated{}
	}
	resp.Username = req.Username
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadUser(ctx context.Context, req *st.ReadUserReq) (resp *st.ReadUserResp, err error) {
	resp = &st.ReadUserResp{}
	if err = ps.db.Get(resp, "SELECT * FROM keeper.user WHERE id=$1 LIMIT 1", req.ID); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, &st.ErrorUserIDNotFound{
				ID: req.ID,
			}
		}
		return nil, err
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadUserByUsername(ctx context.Context, req *st.ReadUserByUsernameReq) (resp *st.ReadUserByUsernameResp, err error) {
	resp = &st.ReadUserByUsernameResp{}
	if err = ps.db.Get(resp, "SELECT * FROM keeper.user WHERE username=$1 LIMIT 1", req.Username); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, &st.ErrorUsernameNotFound{
				Username: req.Username,
			}
		}
		return nil, err
	}
	return resp, nil
}
