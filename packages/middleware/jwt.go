package middleware

import (
	"banking/packages/customResponse"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	jwtKey            = []byte("DMM")
	ONE_HOUR_FROM_NOW = time.Now().Add(time.Hour * 1).Unix()
)

type JWTClaim struct {
	Phone       string
	AccountID   int
	Role        string
	TokenRandom int
	jwt.StandardClaims
}

func GenerateTokenJWT(phone string, accountID int, role string) (string, error) {
	claims := &JWTClaim{
		Phone:     phone,
		AccountID: accountID,
		Role:      role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ONE_HOUR_FROM_NOW,
		},
	}

	fmt.Println("claims", claims.Phone, claims.Role, claims.AccountID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		customResponse.FailErr("GenerateTokenJWT: Generate token fail to signed", err)
		return "", err
	}
	return tokenString, nil
}

func ExtractToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	fmt.Println("claims:---", token.Claims)
	customResponse.FailErr("ExtractToken: Fail to parse claims", err)
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	return claims, nil
}
