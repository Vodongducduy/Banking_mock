package controllers

import (
	"banking/internal/tranfer-service/dtos"
	"banking/internal/user-service/cmd/producer"
	"banking/internal/user-service/usecases"
	"banking/packages/config"
	"banking/packages/customResponse"
	"banking/packages/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserController interface {
	GetAll(c *gin.Context)
	UserTransfer(c *gin.Context)
}

type UserController struct {
	userUsecase    usecases.IUserUsecase
	accountUsecase usecases.IAccountUsecase
}

func (u *UserController) UserTransfer(c *gin.Context) {
	var dto dtos.TranferDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		customResponse.FailRespondAPI(c, http.StatusBadRequest, "Bad input", "Fail to ShouldBindJSON")
		return
	}
	tokenString, err := middleware.GenerateTokenTransferJWT(&dto)
	if err != nil {
		customResponse.FailRespondAPI(c, http.StatusBadRequest, "Bad input", "Fail to GenerateTokenTransferJWT")
		return
	}
	producerUser := producer.NewUseProducer()
	if err := producerUser.ProducerTransfer(tokenString); err != nil {
		customResponse.FailRespondAPI(c, http.StatusUnauthorized, "Bad input", "Fail to GenerateTokenTransferJWT")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully transfer",
	})
}

func (u *UserController) GetAll(c *gin.Context) {
	users, err := u.userUsecase.GetAll()
	roleToken, ok := c.Get(config.Role)
	if !ok {
		customResponse.FailRespondAPI(c, http.StatusInternalServerError, "Server Error", "GetAll: "+err.Error())
		return
	}

	if (roleToken).(string) != "user" {
		customResponse.FailRespondAPI(c, http.StatusInternalServerError, "Server Error", "GetAll:")
		return
		//c.JSON(http.StatusForbidden, gin.H{
		//	"Message": "Forbidden accept",
		//})
		//c.Abort()
		//return
	}
	if err != nil {
		customResponse.FailRespondAPI(c, http.StatusInternalServerError, "Server Error", "GetAll: "+err.Error())
		return
	}
	customResponse.SuccessRespondAPI(c, http.StatusOK, users)

}

//func (u *UserController) checkRole(accountId, roleToken string) error {
//	account := u.userUsecase.
//}

func NewUserController(userUsecase usecases.IUserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}
