package grcp

import (
	"context"
	"os"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

const (
	tokenFile string = ".token"
)

func (ag *agentGRPC) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (resp *pb.RegisterResponse, err error) {
	resp, err = ag.authClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ag *agentGRPC) Login(ctx context.Context, req *pb.LoginRequest) (resp *pb.LoginResponse, err error) {
	resp, err = ag.authClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if err = setToken(resp.AccessToken); err != nil {
		return nil, err
	}
	return resp, nil
}

func setToken(token string) (err error) {
	f, err := os.Create(tokenFile)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write([]byte(token)); err != nil {
		return err
	}
	return nil
}

func GetAuthToken() (authToken string, err error) {
	token, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	authToken = string(token)

	if len(authToken) == 0 {
		return "", &ErrorNoAuthToken{}
	}
	return authToken, nil
}
