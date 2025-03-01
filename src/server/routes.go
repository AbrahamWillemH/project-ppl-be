package server

import (
	"fmt"
	"project-ppl-be/middleware"
	v1 "project-ppl-be/src/api/v1"
	auth "project-ppl-be/src/api/v1/auth"
	students "project-ppl-be/src/api/v1/students"
	teachers "project-ppl-be/src/api/v1/teachers"
	users "project-ppl-be/src/api/v1/users"

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

		// Auth
		v1Group.POST("/auth", auth.AuthHandler)

		//Users group
		usersGroup := v1Group.Group("/users")
		usersGroup.Use(middleware.AuthMiddleware()) // Pasang middleware di sini
		usersGroup.GET("", users.UserGetHandler)
		usersGroup.POST("", users.UserPostHandler)

		//Students group
		studentsGroup := v1Group.Group("/students")
		studentsGroup.Use(middleware.AuthMiddleware()) // Pasang middleware di sini
		studentsGroup.GET("", students.StudentsGetHandler)
		studentsGroup.POST("", students.StudentPostHandler)

		//Teachers group
		teachersGroup := v1Group.Group("/teachers")
		teachersGroup.Use(middleware.AuthMiddleware()) // Pasang middleware di sini
		teachersGroup.GET("", teachers.TeachersGetHandler)
		teachersGroup.POST("", teachers.TeachersPostHandler)
	}

	fmt.Println("Server is running at http://localhost:8000")

	return router
}
