package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Account struct {
	gorm.Model    `json:"-"`
	AccountNumber string `json:"account-number"`
	Balance		int 	`json:"-" gorm:"default:0"`
	Password      string `json:"password"`
	Role          string `json:"role-id" gorm:"default:user"`
}

func (a *Account) HashPassword(password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("HashPassword: Error hash password", err)
		return err
	}
	a.Password = string(hashPassword)
	return nil
}

func (a *Account) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		log.Println("CheckPassword: Error hash password", err)
		return err
	}
	return nil
}
