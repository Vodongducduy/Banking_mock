package grpc_server

import (
	"banking/internal/authen-service/repositorys"
	"banking/internal/authen-service/usecases"
	"banking/internal/user-service/database"
	"banking/packages/config"
	"banking/packages/customResponse"
	"banking/packages/pb/auth"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	cerFile = "packages/ssl/server.crt"
	keyFile = "packages/ssl/server.pem"
)

type AuthServer struct {
	auth.IsAuthServer
	AuthUsecase usecases.IAuthUsecase
}

func RunGRPCServer(enable bool, lis net.Listener) error {
	var otps []grpc.ServerOption
	if enable {
		creds, err := credentials.NewClientTLSFromFile(cerFile, keyFile)
		if err != nil {
			return err
		}
		otps = append(otps, grpc.Creds(creds))
	}

	s := grpc.NewServer(otps...)

	authRepo := repositorys.NewAuthRepository(database.Instance)
	authUsecase := usecases.NewAuthUsecase(authRepo)

	auth.RegisterIsAuthServer(s, &AuthServer{
		AuthUsecase: authUsecase,
	})

	log.Printf("Listen on %s\n", config.GrpcAuthPort)
	return s.Serve(lis)
}

func (a *AuthServer) IsAuth(ctx context.Context, in *auth.IsAuthRequest) (*auth.IsAuthResponse, error) {
	log.Printf("Login request: %v\n", in)
	res, err := a.AuthUsecase.IsUser(in)
	if err != nil {
		customResponse.FailErr("Error to send Request", err)
	}
	if res == nil {
		customResponse.FailErr("Error to get Response", errors.New("response empty"))
		return nil, err
	}
	return res, nil
}
