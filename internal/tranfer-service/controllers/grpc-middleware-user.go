package controllers

import (
	"banking/packages/config"
	"banking/packages/pb/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IMiddlewareUser interface {
	CheckLogin() gin.HandlerFunc
}
type MiddlewareUser struct {
	authClient auth.IsAuthClient
}

func NewMiddlewareUser(authClient auth.IsAuthClient) *MiddlewareUser {
	return &MiddlewareUser{authClient: authClient}
}

const (
	HEADER_KEY_AUTHORIZATION = "Authorization"
	NOT_CONTAIN_ACCESS_TOKEN = "request does not contain an access token"
)

func (m *MiddlewareUser) CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No token found"})
			return
		}
		fmt.Println(tokenString)

		//set token to req
		req := &auth.IsAuthRequest{
			Token: tokenString,
		}

		res, err := m.authClient.IsAuth(c, req)

		fmt.Println("res", res)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}

		c.Set(config.Role, res.Role)
		c.Set(config.AccountId, res.AccountId)
		c.Next()
	}
}

//func getContext(c *gin.Context) (context.Context, context.CancelFunc) {
//	grpc_ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//
//	token := c.GetHeader(HEADER_KEY_AUTHORIZATION)
//	if token == "" {
//		c.JSON(http.StatusUnauthorized, gin.H{
//			"error": NOT_CONTAIN_ACCESS_TOKEN,
//		})
//		return nil, cancel
//	}
//	grpc_ctx = grpcMetadata.AppendToOutgoingContext(grpc_ctx, "authorization", token)
//
//	return grpc_ctx, cancel
//}
