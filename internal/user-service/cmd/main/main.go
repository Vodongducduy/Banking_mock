package main

import (
	grpc_client "banking/internal/user-service/cmd/grpc-client"
	"banking/internal/user-service/controllers"
	"banking/internal/user-service/database"
	"banking/internal/user-service/endpoints"
	"banking/internal/user-service/repositorys"
	"banking/internal/user-service/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = ""
	DB_NAME     = "banking"
)

func main() {
	connectString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_NAME)
	database.Connect(connectString)
	database.Migrate()
	r := route()

	grpcClient := grpc_client.NewGRPCAuthClient()
	authClient := grpcClient.SetUpCAuthClient()
	//DI User
	userRepo := repositorys.NewUserRepository(database.Instance)
	userUC := usecases.NewUserUsecase(userRepo)
	userCtrl := controllers.NewUserController(userUC)

	//DI Account
	accountRepo := repositorys.NewAccountRepository(database.Instance)
	accountUsecase := usecases.NewAccountUsecase(accountRepo, userRepo)
	accountCtr := controllers.NewAccountController(accountUsecase)

	//endpoint
	accountEndpoint := endpoints.NewAccountEndpoint(r, accountCtr, userCtrl, authClient)
	accountEndpoint.SetUp()
	if err := r.Run(":3030"); err != nil {
		log.Println("Connect to port fail", err)
	}
	log.Println("Connect to  DB success")
}

func route() *gin.Engine {
	r := gin.Default()
	return r
}
