package materials

import (
	"context"
	"math"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var materialsRepo = repo.MaterialRepository{}

// MaterialsGetHandler retrieves a list of materials
// @Summary Get Materials
// @Description Fetch all materials from the database with pagination
// @Tags Materials
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/materials [get]
func MaterialsGetHandler(c *gin.Context) {
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
	materials, total, err := materialsRepo.GetAllMaterials(context.Background(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format respons dengan metadata pagination
	c.JSON(http.StatusOK, gin.H{
		"materials": materials,
		"meta": gin.H{
			"page":      page,
			"pageSize":  pageSize,
			"total":     total,
			"totalPage": int(math.Ceil(float64(total) / float64(pageSize))),
		},
	})
}

// MaterialsPostHandler creates a new material
// @Summary Create Material
// @Description Create a new material in the database
// @Tags Materials
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.CreateMaterialRequest true "Material data"
// @Success 200 {object} models.Material
// @Router /api/v1/materials [post]
func MaterialsPostHandler(c *gin.Context) {
	var req models.CreateMaterialRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateUser with the extracted values
	user, err := materialsRepo.CreateMaterial(context.Background(), req.Class_ID, req.Title, req.Description, req.Content, req.Teacher_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, user)
}

// MaterialsUpdateHandler updates an existing material
// @Summary Update Material
// @Description Updates an existing material in the database
// @Tags Materials
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Material ID"
// @Param material body models.UpdateMaterialRequest true "Updated Material Data"
// @Success 200 {object} models.Material
// @Router /api/v1/materials [patch]
func MaterialsUpdateHandler(c *gin.Context) {
	var req models.UpdateMaterialRequest

	// Extract `id` from query parameters
	idStr := c.Query("id") // This will get the ID from query params
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing material ID"})
		return
	}

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call UpdateMaterial with the correct parameters
	material, err := materialsRepo.UpdateMaterial(
		context.Background(),
		id,
		req.Class_ID,
		req.Title,
		req.Description,
		req.Content,
		req.Teacher_ID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated teacher
	c.JSON(http.StatusOK, material)
}

// MaterialsDeleteHandler deletes a material
// @Summary Delete Material
// @Description Deletes a material from the database by ID
// @Tags Materials
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Material ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/materials [delete]
func MaterialsDeleteHandler(c *gin.Context) {
	// Extract `id` from query parameters
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing teacher ID"})
		return
	}

	// Call DeleteTeacher function from repository
	err = materialsRepo.DeleteMaterial(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}
