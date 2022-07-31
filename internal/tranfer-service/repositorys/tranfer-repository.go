package repositorys

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/models"
	"banking/packages/customResponse"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ITranferRepository interface {
	CreateTranfer(tranfer *models.Tranfer) (*models.Tranfer, error)
	GetTranferById(id int) (*models.Tranfer, error)
	GetAllTranferByAccId(accountId int) (*[]models.Tranfer, error)
}

type TranferRepository struct {
	db *gorm.DB
}

func (r *TranferRepository) CreateTranfer(tranfer *models.Tranfer) (*models.Tranfer, error){
}

func (r *TranferRepository) GetTranferById(id int) (*models.Tranfer, error){
}

func (r *TranferRepository) GetAllTranferByAccId(accountId int) (*[]models.Tranfer, error){
}
