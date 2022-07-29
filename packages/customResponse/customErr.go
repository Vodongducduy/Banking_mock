package customResponse

import (
	"github.com/gin-gonic/gin"
	"log"
)

func FailErr(msg string, err error) {
	if err != nil {
		log.Printf("%s : %s \n", msg, err)
	}
}

func FailRespondAPI(c *gin.Context, status int, msg string, err interface{}) {
	if err != nil {
		c.JSON(status, gin.H{
			"Message": msg,
		})
		log.Println("Error to ", err)
		c.Abort()
	}
}
