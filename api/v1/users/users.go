package users

import (
	"context"
	"net/http"
	"project-ppl-be/models" // Import models
	"project-ppl-be/repo"

	"github.com/gin-gonic/gin"
)

var userRepo = repo.UserRepository{}

// UserGetHandler retrieves a list of users
// @Summary Get Users
// @Description Fetch all users from the database
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /api/v1/users [get]
func UserGetHandler(c *gin.Context) {
	users, err := userRepo.GetAllUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UserPostHandler creates a new user
// @Summary Create User
// @Description Create a new user in the database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.CreateUserRequest true "User data"
// @Success 200 {object} models.User
// @Router /api/v1/users [post]
func UserPostHandler(c *gin.Context) {
	var req models.CreateUserRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateUser with the extracted values
	user, err := userRepo.CreateUser(context.Background(), req.Username, req.Email, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, user)
}
