package server

import (
	"fmt"
	"project-ppl-be/middleware"
	v1 "project-ppl-be/src/api/v1"
	auth "project-ppl-be/src/api/v1/auth"
	classes "project-ppl-be/src/api/v1/classes"
	discussions "project-ppl-be/src/api/v1/discussions"
	materials "project-ppl-be/src/api/v1/materials"
	students "project-ppl-be/src/api/v1/students"
	teachers "project-ppl-be/src/api/v1/teachers"
	users "project-ppl-be/src/api/v1/users"
	exercises "project-ppl-be/src/api/v1/exercises"
	exams "project-ppl-be/src/api/v1/exams"

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

		// USERS
		usersGroup := v1Group.Group("/users")
		usersGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		usersGroup.GET("", users.UserGetHandler)
		usersGroup.POST("", users.UserPostHandler)
		usersGroup.PATCH("", users.UserUpdateHandler)

		// STUDENTS
		studentsGroup := v1Group.Group("/students")
		studentsGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		studentsGroup.GET("", students.StudentsGetHandler)
		studentsGroup.POST("", students.StudentPostHandler)
		studentsGroup.PATCH("", students.StudentUpdateHandler)
		studentsGroup.DELETE("", students.StudentDeleteHandler)
		studentsGroup.POST("grade-migrate", students.StudentGradeMigrateHandler)

		// STUDENTS - NO ADMIN
		studentAccessGroup := v1Group.Group("/students")
		studentAccessGroup.Use(middleware.AuthMiddleware())
		studentAccessGroup.GET("/details", students.StudentGetByIDHandler)

		// TEACHERS
		teachersGroup := v1Group.Group("/teachers")
		teachersGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		teachersGroup.GET("", teachers.TeachersGetHandler)
		teachersGroup.POST("", teachers.TeachersPostHandler)
		teachersGroup.PATCH("", teachers.TeachersUpdateHandler)
		teachersGroup.DELETE("", teachers.TeachersDeleteHandler)

		// CLASSES
		classesGroup := v1Group.Group("/classes")
		classesGroup.Use(middleware.AuthMiddleware())
		classesGroup.GET("", classes.ClassGetHandler)
		classesGroup.GET("/class-id", classes.GetClassIDHandler)
		classesGroup.GET("/details", classes.ClassGetByIdHandler)
		classesGroup.POST("/assign-students", classes.ClassAssignStudentsHandler)
		classesGroup.DELETE("/unassign-students", classes.ClassUnassignStudentsHandler)
		classesGroup.POST("", classes.ClassPostHandler)
		classesGroup.PATCH("", classes.ClassUpdateHandler)
		classesGroup.DELETE("", classes.ClassDeleteHandler)

		// CLASSES FOR STUDENTS
		classesForStudentsGroup := v1Group.Group("/classes")
		classesForStudentsGroup.Use(middleware.AuthMiddleware())
		classesForStudentsGroup.GET("/assigned", classes.GetClassForStudentHandler)

		// MATERIALS
		materialsGroup := v1Group.Group("/materials")
		materialsGroup.Use(middleware.AuthMiddleware(), middleware.TeacherMiddleware())
		materialsGroup.GET("", materials.MaterialsGetHandler)
		materialsGroup.POST("", materials.MaterialsPostHandler)
		materialsGroup.PATCH("", materials.MaterialsUpdateHandler)
		materialsGroup.DELETE("", materials.MaterialsDeleteHandler)

		// MATERIALS NO ADMIN
		materialsAccessGroup := v1Group.Group("/materials")
		materialsAccessGroup.Use(middleware.AuthMiddleware())
		materialsAccessGroup.GET("/from-class", materials.MaterialsGetByClassIdHandler)

		// EXERCISES
		exercisesGroup := v1Group.Group("/exercises")
		exercisesGroup.Use(middleware.AuthMiddleware())
		exercisesGroup.GET("", exercises.ExercisesGetByMaterialHandler)
		exercisesGroup.GET("/student", exercises.ExercisesGetByMaterialForStudentHandler)
		exercisesGroup.POST("", exercises.ExercisesPostHandler)
		exercisesGroup.POST("/calculate-grade", exercises.CalculateGradePostHandler)
		exercisesGroup.GET("/get-grade", exercises.ExerciseGradesGetHandler)
		exercisesGroup.PATCH("", exercises.ExercisesUpdateHandler)
		exercisesGroup.DELETE("", exercises.ExercisesDeleteHandler)
		exercisesGroup.GET("/get-all-grade", exercises.ExerciseAllGradesGetHandler)

		// EXERCISE ANSWERS
		exerciseAnswersGroup := v1Group.Group("/exercises-answers")
		exerciseAnswersGroup.Use(middleware.AuthMiddleware())
		exerciseAnswersGroup.GET("", exercises.ExerciseAnswersGetHandler)
		exerciseAnswersGroup.POST("", exercises.ExerciseAnswersPostHandler)
		exerciseAnswersGroup.PATCH("", exercises.ExerciseAnswersUpdateHandler)
		exerciseAnswersGroup.DELETE("", exercises.ExerciseAnswersDeleteHandler)

		// EXAMS
		examsGroup := v1Group.Group("/exams")
		examsGroup.Use(middleware.AuthMiddleware())
		examsGroup.GET("", exams.ExamsGetByClassHandler)
		examsGroup.GET("/student", exams.ExamsGetByClassForStudentHandler)
		examsGroup.POST("", exams.ExamsPostHandler)
		examsGroup.POST("/calculate-grade", exams.CalculateGradePostHandler)
		examsGroup.GET("/get-grade", exams.ExamGradesGetHandler)
		examsGroup.PATCH("", exams.ExamsUpdateHandler)
		examsGroup.DELETE("", exams.ExamsDeleteHandler)
		examsGroup.GET("/get-all-grade", exams.ExamsAllGradesGetHandler)

		// EXERCISE ANSWERS
		examAnswersGroup := v1Group.Group("/exams-answers")
		examAnswersGroup.Use(middleware.AuthMiddleware())
		examAnswersGroup.GET("", exams.ExamAnswersGetHandler)
		examAnswersGroup.POST("", exams.ExamAnswersPostHandler)
		examAnswersGroup.PATCH("", exams.ExamAnswersUpdateHandler)
		examAnswersGroup.DELETE("", exams.ExamAnswersDeleteHandler)

		// DISCUSSIONS
		discussionsGroup := v1Group.Group("/discussions")
		discussionsGroup.Use(middleware.AuthMiddleware(), middleware.StudentMiddleware())
		discussionsGroup.GET("", discussions.DiscussionsGetHandler)
		discussionsGroup.POST("", discussions.DiscussionsPostHandler)
		discussionsGroup.PATCH("", discussions.DiscussionsUpdateHandler)
		discussionsGroup.PATCH("/reply", discussions.DiscussionsReplyHandler)
		discussionsGroup.DELETE("", discussions.DiscussionsDeleteHandler)
	}

	fmt.Println("Server is running at http://localhost:8080")
	return router
}
