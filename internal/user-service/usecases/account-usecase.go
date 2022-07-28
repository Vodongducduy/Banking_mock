package usecases

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/models"
	"banking/internal/user-service/repositorys"
	"banking/packages/customResponse"
	"log"
)

type IAccountUsecase interface {
	CreateAccount(dto *dtos.CreateAccountDTO) (*models.Account, error)
	GetAccount(dto *dtos.GetAccountByIdDTO) (*models.Account, error)
}
type AccountUsecase struct {
	AccountRepository repositorys.IAccountRepository
	UserRepository    repositorys.IUserRepository
}

func (a *AccountUsecase) CreateAccount(dto *dtos.CreateAccountDTO) (*models.Account, error) {
	var account models.Account
	if err := account.HashPassword(dto.Password); err != nil {
		log.Println("HashPassword: Error to hash password pkg usecase", err)
		return nil, err
	}
	accountRep, errAcc := a.AccountRepository.CreateAccount(&account)
	if errAcc != nil {
		log.Println("CreateAccount: Error to create account pkg usecase", errAcc)
		return nil, errAcc
	}
	var user models.User
	user.Name = dto.Name
	user.Phone = dto.Phone
	user.AccountId = int(accountRep.ID)
	_, errUser := a.UserRepository.CreateUser(&user)
	if errUser != nil {
		log.Println("CreateUser: Error to create account pkg usecase", errUser)
		return nil, errUser
	}
	return accountRep, nil
}

func (a *AccountUsecase) GetAccount(dto *dtos.GetAccountByIdDTO) (*models.Account, error) {
	var dtoUser dtos.GetUserByPhone
	dtoUser.Phone = dto.Phone
	user, errUser := a.UserRepository.GetUser(&dtoUser)
	if errUser != nil {
		customResponse.FailErr("GetUser: Error to call repo pkg account-usecase", errUser)
		return nil, errUser
	}
	dto.AccountId = user.AccountId
	return a.AccountRepository.GetAccount(dto)

}

func NewAccountUsecase(accountRepository repositorys.IAccountRepository, userRepository repositorys.IUserRepository) *AccountUsecase {
	return &AccountUsecase{AccountRepository: accountRepository, UserRepository: userRepository}
}
