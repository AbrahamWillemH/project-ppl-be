package server

import (
	"fmt"
	v1 "project-ppl-be/api/v1"
	users "project-ppl-be/api/v1/users"

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
		v1Group.GET("/ping", v1.PingHandler)
		v1Group.GET("/users", users.UserGetHandler)
		v1Group.POST("/users", users.UserPostHandler)
	}

	fmt.Println("Server is running at http://localhost:8000")

	return router
}
