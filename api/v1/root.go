package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootHandler handles the root endpoint
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Gin!",
	})
}
