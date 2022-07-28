package database

import (
	"banking/internal/authen-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		panic("Connect: Error connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	if err := Instance.AutoMigrate(&models.ProfileToken{}); err != nil {
		log.Println(err)
		return
	}
	log.Println("Database Migration Completed!")
}
