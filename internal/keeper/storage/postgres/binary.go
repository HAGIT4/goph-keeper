package postgres

import (
	"database/sql"

	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
	"golang.org/x/net/context"
)

func (ps *keeperPostgresStorage) CreateBinaryData(ctx context.Context, req *st.CreateBinaryDataReq) (resp *st.CreateBinaryDataResp, err error) {
	resp = &st.CreateBinaryDataResp{}
	query := `INSERT INTO keeper.binary (id, user_id, keyword, bin, meta) VALUES ($1, $2, $3, $4, $5)`
	_, err = ps.db.ExecContext(ctx, query, req.ID, req.UserID, req.Keyword, req.Bin, req.Meta)
	if err != nil {
		return nil, st.NewErrorLoginPassNotCreated(err)
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadBinaryData(ctx context.Context, req *st.ReadBinaryDataByKeywordReq) (resp *st.ReadBinaryDataByKeywordResp, err error) {
	resp = &st.ReadBinaryDataByKeywordResp{}
	query := `SELECT * FROM keeper.binary WHERE keyword=$1 AND user_id=$2 LIMIT 1`
	if err = ps.db.Get(resp, query, req.Keyword, req.UserID); err != nil {
		if err != sql.ErrNoRows {
			return nil, st.NewErrorBinaryNotFoundByKeyword(req.Keyword, req.UserID, err)
		}
		return nil, err
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ListBinaryKeywords(ctx context.Context, req *st.ListBinaryDataKeywordsReq) (resp *st.ListBinaryDataKeywordsResp, err error) {
	resp = &st.ListBinaryDataKeywordsResp{}
	var keywords []string
	query := `SELECT keyword FROM keeper.binary WHERE user_id=$1`
	if err = ps.db.SelectContext(ctx, &keywords, query, req.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, st.NewErrorNoBinaryFoundForUser(req.UserID)
		}
		return nil, err
	}
	resp.Keywords = keywords
	return resp, nil
}
