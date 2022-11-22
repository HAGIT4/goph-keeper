package service

import (
	"context"
	"encoding/hex"
	"fmt"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (as *agentService) SaveTextData(ctx context.Context) (err error) {
	var keyword, textdata, meta string
	fmt.Println("Enter keyword: ")
	fmt.Scanln(&keyword)
	fmt.Println("Enter text: ")
	fmt.Scanln(&textdata)
	fmt.Println("Enter meta: ")
	fmt.Scanln(&meta)

	encTextdataBytes, err := as.encrypt([]byte(textdata))
	if err != nil {
		return err
	}
	encMetaBytes, err := as.encrypt([]byte(meta))
	if err != nil {
		return err
	}
	encTextData := hex.EncodeToString(encTextdataBytes)
	encMeta := hex.EncodeToString(encMetaBytes)

	grpcReq := &pb.SaveTextDataRequest{
		EncKeyword:  string(keyword),
		EncTextData: string(encTextData),
		EncMeta:     string(encMeta),
	}
	grpcResp, err := as.agentGRPC.SaveTextData(ctx, grpcReq)
	if err != nil {
		return err
	}
	fmt.Println(grpcResp.EncKeyword)
	return nil
}

func (as *agentService) GetTextData(ctx context.Context) (err error) {
	var keyword string
	fmt.Println("Enter keyword:")
	fmt.Scanln(&keyword)

	grpcReq := &pb.GetTextDataRequest{
		EncKeyword: keyword,
	}
	grpcResp, err := as.agentGRPC.GetTextData(ctx, grpcReq)
	if err != nil {
		return err
	}

	decodedTextBytes, err := hex.DecodeString(grpcResp.GetEncTextData())
	if err != nil {
		return err
	}
	decriptedTextBytes, err := as.decrypt(decodedTextBytes)
	if err != nil {
		return err
	}
	decriptedText := string(decriptedTextBytes)

	decodedMeta, err := hex.DecodeString(grpcResp.GetEncMeta())
	if err != nil {
		return err
	}
	decriptedMetaBytes, err := as.decrypt(decodedMeta)
	if err != nil {
		return err
	}
	decriptedMeta := string(decriptedMetaBytes)

	fmt.Println(grpcResp.GetEncKeyword())
	fmt.Println(decriptedText)
	fmt.Println(decriptedMeta)
	return nil
}

func (as *agentService) ListTextDataKeywords(ctx context.Context) (err error) {
	grpcReq := &pb.ListTextDataKeywordsRequest{}
	grpcResp, err := as.agentGRPC.ListTextDataKeywords(ctx, grpcReq)
	if err != nil {
		return err
	}
	for _, keyword := range grpcResp.Keywords {
		fmt.Println(keyword)
	}
	return nil
}
