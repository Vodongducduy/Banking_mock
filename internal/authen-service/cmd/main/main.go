package main

import (
	grpc_server "banking/internal/authen-service/cmd/grpc-server"
	"banking/internal/authen-service/repositorys"
	"banking/internal/authen-service/usecases"
	"banking/internal/user-service/database"
	"banking/packages/config"
	"banking/packages/customResponse"
	"fmt"
	"log"
	"net"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = ""
	DB_NAME     = "banking"
)

func main() {
	database.Connect(fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_NAME))
	database.Migrate()
	autheRepo := repositorys.NewAuthRepository(database.Instance)
	usecases.NewAuthUsecase(autheRepo)

	lis, err := net.Listen("tcp", config.GrpcAuthPort)
	if err != nil {
		customResponse.FailErr("Fail to listen", err)
		return
	}

	if err = grpc_server.RunGRPCServer(false, lis); err != nil {
		customResponse.FailErr("Fail to serve", err)
		return
	}
	log.Println("gRPC server auth is running....")

}
