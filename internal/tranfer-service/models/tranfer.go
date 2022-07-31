package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Tranfer struct{
	gorm.Model `json:"-"`
	FromAccountID int `json:"from-account-id"`
	ToAccountID int `json:"to-account-id"`
	Amount int `json:"amount"`
}
