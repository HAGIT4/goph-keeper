package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	st "github.com/hagit4/goph-keeper/internal/keeper/storage"
)

type SaveBinaryReq struct {
	UserID     uuid.UUID
	EncKeyword string
	EncBin     []byte
	EncMeta    string
}

type SaveBinaryResp struct {
	EncKeyword string
}

func (ks *keeperService) SaveBinary(ctx context.Context, req *SaveBinaryReq) (resp *SaveBinaryResp, err error) {
	stReq := &st.CreateBinaryDataReq{
		ID:      uuid.New(),
		UserID:  req.UserID,
		Keyword: req.EncKeyword,
		Bin:     req.EncBin,
		Meta:    req.EncMeta,
	}
	_, err = ks.storage.CreateBinaryData(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &SaveBinaryResp{
		EncKeyword: req.EncKeyword,
	}
	return resp, nil
}

type GetBinaryReq struct {
	UserID  uuid.UUID
	Keyword string
}

type GetBinaryResp struct {
	UserID  uuid.UUID
	Keyword string
	Bin     []byte
	Meta    string
}

func (ks *keeperService) GetBinary(ctx context.Context, req *GetBinaryReq) (resp *GetBinaryResp, err error) {
	stReq := &st.ReadBinaryDataByKeywordReq{
		UserID:  req.UserID,
		Keyword: req.Keyword,
	}
	stResp, err := ks.storage.ReadBinaryDataByKeyword(ctx, stReq)
	if err != nil {
		return nil, err
	}
	resp = &GetBinaryResp{
		UserID:  stResp.UserID,
		Keyword: stResp.Keyword,
		Bin:     stResp.Bin,
		Meta:    stResp.Meta,
	}
	return resp, nil
}

type ListBinaryKeywordsReq struct {
	UserID uuid.UUID
}

type ListBinaryKeywordsResp struct {
	Keywords []string
}

func (ks *keeperService) ListBinaryKeywords(ctx context.Context, req *ListBinaryKeywordsReq) (resp *ListBinaryKeywordsResp, err error) {
	stReq := &st.ListBinaryDataKeywordsReq{
		UserID: req.UserID,
	}
	stResp, err := ks.storage.ListBinaryKeywords(ctx, stReq)
	if err != nil {
		if errors.Is(err, &st.ErrorNoBinaryFoundForUser{}) {
			return resp, nil
		}
		return nil, err
	}
	resp = &ListBinaryKeywordsResp{
		Keywords: stResp.Keywords,
	}
	return resp, nil
}
