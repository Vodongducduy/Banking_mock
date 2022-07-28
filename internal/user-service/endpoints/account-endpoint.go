package endpoints

import (
	"banking/internal/user-service/controllers"
	"github.com/gin-gonic/gin"
)

type IAccountEndpoint interface {
	SetUp()
}

type AccountEndpoint struct {
	Route             *gin.Engine
	AccountController controllers.IAccountController
}

func (a *AccountEndpoint) SetUp() {

	accountApi := a.Route.Group("api/account")
	{
		accountApi.POST("/", a.AccountController.CreateAccount)
		accountApi.POST("/login", a.AccountController.GetAccount)
	}

}

func NewAccountEndpoint(route *gin.Engine, accountController controllers.IAccountController) *AccountEndpoint {
	return &AccountEndpoint{Route: route, AccountController: accountController}
}
