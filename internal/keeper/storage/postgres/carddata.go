package postgres

import (
	"database/sql"

	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
	"golang.org/x/net/context"
)

func (ps *keeperPostgresStorage) CreateCardData(ctx context.Context, req *st.CreateCardDataReq) (resp *st.CreateCardDataResp, err error) {
	resp = &st.CreateCardDataResp{}
	query := `INSERT INTO keeper.carddata (id, user_id, keyword, card_number, card_holder, card_code, meta) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = ps.db.ExecContext(ctx, query, req.ID, req.UserID, req.Keyword, req.CardNumber, req.CardHolder, req.CardCode, req.Meta)
	if err != nil {
		return nil, st.NewErrorLoginPassNotCreated(err)
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadCardDataByKeyword(ctx context.Context, req *st.ReadCardDataByKeywordReq) (resp *st.ReadCardDataByKeywordResp, err error) {
	resp = &st.ReadCardDataByKeywordResp{}
	query := `SELECT * FROM keeper.carddata WHERE keyword=$1 AND user_id=$2 LIMIT 1`
	if err = ps.db.Get(resp, query, req.Keyword, req.UserID); err != nil {
		if err != sql.ErrNoRows {
			return nil, st.NewErrorCardDataNotFoundByKeyword(req.Keyword, req.UserID, err)
		}
		return nil, err
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ListCardDataKeywords(ctx context.Context, req *st.ListCardDataKeywordsReq) (resp *st.ListCardDataKeywordsResp, err error) {
	resp = &st.ListCardDataKeywordsResp{}
	var keywords []string
	query := `SELECT keyword FROM keeper.carddata WHERE user_id=$1`
	if err = ps.db.SelectContext(ctx, &keywords, query, req.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, st.NewErrorNoLoginPassFoundForUser(req.UserID)
		}
		return nil, err
	}
	resp.Keywords = keywords
	return resp, nil
}
