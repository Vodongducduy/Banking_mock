package repositorys

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/models"
	"gorm.io/gorm"
	"log"
)

type IAccountRepository interface {
	CreateAccount(account *models.Account) (*models.Account, error)
	GetAccount(dto *dtos.GetAccountByIdDTO) (*models.Account, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func (u *AccountRepository) CreateAccount(account *models.Account) (*models.Account, error) {
	record := u.db.Create(&account)
	if record.Error != nil {
		log.Println("CreateAccount: Error to create Account", record.Error)
		return nil, record.Error
	}
	return account, nil
}

func (u *AccountRepository) GetAccount(dto *dtos.GetAccountByIdDTO) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}
