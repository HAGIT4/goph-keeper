package service

import (
	"context"

	"github.com/google/uuid"
	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type SaveLoginPassReq struct {
	UserID      uuid.UUID
	EncKeyword  string
	EncLogin    string
	EncPassword string
	EncMeta     string
}
type SaveLoginPassResp struct {
	EncKeyword string
}

func (ks *keeperService) SaveLoginPass(ctx context.Context, req *SaveLoginPassReq) (resp *SaveLoginPassResp, err error) {
	stReq := &st.CreateLoginPassReq{
		LoginPass: st.LoginPass{
			ID:       uuid.New(),
			UserID:   req.UserID,
			Keyword:  req.EncKeyword,
			Login:    req.EncLogin,
			Password: []byte(req.EncPassword),
			Meta:     req.EncMeta,
		},
	}
	_, err = ks.storage.CreateLoginPass(ctx, stReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
