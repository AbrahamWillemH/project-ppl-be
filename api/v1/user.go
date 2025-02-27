package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles requests to "/user/:name"
func UserHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"user":    name,
		"message": "Hello, " + name + "!",
	})
}
