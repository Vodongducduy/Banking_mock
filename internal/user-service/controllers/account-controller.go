package controllers

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/usecases"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAccountController interface {
	CreateAccount(c *gin.Context)
}

type AccountController struct {
	AccountUsecase usecases.IAccountUsecase
}

func (a *AccountController) CreateAccount(c *gin.Context) {
	var dto dtos.CreateAccountDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create account",
		})
		log.Println("CreateAccount: Error to ShouldBindJSON", err)
		c.Abort()
		return
	}
	_, errAccount := a.AccountUsecase.CreateAccount(&dto)
	if errAccount != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create account",
		})
		log.Println("CreateAccount: Error to call usecase", errAccount)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "Create account successfully",
	})
}

func NewAccountController(accountUsecase usecases.IAccountUsecase) *AccountController {
	return &AccountController{AccountUsecase: accountUsecase}
}
