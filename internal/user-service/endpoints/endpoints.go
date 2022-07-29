package endpoints

import (
	"banking/internal/user-service/controllers"
	"banking/packages/pb/auth"
	"github.com/gin-gonic/gin"
)

type IAccountEndpoint interface {
	SetUp()
}

type AccountEndpoint struct {
	Route             *gin.Engine
	AccountController controllers.IAccountController
	UserController    controllers.IUserController
	authClient        auth.IsAuthClient
}

func NewAccountEndpoint(route *gin.Engine, accountController controllers.IAccountController, userController controllers.IUserController, authClient auth.IsAuthClient) *AccountEndpoint {
	return &AccountEndpoint{Route: route, AccountController: accountController, UserController: userController, authClient: authClient}
}

func (a *AccountEndpoint) SetUp() {

	accountApi := a.Route.Group("api/account")
	{
		accountApi.POST("/", a.AccountController.CreateAccount)
		accountApi.POST("/login", a.AccountController.GetAccount)
	}

	//setup middleware
	middlewareUser := controllers.NewMiddlewareUser(a.authClient)

	userApi := a.Route.Group("api/user")
	{
		userApi.GET("/", middlewareUser.CheckLogin(), a.UserController.GetAll)
	}

}
