package main

import (
	"github.com/gin-gonic/gin"
	"backend-demo/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.Run("localhost:8888")
}