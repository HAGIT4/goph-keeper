package postgres

import (
	"database/sql"

	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
	"golang.org/x/net/context"
)

func (ps *keeperPostgresStorage) CreateTextData(ctx context.Context, req *st.CreateTextDataReq) (resp *st.CreateTextDataResp, err error) {
	resp = &st.CreateTextDataResp{}
	query := `INSERT INTO keeper.textdata (id, user_id, keyword, textdata, meta) VALUES ($1, $2, $3, $4, $5)`
	_, err = ps.db.ExecContext(ctx, query, req.TextData.ID, req.TextData.UserID, req.TextData.Keyword,
		req.TextData.TextData, req.TextData.Meta)
	if err != nil {
		return nil, st.NewErrorTextDataNotCreated(err)
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) ReadTextData(ctx context.Context, req *st.ReadTextDataReq) (resp *st.ReadTextDataResp, err error) {
	resp = &st.ReadTextDataResp{}
	query := `SELECT * FROM keeper.textdata WHERE id=$1 AND user_id=$2 LIMIT 1`
	if err = ps.db.Get(resp, query, req.ID, req.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, st.NewErrorTextDataNotFoundByID(req.ID, req.UserID, err)
		}
		return nil, err
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) UpdateTextData(ctx context.Context, req *st.UpdateTextDataReq) (resp *st.UpdateTextDataResp, err error) {
	resp = &st.UpdateTextDataResp{}
	query := `UPDATE keeper.textdata SET textdata=$1 WHERE id=$2 AND user_id=$3 AND keyword=$4`
	if _, err = ps.db.ExecContext(ctx, query, req.TextData.TextData, req.TextData.ID,
		req.TextData.UserID, req.TextData.Keyword); err != nil {
		return nil, st.NewErrorTextDataNotUpdated(req.TextData.ID, req.TextData.UserID, err)
	}
	return resp, nil
}

func (ps *keeperPostgresStorage) DeleteTextData(ctx context.Context, req *st.DeleteTextDataReq) (resp *st.DeleteTextDataResp, err error) {
	resp = &st.DeleteTextDataResp{}
	query := `DELETE FROM keeper.textdata WHERE id=$1 AND user_id=$2`
	if _, err = ps.db.ExecContext(ctx, query, req.ID, req.UserID); err != nil {
		return nil, st.NewErrorTextDataNotDeleted(req.ID, req.UserID, err)
	}
	return resp, nil
}
