package interceptor

import (
	"context"
	"fmt"

	ag "github.com/hagit4/goph-keeper/internal/agent/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewAuthInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		if method == "/keeper.Auth/Login" || method == "/keeper.Auth/Register" {
			if err = invoker(ctx, method, req, reply, cc, opts...); err != nil {
				fmt.Println("interceptor err", err)
				return err
			}
			return nil
		}
		token, err := ag.GetAuthToken()
		if err != nil {
			return err
		}
		ctx = metadata.AppendToOutgoingContext(ctx, "token", token)
		if err = invoker(ctx, method, req, reply, cc, opts...); err != nil {
			return err
		}
		return nil
	}
}
