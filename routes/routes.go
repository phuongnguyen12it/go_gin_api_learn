package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-api/handler"
	"user-api/middleware"
)

func Routes(router *gin.Engine) {
	log.Println("-------------------------routes----------------------------")
	publicAPI := router.Group("/", middleware.CheckToken)
	publicAPI.POST("/user/create", handler.CreateUser)
	publicAPI.PUT("/user/update/:id", handler.UpdateUser)
	publicAPI.DELETE("/user/delete/:id", handler.DeleteUser)
	publicAPI.GET("/user/:id", handler.ReadUser)
}
