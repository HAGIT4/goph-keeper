package service

import (
	"context"
	"errors"

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

type ListLoginPassKeywordsReq struct {
	UserID uuid.UUID
}

type ListLoginPassKeywordsResp struct {
	Keywords []string
}

func (ks *keeperService) ListLoginPassKeywords(ctx context.Context, req *ListLoginPassKeywordsReq) (resp *ListLoginPassKeywordsResp, err error) {
	stReq := &st.ListLoginPassKeywordsReq{
		UserID: req.UserID,
	}
	stResp, err := ks.storage.ListLoginPassKeywords(ctx, stReq)
	if err != nil {
		if errors.Is(err, &st.ErrorNoLoginPassFoundForUser{}) {
			return resp, nil
		}
		return nil, err
	}
	resp.Keywords = stResp.Keywords
	return resp, nil
}
