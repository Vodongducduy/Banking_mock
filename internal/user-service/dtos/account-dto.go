package dtos

type CreateAccountDTO struct {
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type GetAccountByIdDTO struct {
	AccountId int
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
