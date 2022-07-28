package customResponse

import "github.com/gin-gonic/gin"

func SuccessRespondAPI(c *gin.Context, status int, msg interface{}) {
	c.JSON(status, gin.H{
		"Message": msg,
	})
	c.Abort()
	return
}
