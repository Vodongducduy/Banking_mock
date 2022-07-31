package dtos

import "gorm.io/gorm"
type TranferDTO struct{
	FromAccountID int `json:"from-account-id"`
	ToAccountID int `json:"to-account-id"`
	Amount int `json:"amount"`
}
