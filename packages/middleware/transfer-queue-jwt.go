package middleware

import (
	"banking/internal/tranfer-service/dtos"
	"banking/packages/customResponse"
	"errors"
	"github.com/golang-jwt/jwt"
)

type JWTTransferClaim struct {
	TransferInfo *dtos.TranferDTO
	jwt.StandardClaims
}

func GenerateTokenTransferJWT(TransferInfo *dtos.TranferDTO) (string, error) {
	claims := &JWTTransferClaim{
		TransferInfo: TransferInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ONE_HOUR_FROM_NOW,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	customResponse.FailErr("GenerateTokenJWT: Generate token fail to signed", err)
	return tokenString, err
}

func ExtractTokenTransfer(signedToken string) (*JWTTransferClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTTransferClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	customResponse.FailErr("ExtractToken: Fail to parse claims", err)
	claims, ok := token.Claims.(*JWTTransferClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	return claims, nil
}
