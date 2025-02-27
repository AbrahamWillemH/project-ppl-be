package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define a basic route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin!",
		})
	})

	// Ping test route
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("Ping endpoint was called") // Print in terminal
		c.String(http.StatusOK, "pong")
	})

	// Dynamic route with parameters
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"user":    name,
			"message": "Hello, " + name + "!",
		})
	})

	// Run the server on port 8080
	fmt.Println("Server is running at http://localhost:8000")
	r.Run(":8000")
}
