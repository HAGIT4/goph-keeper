package service

import (
	"context"
	"fmt"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

const (
	hashingCost int = 14
)

func (as *agentService) RegisterUser(ctx context.Context) (err error) {
	username, password, err := readUsernameAndPassword()
	if err != nil {
		return err
	}
	passwordHash, err := hashPassword(password)
	if err != nil {
		return err
	}
	grpcReq := &pb.RegisterRequest{
		Username: username,
		PassHash: passwordHash,
	}
	_, err = as.agentGRPC.RegisterUser(ctx, grpcReq)
	if err != nil {
		return err
	}
	return nil
}

func (as *agentService) Login(ctx context.Context) (err error) {
	username, password, err := readUsernameAndPassword()
	if err != nil {
		return err
	}
	passwordHash, err := hashPassword(password)
	if err != nil {
		return err
	}
	grpcReq := &pb.LoginRequest{
		Username: username,
		PassHash: passwordHash,
	}
	_, err = as.agentGRPC.Login(ctx, grpcReq)
	if err != nil {
		return err
	}
	return nil
}

func readUsernameAndPassword() (username string, password string, err error) {
	fmt.Println("Enter new users login and password")
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

func hashPassword(password string) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(hash), hashingCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}
