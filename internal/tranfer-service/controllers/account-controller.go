package controllers

import (
	"banking/internal/user-service/dtos"
	"banking/internal/user-service/usecases"
	"banking/packages/customResponse"
	"banking/packages/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAccountController interface {
	CreateAccount(c *gin.Context)
	GetAccount(c *gin.Context)
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

func (a *AccountController) GetAccount(c *gin.Context) {
	var dto dtos.GetAccountByIdDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		customResponse.FailRespondAPI(c, http.StatusBadRequest, "GetAccount: Fail to ShouldBindJSON", err)
		return
	}
	account, errAccount := a.AccountUsecase.GetAccount(&dto)
	if errAccount != nil {
		customResponse.FailRespondAPI(c, http.StatusBadRequest, "GetAccount: Fail to get account", errAccount)
		return
	}

	//Account not found
	if account == nil {
		customResponse.FailRespondAPI(c, http.StatusBadRequest, "Not found phone", errors.New("Not fouund bro"))
		return
	}

	//account.Password = dto.Password

	//Wrong password
	if err := account.CheckPassword(dto.Password); err != nil {
		customResponse.FailRespondAPI(c, http.StatusBadRequest, "Phone or password is invalid", err)
		return
	}
	tokenString, errToken := middleware.GenerateTokenJWT(dto.Phone, int(account.ID), account.Role)
	if errToken != nil {
		customResponse.FailRespondAPI(c, http.StatusInternalServerError, "Generate token fail", errToken)
		return
	}
	customResponse.SuccessRespondAPI(c, http.StatusOK, tokenString)
}

func NewAccountController(accountUsecase usecases.IAccountUsecase) *AccountController {
	return &AccountController{AccountUsecase: accountUsecase}
}
