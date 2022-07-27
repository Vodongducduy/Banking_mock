package dtos

type CreateUserDTO struct {
	Name  string `json:"name" `
	Phone string `json:"phone" gorm:"unique"`
}

type GetUserByID struct {
	UserId int `json:"user-id"`
}
