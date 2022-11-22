package service

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"os"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

func (as *agentService) SaveBinary(ctx context.Context) (err error) {
	var keyword, path, meta string
	fmt.Println("Enter keyword:")
	fmt.Scanln(&keyword)
	fmt.Println("Enter path to file:")
	fmt.Scanln(&path)
	fmt.Println("Enter meta:")
	fmt.Scanln(&meta)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fStat, err := f.Stat()
	if err != nil {
		return err
	}

	bin := make([]byte, fStat.Size())
	reader := bufio.NewReader(f)
	_, err = reader.Read(bin)

	encMetaBytes, err := as.encrypt([]byte(meta))
	if err != nil {
		return err
	}
	encMeta := hex.EncodeToString(encMetaBytes)

	grpcReq := &pb.SaveBinaryRequest{
		Keyword: keyword,
		Bin:     bin,
		Meta:    encMeta,
	}
	_, err = as.agentGRPC.SaveBinary(ctx, grpcReq)
	if err != nil {
		return err
	}
	return nil
}

func (as *agentService) GetBinary(ctx context.Context) (err error) {
	var keyword string
	fmt.Println("Enter keyword:")
	fmt.Scanln(&keyword)

	grpcReq := &pb.GetBinaryRequest{
		Keyword: keyword,
	}
	grpcResp, err := as.agentGRPC.GetBinary(ctx, grpcReq)
	if err != nil {
		return err
	}

	decodedMeta, err := hex.DecodeString(grpcResp.GetMeta())
	if err != nil {
		return err
	}
	decriptedMetaBytes, err := as.decrypt(decodedMeta)
	if err != nil {
		return err
	}
	decriptedMeta := string(decriptedMetaBytes)

	fmt.Println(grpcResp.Keyword)
	fmt.Println(decriptedMeta)
	err = os.WriteFile("keep.bin", grpcResp.GetBin(), 0644)
	if err != nil {
		return err
	}
	fmt.Println("file saved")
	return nil
}

func (as *agentService) ListBinaryKeywords(ctx context.Context) (err error) {
	grpcReq := &pb.ListBinaryKeywordsRequest{}
	grpcResp, err := as.agentGRPC.ListBinaryKeywords(ctx, grpcReq)
	if err != nil {
		return err
	}
	for _, keyword := range grpcResp.Keywords {
		fmt.Println(keyword)
	}
	return nil
}
