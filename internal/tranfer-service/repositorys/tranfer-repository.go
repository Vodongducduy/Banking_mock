package repositorys

import (
	"banking/internal/tranfer-service/models"
	"banking/packages/customResponse"
	"gorm.io/gorm"
)

type ITranferRepository interface {
	CreateTranfer(tranfer *models.Tranfer) (*models.Tranfer, error)
	GetTranferById(id int) (*models.Tranfer, error)
	GetAllTranferByAccId(accountId int) (*[]models.Tranfer, error)
}

type TranferRepository struct {
	db *gorm.DB
}

func NewTranferRepository(db *gorm.DB) *TranferRepository {
	return &TranferRepository{db: db}
}

func (r *TranferRepository) CreateTranfer(tranfer *models.Tranfer) (*models.Tranfer, error) {
	record := r.db.Create(&tranfer)
	if record.Error != nil {
		customResponse.FailErr("Error to Create Tranfer", record.Error)
		return nil, record.Error
	}
	return tranfer, nil
}

func (r *TranferRepository) GetTranferById(id int) (*models.Tranfer, error) {
	return nil, nil
}

func (r *TranferRepository) GetAllTranferByAccId(accountId int) (*[]models.Tranfer, error) {
	return nil, nil
}
