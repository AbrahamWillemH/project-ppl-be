package server

import (
	"fmt"
	"project-ppl-be/middleware"
	v1 "project-ppl-be/src/api/v1"
	auth "project-ppl-be/src/api/v1/auth"
	classes "project-ppl-be/src/api/v1/classes"
	materials "project-ppl-be/src/api/v1/materials"
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
		v1Group.GET("/ping", v1.PingHandler)

		v1Group.POST("/auth", auth.AuthHandler)

		usersGroup := v1Group.Group("/users")
		usersGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		usersGroup.GET("", users.UserGetHandler)
		usersGroup.POST("", users.UserPostHandler)
		usersGroup.PATCH("", users.UserUpdateHandler)

		studentsGroup := v1Group.Group("/students")
		studentsGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		studentsGroup.GET("", students.StudentsGetHandler)
		studentsGroup.POST("", students.StudentPostHandler)
		studentsGroup.PATCH("", students.StudentUpdateHandler)
		studentsGroup.DELETE("", students.StudentDeleteHandler)
		studentsGroup.POST("grade-migrate", students.StudentGradeMigrateHandler)

		teachersGroup := v1Group.Group("/teachers")
		teachersGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		teachersGroup.GET("", teachers.TeachersGetHandler)
		teachersGroup.POST("", teachers.TeachersPostHandler)
		teachersGroup.PATCH("", teachers.TeachersUpdateHandler)
		teachersGroup.DELETE("", teachers.TeachersDeleteHandler)

		classesGroup := v1Group.Group("/classes")
		classesGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		classesGroup.GET("", classes.ClassGetHandler)
		classesGroup.POST("", classes.ClassPostHandler)
		classesGroup.PATCH("", classes.ClassUpdateHandler)
		classesGroup.DELETE("", classes.ClassDeleteHandler)

		materialsGroup := v1Group.Group("/materials")
		materialsGroup.Use(middleware.AuthMiddleware(), middleware.TeacherMiddleware())
		materialsGroup.GET("", materials.MaterialsGetHandler)
		materialsGroup.POST("", materials.MaterialsPostHandler)
		materialsGroup.PATCH("", materials.MaterialsUpdateHandler)
		materialsGroup.DELETE("", materials.MaterialsDeleteHandler)

		exercisesGroup := v1Group.Group("/exercises")
		exercisesGroup.Use(middleware.AuthMiddleware(), middleware.TeacherMiddleware())
		exercisesGroup.GET("", materials.MaterialsGetHandler)
		exercisesGroup.POST("", materials.MaterialsPostHandler)
		exercisesGroup.PATCH("", materials.MaterialsUpdateHandler)
		exercisesGroup.DELETE("", materials.MaterialsDeleteHandler)
	}

	fmt.Println("Server is running at http://localhost:8080")
	return router
}
