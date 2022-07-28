package usecases

import (
	"banking/internal/authen-service/repositorys"
	"banking/packages/customResponse"
	"banking/packages/middleware"
	"banking/packages/pb/auth"
	"errors"
	"time"
)

type IAuthUsecase interface {
	AddToken(token string) error
	IsUser(req *auth.IsAuthRequest) (*auth.IsAuthResponse, error)
}

type AuthUsecase struct {
	AuthRepository repositorys.IAuthRepository
}

func (a *AuthUsecase) IsUser(req *auth.IsAuthRequest) (*auth.IsAuthResponse, error) {
	tokenString := req.Token
	//if err := a.AddToken(tokenString); err != nil {
	//	customResponse.FailErr("IsUser: add token fail", err)
	//}
	claims, err := middleware.ExtractToken(tokenString)
	if claims == nil {
		customResponse.FailErr("ExtractToken: Error", err)
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return nil, err
	}

	return &auth.IsAuthResponse{
		Role:      claims.Role,
		AccountId: int32(claims.AccountID),
	}, err
}

func (a *AuthUsecase) AddToken(token string) error {
	return a.AuthRepository.AddToken(token)
}

func NewAuthUsecase(authRepository repositorys.IAuthRepository) *AuthUsecase {
	return &AuthUsecase{AuthRepository: authRepository}
}
