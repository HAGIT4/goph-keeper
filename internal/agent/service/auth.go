package service

import (
	"context"
	"fmt"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"golang.org/x/term"
)

func (as *agentService) RegisterUser(ctx context.Context) (err error) {
	username, password, err := readUsernameAndPassword()
	if err != nil {
		return err
	}
	grpcReq := &pb.RegisterRequest{
		Username: username,
		Password: password,
	}
	grpcResp, err := as.agentGRPC.RegisterUser(ctx, grpcReq)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(grpcResp.Username)
	return nil
}

func (as *agentService) Login(ctx context.Context) (err error) {
	username, password, err := readUsernameAndPassword()
	if err != nil {
		return err
	}
	grpcReq := &pb.LoginRequest{
		Username: username,
		Password: password,
	}
	_, err = as.agentGRPC.Login(ctx, grpcReq)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func readUsernameAndPassword() (username string, password string, err error) {
	fmt.Println("Enter login and password")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("(Input masked)Password: ")
	var passwordBytes []byte
	passwordBytes, err = term.ReadPassword(0)
	if err != nil {
		return "", "", err
	}
	fmt.Print("\n")
	password = string(passwordBytes)
	return username, password, nil
}
