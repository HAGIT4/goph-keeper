package service

import (
	"context"
	"fmt"
)

type SaveLoginPassReq struct {
	Username    string
	EncKeyword  string
	EncLogin    string
	EncPassword string
	EncMeta     string
}
type SaveLoginPassResp struct {
	EncKeyword string
}

func (ks *keeperService) SaveLoginPass(ctx context.Context, req *SaveLoginPassReq) (resp *SaveLoginPassResp, err error) {
	fmt.Println("loginpass saved")
	return &SaveLoginPassResp{}, nil
}
