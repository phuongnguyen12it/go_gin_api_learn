package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	// "user-api/model"
	"user-api/repo"
)

func CreateUser(c *gin.Context) {
	log.Println("-------------------------CreateUser----------------------------")
	type myRequest struct {
		Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
		Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
		Password string `form:"password" json:"password" xml:"password"  binding:"required"`
		Note     string `form:"note" json:"note" xml:"note"`
	}
	request := new(myRequest)
	err := c.Bind(&request)

	if err != nil {
		c.AbortWithError(400, err)
		log.Println(err)
	}
	// user := model.User{}
	err = repo.Add(request.Name, request.Email, request.Password, request.Note)

	if err != nil {
		c.AbortWithError(400, err)
	}

	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "User created",
		"data":   request,
	})
}
