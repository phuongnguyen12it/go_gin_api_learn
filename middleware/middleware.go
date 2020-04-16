package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckToken(c *gin.Context) {
	log.Println("-------------------------middleware----------------------------")
	if c.GetHeader("token") == "token_value" {
		c.Next()
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "token is invalid"})
		c.Abort()
		return
	}
}
