package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-api/repo"
)

func ReadUser(c *gin.Context) {
	log.Println("-------------------------CreateUser----------------------------")
	id := c.Param("id")
	log.Println("ID: " + id)

	data, err := repo.Read(id)
	if err != nil {
		c.AbortWithError(400, err)
	}

	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   data,
	})
}
