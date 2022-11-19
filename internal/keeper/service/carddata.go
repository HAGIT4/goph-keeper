package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type SaveCardDataReq struct {
	UserID        uuid.UUID
	EncKeyword    string
	EncCardData   string
	EncCardHolder string
	EncCardCode   string
	EncMeta       string
}
type SaveCardDataResp struct {
	EncKeyword string
}

func (ks *keeperService) SaveCardData(ctx context.Context, req *SaveCardDataReq) (resp *SaveCardDataResp, err error) {
	stReq := &st.CreateCardDataReq{
		ID:         uuid.New(),
		UserID:     req.UserID,
		Keyword:    req.EncKeyword,
		CardNumber: req.EncCardData,
		CardHolder: req.EncCardHolder,
		CardCode:   req.EncCardCode,
		Meta:       req.EncMeta,
	}
	_, err = ks.storage.CreateCardData(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &SaveCardDataResp{
		EncKeyword: req.EncKeyword,
	}
	return resp, nil
}

type GetCardDataReq struct {
	UserID  uuid.UUID
	Keyword string
}

type GetCardDataResp struct {
	UserID        uuid.UUID
	EncKeyword    string
	EncCardNumber string
	EncCardHolder string
	EncCardCode   string
	EncMeta       string
}

func (ks *keeperService) GetCardData(ctx context.Context, req *GetCardDataReq) (resp *GetCardDataResp, err error) {
	stReq := &st.ReadCardDataByKeywordReq{
		UserID:  req.UserID,
		Keyword: req.Keyword,
	}
	stResp, err := ks.storage.ReadCardDataByKeyword(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &GetCardDataResp{
		UserID:        stResp.UserID,
		EncKeyword:    stResp.Keyword,
		EncCardNumber: stResp.CardNumber,
		EncCardHolder: stResp.CardHolder,
		EncCardCode:   stResp.CardCode,
		EncMeta:       stResp.Meta,
	}
	return resp, nil
}

type ListCardDataKeywordsReq struct {
	UserID uuid.UUID
}
type ListCardDataKeywordsResp struct {
	Keywords []string
}

func (ks *keeperService) ListCardDataKeywords(ctx context.Context, req *ListCardDataKeywordsReq) (resp *ListCardDataKeywordsResp, err error) {
	stReq := &st.ListCardDataKeywordsReq{
		UserID: req.UserID,
	}
	stResp, err := ks.storage.ListCardDataKeywords(ctx, stReq)
	if err != nil {
		if errors.Is(err, &st.ErrorNoCardDataFoundForUser{}) {
			return resp, nil
		}
		return nil, err
	}
	resp = &ListCardDataKeywordsResp{
		Keywords: stResp.Keywords,
	}
	return resp, nil
}
