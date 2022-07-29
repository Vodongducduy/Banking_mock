package usecases

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/models"
	"banking/internal/user-service/repositorys"
)

type IUserUsecase interface {
	GetUser(dto *dtos.GetUserByPhone) (*models.User, error)
	GetAll() (*[]models.User, error)
}

type UserUsecase struct {
	UserRepository repositorys.IUserRepository
}

func NewUserUsecase(userRepository repositorys.IUserRepository) *UserUsecase {
	return &UserUsecase{UserRepository: userRepository}
}

func (u *UserUsecase) GetUser(dto *dtos.GetUserByPhone) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUsecase) GetAll() (*[]models.User, error) {
	return u.UserRepository.GetAll()
}
