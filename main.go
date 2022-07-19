package main

import (
	"votar/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/signin", controllers.SignIn)

	router.Run(":8000")
}
