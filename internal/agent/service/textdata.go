package service

import (
	"context"
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

	encKeywordBytes, err := as.encrypt([]byte(keyword))
	if err != nil {
		return err
	}
	encTextdataBytes, err := as.encrypt([]byte(textdata))
	if err != nil {
		return err
	}
	encMetaBytes, err := as.encrypt([]byte(meta))
	if err != nil {
		return err
	}

	grpcReq := &pb.SaveTextDataRequest{
		EncKeyword:  string(encKeywordBytes),
		EncTextData: string(encTextdataBytes),
		EncMeta:     string(encMetaBytes),
	}
	grpcResp, err := as.agentGRPC.SaveTextData(ctx, grpcReq)
	if err != nil {
		return err
	}
	fmt.Println(grpcResp.EncKeyword)
	return nil
}
