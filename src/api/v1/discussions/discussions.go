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

var discussionsRepo = repo.DiscussionRepository{}

// DiscussionGetHandler retrieves a list of discussions
// @Summary Get Discussion
// @Description Fetch all discussions from the database with pagination
// @Tags discussions
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/discussions [get]
func DiscussionsGetHandler(c *gin.Context) {
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

// DiscussionPostHandler creates a new discussion
// @Summary Create Discussion
// @Description Create a new discussion in the database
// @Tags Discussiones
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.CreateDiscussionRequest true "Discussion data"
// @Success 200 {object} models.Discussion
// @Router /api/v1/discussions [post]
func DiscussionPostHandler(c *gin.Context) {
	var req models.CreateDiscussionRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateUser with the extracted values
	user, err := discussionsRepo.CreateDiscussion(context.Background(), req.Name, req.Description, req.Teacher_ID, req.Grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, user)
}

// DiscussionUpdateHandler updates an existing discussion
// @Summary Update Discussion
// @Description Updates an existing discussion in the database
// @Tags Discussiones
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Discussion ID"
// @Param discussion body models.CreateDiscussionRequest true "Updated Discussion Data"
// @Success 200 {object} models.Discussion
// @Router /api/v1/discussions [patch]
func DiscussionUpdateHandler(c *gin.Context) {
	var req models.CreateDiscussionRequest

	// Extract `id` from query parameters
	idStr := c.Query("id") // This will get the ID from query params
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing discussion ID"})
		return
	}

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call UpdateTeacher with the correct parameters
	discussion, err := discussionsRepo.UpdateDiscussion(
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
	c.JSON(http.StatusOK, discussion)
}

// DiscussionDeleteHandler deletes a discussion
// @Summary Delete Discussion
// @Description Deletes a discussion from the database by ID
// @Tags Discussiones
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Discussion ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/discussions [delete]
func DiscussionDeleteHandler(c *gin.Context) {
	// Extract `id` from query parameters
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing teacher ID"})
		return
	}

	// Call DeleteTeacher function from repository
	err = discussionsRepo.DeleteDiscussion(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Discussion deleted successfully"})
}
