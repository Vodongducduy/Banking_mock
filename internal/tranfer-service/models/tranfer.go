package models

import "gorm.io/gorm"

type Tranfer struct {
	gorm.Model    `json:"-"`
	FromAccountID int `json:"from-account-id"`
	ToAccountID   int `json:"to-account-id"`
	Amount        int `json:"amount"`
}
