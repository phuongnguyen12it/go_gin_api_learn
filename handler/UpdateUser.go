package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-api/repo"
)

func UpdateUser(c *gin.Context) {
	log.Println("-------------------------Update User----------------------------")
	type myRequest struct {
		Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
		Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
		Password string `form:"password" json:"password" xml:"password"  binding:"required"`
		Note     string `form:"note" json:"note" xml:"note"`
	}
	id := c.Param("id")
	request := new(myRequest)
	err := c.Bind(&request)

	if err != nil {
		c.AbortWithError(400, err)
		log.Println(err)
	}
	err = repo.Update(id, request.Name, request.Email, request.Password, request.Note)

	if err != nil {
		c.AbortWithError(400, err)
	}

	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "User updated",
		"data":   request,
	})
}
