package repositorys

import (
	"banking/internal/authen-service/models"
	"banking/packages/customResponse"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	AddToken(token string) error
	//(token string) error
}

type AuthRepository struct {
	db *gorm.DB
}

func (a *AuthRepository) AddToken(token string) error {
	var storedToken *models.ProfileToken
	storedToken.Token = token
	record := a.db.Create(&storedToken)
	if record.Error != nil {
		customResponse.FailErr("IsValidToken: Error to store token", record.Error)
		return record.Error
	}
	return nil
}

//func (a *AuthRepository) CheckTokenExist(token string) error{
//	var
//	return a.db.Where("token = ?", token).Find()
//}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
