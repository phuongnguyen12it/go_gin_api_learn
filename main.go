package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-api/routes"
)

func main() {
	log.Println("-------------------------main----------------------------")
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":8080"))
}
