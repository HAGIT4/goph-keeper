package postgres

import (
	"database/sql"

	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
	"golang.org/x/net/context"
)

func (ps *keeperPostgresStorage) CreateLoginPass(ctx context.Context, req *st.CreateLoginPassReq) (resp *st.CreateLoginPassResp, err error) {
	resp = &st.CreateLoginPassResp{}
	query := `INSERT INTO keeper.loginpass (id, user_id, keyword, login, password, meta) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = ps.db.ExecContext(ctx, query, req.ID, req.UserID, req.Keyword, req.Login, req.Password, req.Meta)
	if err != nil {
		return nil, st.NewErrorLoginPassNotCreated(err)
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadLoginPassByID(ctx context.Context, req *st.ReadLoginPassByIDreq) (resp *st.ReadLoginPassByIDresp, err error) {
	resp = &st.ReadLoginPassByIDresp{}
	query := `SELECT * FROM keeper.loginpass WHERE id=$1 AND user_id=$2 LIMIT 1`
	if err = ps.db.Get(resp, query, req.ID, req.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, st.NewErrorLoginPassNotFoundByID(req.ID, req.UserID, err)
		}
		return nil, err
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadLoginPassByKeyword(ctx context.Context, req *st.ReadLoginPassByKeywordReq) (resp *st.ReadLoginPassByKeywordResp, err error) {
	resp = &st.ReadLoginPassByKeywordResp{}
	query := `SELECT * FROM keeper.loginpass WHERE keyword=$1 AND user_id=$2 LIMIT 1`
	if err = ps.db.Get(resp, query, req.Keyword, req.UserID); err != nil {
		if err != sql.ErrNoRows {
			return nil, st.NewErrorLoginPassNotFoundByKeyword(req.Keyword, req.UserID, err)
		}
		return nil, err
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) UpdateLogin(ctx context.Context, req *st.UpdateLoginByIDreq) (resp *st.UpdateLoginByIDresp, err error) {
	resp = &st.UpdateLoginByIDresp{}
	query := `UPDATE keeper.loginpass SET login=$1 WHERE id=$2 AND user_id=$3`
	if _, err = ps.db.ExecContext(ctx, query, req.Login, req.ID, req.UserID); err != nil {
		return nil, st.NewErrorLoginNotUpdatedByID(req.ID, req.UserID, req.Login, err)
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ListLoginPassKeywords(ctx context.Context, req *st.ListLoginPassKeywordsReq) (resp *st.ListLoginPassKeywordsResp, err error) {
	resp = &st.ListLoginPassKeywordsResp{}
	var keywords []string
	query := `SELECT keyword FROM keeper.loginpass WHERE user_id=$1`
	if err = ps.db.SelectContext(ctx, &keywords, query, req.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, st.NewErrorNoLoginPassFoundForUser(req.UserID)
		}
		return nil, err
	}
	resp.Keywords = keywords
	return resp, nil
}
