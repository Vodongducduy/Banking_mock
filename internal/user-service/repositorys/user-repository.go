package repositorys

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/models"
	"banking/packages/customResponse"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(dto *dtos.GetUserByPhone) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	record := u.db.Create(&user)
	if record.Error != nil {
		log.Println("CreateUser: Error to create user", record.Error)
		return nil, record.Error
	}
	return user, nil
}

func (u *UserRepository) GetUser(dto *dtos.GetUserByPhone) (*models.User, error) {
	var user *models.User
	record := u.db.Where("phone = ?", dto.Phone).Find(&user)
	if record.Error != nil {
		customResponse.FailErr("GetUser: Error to find user by phone", record.Error)
		return nil, record.Error
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
