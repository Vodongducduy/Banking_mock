package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" `
	Address    string `json:"address"`
	Phone      string `json:"phone" gorm:"unique"`
	Email      string `json:"email" gorm:"type:varchar(50)"`
	AccountId  int    `json:"account-id" gorm:"unique"`
}
