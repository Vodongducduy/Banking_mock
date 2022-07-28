package dtos

type CreateUserDTO struct {
	Name  string `json:"name" `
	Phone string `json:"phone" gorm:"unique"`
}

type GetUserByPhone struct {
	Phone string `json:"Phone"`
}
