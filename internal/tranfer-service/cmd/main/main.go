package main

import (
	"banking/internal/tranfer-service/cmd/qServer"
	"banking/internal/tranfer-service/database"
	"banking/internal/tranfer-service/repositorys"
	"banking/internal/tranfer-service/usecases"
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

	//DI Transfer
	tranferRepo := repositorys.NewTranferRepository(database.Instance)
	tranferUsecase := usecases.NewTranferUsecase(tranferRepo)
	tranferMq := qServer.NewTransferMQ(tranferUsecase)

	//MQServer
	go tranferMq.Consumer()

	if err := r.Run(":3031"); err != nil {
		log.Println("Connect to port fail", err)
	}
	log.Println("Connect to  DB success")
}

func route() *gin.Engine {
	r := gin.Default()
	return r
}
