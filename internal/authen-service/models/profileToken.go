package models

import "gorm.io/gorm"

type ProfileToken struct {
	gorm.Model
	Token string `json:"token"`
}
