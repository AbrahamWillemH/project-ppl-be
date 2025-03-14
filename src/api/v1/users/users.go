package students

import (
	"context"
	"math"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userRepo = repo.UserRepository{}

// UserGetHandler retrieves a list of users
// @Summary Get Users
// @Description Fetch all users from the database with pagination, filtering by role, and sorting by username
// @Tags Users
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Param role query string false "Filter by role (e.g., 'admin', 'student')"
// @Param sortByUsername query bool false "Sort by username (true for ascending, false for descending)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/users [get]
func UserGetHandler(c *gin.Context) {
	// Ambil parameter query dari request
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	role := c.DefaultQuery("role", "")
	sortByUsername, _ := strconv.ParseBool(c.DefaultQuery("sortByUsername", "false"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	// Ambil data dengan pagination dan filter role
	users, total, err := userRepo.GetAllUsers(context.Background(), page, pageSize, role, sortByUsername)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format respons dengan metadata pagination
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"meta": gin.H{
			"page":      page,
			"pageSize":  pageSize,
			"total":     total,
			"totalPage": int(math.Ceil(float64(total) / float64(pageSize))),
		},
	})
}

// UserPostHandler creates a new user
// @Summary Create User
// @Description Create a new user in the database
// @Tags Users
// @Security BearerAuth
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

// UserUpdateHandler updates an existing user
// @Summary Update User
// @Description Updates an existing user in the database
// @Tags Users
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "User ID"
// @Param teacher body models.CreateUserRequest true "Updated User Data"
// @Success 200 {object} models.User
// @Router /api/v1/users [patch]
func UserUpdateHandler(c *gin.Context) {
	var req models.CreateUserRequest

	// Extract `id` from query parameters
	idStr := c.Query("id") // This will get the ID from query params
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user ID"})
		return
	}

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call UpdateUser with the correct parameters
	user, err := userRepo.UpdateUser(
		context.Background(),
		id,
		req.Username,
		req.Email,
		req.Password,
		req.Role,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated teacher
	c.JSON(http.StatusOK, user)
}
