package main

import (
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
	//DI User
	userRepo := repositorys.NewUserRepository(database.Instance)

	//DI Account
	accountRepo := repositorys.NewAccountRepository(database.Instance)
	accountUsecase := usecases.NewAccountUsecase(accountRepo, userRepo)
	accountCtr := controllers.NewAccountController(accountUsecase)
	accountEndpoint := endpoints.NewAccountEndpoint(r, accountCtr)
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
