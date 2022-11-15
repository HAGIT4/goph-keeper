package service

import (
	"context"
	"fmt"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (as *agentService) SaveLoginPass(ctx context.Context) (err error) {
	var keyword, login, pass, meta string
	fmt.Println("Enter keyword: ")
	fmt.Scanln(&keyword)
	fmt.Println("Enter login: ")
	fmt.Scanln(&login)
	fmt.Println("Enter password: ")
	fmt.Scanln(&pass)
	fmt.Println("Enter meta: ")
	fmt.Scanln(&meta)

	grpcReq := &pb.SaveLoginPassRequest{
		EncKeyword:  keyword,
		EncLogin:    login,
		EncPassword: pass,
		EncMeta:     meta,
	}
	_, err = as.agentGRPC.SaveLoginPass(ctx, grpcReq)
	if err != nil {
		return err
	}
	return nil
}

func (as *agentService) ListLoginPassKeywords(ctx context.Context) (err error) {
	grpcReq := &pb.ListLoginPassKeywordsRequest{}
	grpcResp, err := as.agentGRPC.ListLoginPassKeywords(ctx, grpcReq)
	if err != nil {
		return nil
	}
	for _, keyword := range grpcResp.Keywords {
		fmt.Println(keyword)
	}
	return nil
}

func (as *agentService) GetLoginPass(ctx context.Context) (err error) {
	var keyword string
	fmt.Println("Enter keyword:")
	fmt.Scanln(&keyword)

	grpcReq := &pb.GetLoginPassRequest{
		EncKeyword: keyword,
	}
	grpcResp, err := as.agentGRPC.GetLoginPass(ctx, grpcReq)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(grpcResp.EncKeyword)
	fmt.Println(grpcResp.EncLogin)
	fmt.Println(grpcResp.EncPassword)
	fmt.Println(grpcResp.EncMeta)
	return nil
}
