package service

import (
	"context"
	"fmt"

	"encoding/hex"

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

	encLoginBytes, err := as.encrypt([]byte(login))
	if err != nil {
		return err
	}
	encPassBytes, err := as.encrypt([]byte(pass))
	if err != nil {
		return err
	}
	encMetaBytes, err := as.encrypt([]byte(meta))
	if err != nil {
		return err
	}

	encLogin := hex.EncodeToString(encLoginBytes)
	encPass := hex.EncodeToString(encPassBytes)
	encMeta := hex.EncodeToString(encMetaBytes)
	grpcReq := &pb.SaveLoginPassRequest{
		EncKeyword:  keyword,
		EncLogin:    encLogin,
		EncPassword: encPass,
		EncMeta:     encMeta,
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

	decodedLoginBytes, err := hex.DecodeString(grpcResp.GetEncLogin())
	if err != nil {
		return err
	}
	decriptedLoginBytes, err := as.decrypt(decodedLoginBytes)
	if err != nil {
		return err
	}
	decriptedLogin := string(decriptedLoginBytes)

	decodedPass, err := hex.DecodeString(grpcResp.GetEncPassword())
	if err != nil {
		return err
	}
	decriptedPassBytes, err := as.decrypt(decodedPass)
	if err != nil {
		return err
	}
	decriptedPass := string(decriptedPassBytes)

	decodedMeta, err := hex.DecodeString(grpcResp.GetEncMeta())
	if err != nil {
		return err
	}
	decriptedMetaBytes, err := as.decrypt(decodedMeta)
	if err != nil {
		return err
	}
	decriptedMeta := string(decriptedMetaBytes)

	fmt.Println(grpcResp.EncKeyword)
	fmt.Println(decriptedLogin)
	fmt.Println(decriptedPass)
	fmt.Println(decriptedMeta)
	return nil
}
