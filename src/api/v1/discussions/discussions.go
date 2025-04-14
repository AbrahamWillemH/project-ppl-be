package discussions

import (
	"context"
	"math"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var discussionsRepo = repo.ClassRepository{}

// ClassGetHandler retrieves a list of discussions
// @Summary Get Class
// @Description Fetch all discussions from the database with pagination
// @Tags discussions
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/discussions [get]
func ClassGetHandler(c *gin.Context) {
	// Ambil parameter query dari request
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	// Ambil data dengan pagination dan filter grade
	discussions, total, err := discussionsRepo.GettAllDiscussions(context.Background(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format respons dengan metadata pagination
	c.JSON(http.StatusOK, gin.H{
		"discussions": discussions,
		"meta": gin.H{
			"page":      page,
			"pageSize":  pageSize,
			"total":     total,
			"totalPage": int(math.Ceil(float64(total) / float64(pageSize))),
		},
	})
}

// ClassPostHandler creates a new class
// @Summary Create Class
// @Description Create a new class in the database
// @Tags Classes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.CreateClassRequest true "Class data"
// @Success 200 {object} models.Class
// @Router /api/v1/discussions [post]
func ClassPostHandler(c *gin.Context) {
	var req models.CreateClassRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateUser with the extracted values
	user, err := discussionsRepo.CreateClass(context.Background(), req.Name, req.Description, req.Teacher_ID, req.Grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, user)
}

// ClassUpdateHandler updates an existing class
// @Summary Update Class
// @Description Updates an existing class in the database
// @Tags Classes
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Class ID"
// @Param class body models.CreateClassRequest true "Updated Class Data"
// @Success 200 {object} models.Class
// @Router /api/v1/discussions [patch]
func ClassUpdateHandler(c *gin.Context) {
	var req models.CreateClassRequest

	// Extract `id` from query parameters
	idStr := c.Query("id") // This will get the ID from query params
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing class ID"})
		return
	}

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call UpdateTeacher with the correct parameters
	class, err := discussionsRepo.UpdateClass(
		context.Background(),
		id,
		req.Name,
		req.Description,
		req.Teacher_ID,
		req.Grade,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated teacher
	c.JSON(http.StatusOK, class)
}

// ClassDeleteHandler deletes a class
// @Summary Delete Class
// @Description Deletes a class from the database by ID
// @Tags Classes
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Class ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/discussions [delete]
func ClassDeleteHandler(c *gin.Context) {
	// Extract `id` from query parameters
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing teacher ID"})
		return
	}

	// Call DeleteTeacher function from repository
	err = discussionsRepo.DeleteClass(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
