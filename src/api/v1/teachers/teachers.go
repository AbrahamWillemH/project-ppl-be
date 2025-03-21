package teachers

import (
	"context"
	"math"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var teacherRepo = repo.TeacherRepository{}

// TeachersGetHandler retrieves a list of students
// @Summary Get Teachers
// @Description Fetch all teachers from the database with pagination, filtering by specialization, and sorting by NIP
// @Tags Teachers
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Param specialization query string false "Filter by specialization (e.g., 'IPA')"
// @Param sortByNIP query bool false "Sort by NIP (true for ascending, false for descending)"
// @Param search query string false "Search by Name or NIP"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/teachers [get]
func TeachersGetHandler(c *gin.Context) {
	// Ambil parameter query dari request
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	specialization := c.DefaultQuery("specialization", "")
	sortByNIP, _ := strconv.ParseBool(c.DefaultQuery("sortByNIP", "false"))
	search := c.DefaultQuery("search", "")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	// Ambil data dengan pagination dan filter specialization
	students, total, err := teacherRepo.GetAllTeachers(context.Background(), page, pageSize, specialization, sortByNIP, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format respons dengan metadata pagination
	c.JSON(http.StatusOK, gin.H{
		"students": students,
		"meta": gin.H{
			"page":      page,
			"pageSize":  pageSize,
			"total":     total,
			"totalPage": int(math.Ceil(float64(total) / float64(pageSize))),
		},
	})
}

// TeachersPostHandler creates a new teacher
// @Summary Create Teacher
// @Description Create a new teacher in the database
// @Tags Teachers
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param teacher body models.CreateTeacherRequest true "Teacher data"
// @Success 200 {object} models.Teacher
// @Router /api/v1/teachers [post]
func TeachersPostHandler(c *gin.Context) {
	var req models.CreateTeacherRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateTeacher with the extracted values
	teacher, err := teacherRepo.CreateTeacher(context.Background(), req.Name, req.NIP, req.Specialization, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created teacher
	c.JSON(http.StatusOK, teacher)
}

// TeachersUpdateHandler updates an existing teacher
// @Summary Update Teacher
// @Description Updates an existing teacher in the database
// @Tags Teachers
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Teacher ID"
// @Param teacher body models.UpdateTeacherRequest true "Updated Teacher Data"
// @Success 200 {object} models.Teacher
// @Router /api/v1/teachers [patch]
func TeachersUpdateHandler(c *gin.Context) {
	var req models.UpdateTeacherRequest

	// Extract `id` from query parameters
	idStr := c.Query("id") // This will get the ID from query params
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing teacher ID"})
		return
	}

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle nil pointers for Phone_Number and Profile_Picture_URL
	phoneNumber := ""
	if req.Phone_Number != nil {
		phoneNumber = *req.Phone_Number // Dereference the pointer to get the string value
	}

	profilePictureURL := ""
	if req.Profile_Picture_URL != nil {
		profilePictureURL = *req.Profile_Picture_URL // Dereference the pointer to get the string value
	}

	// Call UpdateTeacher with the correct parameters
	teacher, err := teacherRepo.UpdateTeacher(
		context.Background(),
		id,
		req.Name,
		req.NIP,
		phoneNumber, // Pass the dereferenced phone number
		req.Specialization,
		req.Status,
		profilePictureURL, // Pass the dereferenced profile picture URL
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated teacher
	c.JSON(http.StatusOK, teacher)
}

// TeachersDeleteHandler deletes a teacher
// @Summary Delete Teacher
// @Description Deletes a teacher from the database by ID
// @Tags Teachers
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Teacher ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/teachers [delete]
func TeachersDeleteHandler(c *gin.Context) {
	// Extract `id` from query parameters
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing teacher ID"})
		return
	}

	// Call DeleteTeacher function from repository
	err = teacherRepo.DeleteTeacher(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}
