package server

import (
	"fmt"
	v1 "project-ppl-be/api/v1"
	"project-ppl-be/api/v1/auth"
	"project-ppl-be/api/v1/users"
	"project-ppl-be/middleware" // Import middleware

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API Version 1 group
	v1Group := router.Group("/api/v1")
	{
		// Test
		v1Group.GET("/ping", v1.PingHandler)

		// Users - menggunakan middleware AuthMiddleware
		usersGroup := v1Group.Group("/users")
		usersGroup.Use(middleware.AuthMiddleware()) // Middleware diterapkan di sini
		{
			usersGroup.GET("", users.UserGetHandler)
			usersGroup.POST("", users.UserPostHandler)
		}

		// Auth
		v1Group.POST("/auth", auth.AuthHandler)
	}

	fmt.Println("Server is running at http://localhost:8000")

	return router
}
