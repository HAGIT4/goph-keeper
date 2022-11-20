package service

import (
	"context"
	"encoding/hex"
	"fmt"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (as *agentService) SaveCardData(ctx context.Context) (err error) {
	var keyword, cardNumber, cardHolder, cardCode, meta string
	fmt.Println("Enter keyword:")
	fmt.Scanln(&keyword)
	fmt.Println("Enter card number:")
	fmt.Scanln(&cardNumber)
	fmt.Println("Enter card holder:")
	fmt.Scanln(&cardHolder)
	fmt.Println("Enter card code:")
	fmt.Scanln(&cardCode)
	fmt.Println("Enter meta:")
	fmt.Scanln(&meta)

	encCardNumberBytes, err := as.encrypt([]byte(cardNumber))
	if err != nil {
		return err
	}
	encCardHolderBytes, err := as.encrypt([]byte(cardHolder))
	if err != nil {
		return err
	}
	encCardCodeBytes, err := as.encrypt([]byte(cardCode))
	if err != nil {
		return err
	}
	encMetaBytes, err := as.encrypt([]byte(cardCode))
	if err != nil {
		return err
	}

	encCardNumber := hex.EncodeToString(encCardNumberBytes)
	encCardHolder := hex.EncodeToString(encCardHolderBytes)
	encCardCode := hex.EncodeToString(encCardCodeBytes)
	encMeta := hex.EncodeToString(encMetaBytes)
	grpcReq := &pb.SaveCardDataRequest{
		EncKeyword:    keyword,
		EncCardNumber: encCardNumber,
		EncCardHolder: encCardHolder,
		EncCardCode:   encCardCode,
		EncMeta:       encMeta,
	}
	_, err = as.agentGRPC.SaveCardData(ctx, grpcReq)
	if err != nil {
		return err
	}
	return nil
}

func (as *agentService) GetCardData(ctx context.Context) (err error) {
	var keyword string
	fmt.Println("Enter keyword:")
	fmt.Scanln(&keyword)

	grpcReq := &pb.GetCardDataRequest{
		EncKeyword: keyword,
	}
	grpcResp, err := as.agentGRPC.GetCardData(ctx, grpcReq)
	if err != nil {
		return err
	}

	decodedCardNumberBytes, err := hex.DecodeString(grpcResp.GetEncCardCode())
	if err != nil {
		return err
	}
	decriptedCardNumberBytes, err := as.decrypt(decodedCardNumberBytes)
	if err != nil {
		return err
	}
	decriptedCardNumber := string(decriptedCardNumberBytes)

	decodedCardHolderBytes, err := hex.DecodeString(grpcResp.GetEncCardHolder())
	if err != nil {
		return err
	}
	decriptedCardHolderBytes, err := as.decrypt(decodedCardHolderBytes)
	if err != nil {
		return err
	}
	decriptedCardHolder := string(decriptedCardHolderBytes)

	decodedCardCodeBytes, err := hex.DecodeString(grpcResp.GetEncCardCode())
	if err != nil {
		return err
	}
	decriptedCardCodeBytes, err := as.decrypt(decodedCardCodeBytes)
	if err != nil {
		return err
	}
	decriptedCardCode := string(decriptedCardCodeBytes)

	decodedMetaBytes, err := hex.DecodeString(grpcResp.GetEncMeta())
	if err != nil {
		return nil
	}
	decriptedMetaBytes, err := as.decrypt(decodedMetaBytes)
	if err != nil {
		return nil
	}
	decriptedMeta := string(decriptedMetaBytes)

	fmt.Println(grpcResp.GetEncKeyword())
	fmt.Println(decriptedCardNumber)
	fmt.Println(decriptedCardHolder)
	fmt.Println(decriptedCardCode)
	fmt.Println(decriptedMeta)
	return nil
}

func (as *agentService) ListCardDataKeywords(ctx context.Context) (err error) {
	grpcReq := &pb.ListCardDataKeywordsRequest{}
	grpcResp, err := as.agentGRPC.ListCardDataKeywords(ctx, grpcReq)
	if err != nil {
		return err
	}
	for _, keyword := range grpcResp.Keywords {
		fmt.Println(keyword)
	}
	return nil
}
