package server

import (
	"fmt"

	v1 "project-ppl-be/api/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1Group := router.Group("/api/v1")

	v1Group.GET("/", v1.RootHandler)
	v1Group.GET("/ping", v1.PingHandler)
	v1Group.GET("/user/:name", v1.UserHandler)

	fmt.Println("Server is running at http://localhost:8000")
	return router
}
