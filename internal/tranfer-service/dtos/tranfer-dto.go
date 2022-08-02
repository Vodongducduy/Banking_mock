package dtos

type TranferDTO struct {
	FromAccountID int `json:"from-account-id"`
	ToAccountID   int `json:"to-account-id"`
	Amount        int `json:"amount"`
}
