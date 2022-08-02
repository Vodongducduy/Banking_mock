package usecases

import (
	"banking/internal/tranfer-service/models"
	"banking/internal/tranfer-service/repositorys"
	"banking/packages/customResponse"
	"banking/packages/middleware"
)

type ITranferUsecase interface {
	CreateTranfer(msgToken string) error
	GetTranferById(id int) (*models.Tranfer, error)
	GetAllTranferByAccId(accountId int) (*[]models.Tranfer, error)
}

type TranferUsecase struct {
	TranferRepository repositorys.ITranferRepository
}

func (t *TranferUsecase) CreateTranfer(msgToken string) error {
	claims, err := middleware.ExtractTokenTransfer(msgToken)
	if err != nil {
		customResponse.FailErr("Fail to claims", err)
		return nil
	}
	dto := &models.Tranfer{
		FromAccountID: claims.TransferInfo.FromAccountID,
		ToAccountID:   claims.TransferInfo.ToAccountID,
		Amount:        claims.TransferInfo.Amount,
	}
	_, errCreate := t.TranferRepository.CreateTranfer(dto)
	if errCreate != nil {
		customResponse.FailErr("Fail to Create Transfer", errCreate)
		return err
	}
	return nil
}

func (t TranferUsecase) GetTranferById(id int) (*models.Tranfer, error) {
	//TODO implement me
	panic("implement me")
}

func (t TranferUsecase) GetAllTranferByAccId(accountId int) (*[]models.Tranfer, error) {
	//TODO implement me
	panic("implement me")
}

func NewTranferUsecase(tranferRepository repositorys.ITranferRepository) *TranferUsecase {
	return &TranferUsecase{TranferRepository: tranferRepository}
}
