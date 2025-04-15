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

// DiscussionsGetHandler retrieves a list of discussions
// @Summary Get Discussions
// @Description Fetch all discussions from the database with pagination
// @Tags Discussions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/discussions [get]
func DiscussionsGetHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	discussions, total, err := discussionsRepo.GetAllDiscussions(context.Background(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
// @Tags Discussions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param discussion body models.CreateDiscussionRequest true "Discussion data"
// @Success 200 {object} models.Discussion
// @Router /api/v1/discussions [post]
func DiscussionsPostHandler(c *gin.Context) {
	var req models.CreateDiscussionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discussion, err := discussionsRepo.CreateDiscussion(
		context.Background(),
		req.Student_ID,
		req.Topic,
		req.Description,
		req.Replies,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discussion)
}

// DiscussionUpdateHandler updates an existing discussion
// @Summary Update Discussion
// @Description Updates an existing discussion in the database
// @Tags Discussions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Discussion ID"
// @Param discussion body models.CreateDiscussionRequest true "Updated Discussion Data"
// @Success 200 {object} models.Discussion
// @Router /api/v1/discussions [patch]
func DiscussionsUpdateHandler(c *gin.Context) {
	var req models.CreateDiscussionRequest

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing discussion ID"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discussion, err := discussionsRepo.UpdateDiscussion(
		context.Background(),
		id,
		req.Student_ID,
		req.Topic,
		req.Description,
		req.Replies,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discussion)
}

// DiscussionDeleteHandler deletes a discussion
// @Summary Delete Discussion
// @Description Deletes a discussion from the database by ID
// @Tags Discussions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Discussion ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/discussions [delete]
func DiscussionsDeleteHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing discussion ID"})
		return
	}

	err = discussionsRepo.DeleteDiscussion(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discussion deleted successfully"})
}
