package service

import (
	"context"

	"github.com/google/uuid"
	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type SaveTextDataReq struct {
	UserID      uuid.UUID
	EncKeyword  string
	EncTextData string
	EncMeta     string
}

type SaveTextDataResp struct {
	EncKeyword string
}

func (ks *keeperService) SaveTextData(ctx context.Context, req *SaveTextDataReq) (resp *SaveLoginPassResp, err error) {
	stReq := &st.CreateTextDataReq{
		TextData: st.TextData{
			ID:       uuid.New(),
			UserID:   req.UserID,
			Keyword:  req.EncKeyword,
			TextData: req.EncTextData,
			Meta:     req.EncMeta,
		},
	}
	stResp, err := ks.storage.CreateTextData(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &SaveLoginPassResp{
		EncKeyword: stResp.Keyword,
	}
	return resp, nil
}

type ReadTextDataByIDreq struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

type ReadTextDataByIDresp struct {
	Keyword  string
	TextData string
	Meta     string
}

func (ks *keeperService) ReadTextDataByID(ctx context.Context, req *ReadTextDataByIDreq) (resp *ReadTextDataByIDresp, err error) {
	stReq := &st.ReadTextDataByIDreq{
		UserID: req.UserID,
		ID:     req.ID,
	}
	stResp, err := ks.storage.ReadTextDataByID(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &ReadTextDataByIDresp{
		Keyword:  stResp.TextData.Keyword,
		TextData: stResp.TextData.TextData,
		Meta:     stResp.TextData.Meta,
	}
	return resp, nil
}

type ReadTextDataByKeywordReq struct {
	UserID  uuid.UUID
	Keyword string
}

type ReadTextDataByKeywordResp struct {
	Keyword  string
	TextData string
	Meta     string
}

func (ks *keeperService) ReadTextDataByKeyword(ctx context.Context, req *ReadTextDataByKeywordReq) (resp *ReadTextDataByKeywordResp, err error) {
	stReq := &st.ReadTextDataByKeywordReq{
		UserId:  req.UserID,
		Keyword: req.Keyword,
	}
	stResp, err := ks.storage.ReadTextDataByKeyword(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &ReadTextDataByKeywordResp{
		Keyword:  stResp.Keyword,
		TextData: stResp.TextData,
		Meta:     stResp.Meta,
	}
	return resp, nil
}

type ListTextDataKeywordsReq struct {
	UserID uuid.UUID
}
type ListTextDataKeywordsResp struct {
	Keywords []string
}

func (ks *keeperService) ListTextDataKeywords(ctx context.Context, req *ListTextDataKeywordsReq) (resp *ListTextDataKeywordsResp, err error) {
	stReq := &st.ListTextDataKeywordsReq{
		UserID: req.UserID,
	}
	stResp, err := ks.storage.ListTextDataKeywords(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &ListTextDataKeywordsResp{
		Keywords: stResp.Keywords,
	}
	return resp, nil
}
