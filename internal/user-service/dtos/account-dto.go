package dtos

type CreateAccountDTO struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}

type GetAccountByIdDTO struct {
	AccountId int
}
