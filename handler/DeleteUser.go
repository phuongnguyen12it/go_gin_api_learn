package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-api/repo"
)

func DeleteUser(c *gin.Context) {
	log.Println("-------------------------Delete User----------------------------")

	id := c.Param("id")
	err := repo.Delete(id)

	if err != nil {
		c.AbortWithError(400, err)
	}

	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "User deleted",
	})
}
