package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler responds with "pong"
func PingHandler(c *gin.Context) {
	fmt.Println("Ping endpoint was called")
	c.String(http.StatusOK, "pong")
}
