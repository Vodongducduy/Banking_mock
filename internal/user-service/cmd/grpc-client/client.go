package grpc_client

import (
	"banking/packages/config"
	"banking/packages/customResponse"
	"banking/packages/pb/auth"
	"fmt"
	"google.golang.org/grpc"
)

type IGRPCAuthClient interface {
	SetUpCAuthClient() auth.IsAuthClient
}
type GRPCAuthClient struct{}

func NewGRPCAuthClient() *GRPCAuthClient {
	return &GRPCAuthClient{}
}

func (G *GRPCAuthClient) SetUpCAuthClient() auth.IsAuthClient {
	addr := "localhost" + config.GrpcAuthPort
	conn, diaErr := grpc.Dial(addr, grpc.WithInsecure())
	if diaErr != nil {
		customResponse.FailErr("SetUpClientUser: fail connecting to sever", diaErr)
		return nil
	}
	authClient := auth.NewIsAuthClient(conn)
	fmt.Println("Listen to AuthSever on port", config.GrpcAuthPort)
	return authClient
}
