package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler responds with "pong"
// @Summary Ping the server
// @Description A simple ping-pong endpoint
// @Tags HealthCheck
// @Accept  json
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /api/v1/ping [get]
func PingHandler(c *gin.Context) {
	fmt.Println("Ping endpoint was called")
	c.String(http.StatusOK, "pong")
}
