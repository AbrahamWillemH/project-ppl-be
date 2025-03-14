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

		// Users group
		usersGroup := v1Group.Group("/users")
		usersGroup.Use(middleware.AuthMiddleware())
		usersGroup.GET("", users.UserGetHandler)
		usersGroup.POST("", users.UserPostHandler)
		usersGroup.PATCH("", users.UserUpdateHandler)

		// Students group
		studentsGroup := v1Group.Group("/students")
		studentsGroup.Use(middleware.AuthMiddleware())
		studentsGroup.GET("", students.StudentsGetHandler)
		studentsGroup.POST("", students.StudentPostHandler)
		studentsGroup.PATCH("", students.StudentUpdateHandler)
		studentsGroup.DELETE("", students.StudentDeleteHandler)

		// Teachers group
		teachersGroup := v1Group.Group("/teachers")
		teachersGroup.Use(middleware.AuthMiddleware())
		teachersGroup.GET("", teachers.TeachersGetHandler)
		teachersGroup.POST("", teachers.TeachersPostHandler)
		teachersGroup.PATCH("", teachers.TeachersUpdateHandler)
		teachersGroup.DELETE("", teachers.TeachersDeleteHandler)

		// Materials group
		materialsGroup := v1Group.Group("/materials")
		materialsGroup.Use(middleware.AuthMiddleware())
	}

	fmt.Println("Server is running at http://localhost:8080")

	return router
}
