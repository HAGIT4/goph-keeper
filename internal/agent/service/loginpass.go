package service

import (
	"context"
	"fmt"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (as agentService) SaveLoginPass(ctx context.Context) (err error) {
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
