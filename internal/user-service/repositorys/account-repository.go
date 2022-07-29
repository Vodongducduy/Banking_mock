package repositorys

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/models"
	"banking/packages/customResponse"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IAccountRepository interface {
	CreateAccount(account *models.Account) (*models.Account, error)
	GetAccount(dto *dtos.GetAccountByIdDTO) (*models.Account, error)
	GetAccountById(id int) (*models.Account, error)
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
	var account *models.Account
	var count int64
	record := u.db.Where("id = ?", dto.AccountId).Find(&account).Count(&count)
	if record.Error != nil {
		customResponse.FailErr("GetAccount: Fail to find", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, nil
	}
	return account, nil
}

func (u *AccountRepository) GetAccountById(id int) (*models.Account, error) {
	var account models.Account
	var count int64
	err := u.db.Where("id = ?", id).Find(&account).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("record not found")
	}

	return &account, err
}
func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}
