package main

import (
	"github.com/gin-gonic/gin"
	"backend-demo/handlers"
	"backend-demo/middleware"
)

func main() {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handlers.Profile)
	}


	r.Run("localhost:8888")
}