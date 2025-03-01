package users

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
// @Param pageSize query int false "Number of items per page (default: 40)"
// @Param specialization query string false "Filter by specialization (e.g., 'IPA')"
// @Param sortByNIP query bool false "Sort by NIP (true for ascending, false for descending)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/teachers [get]
func TeachersGetHandler(c *gin.Context) {
	// Ambil parameter query dari request
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "40"))
	specialization := c.DefaultQuery("specialization", "")
	sortByNIP, _ := strconv.ParseBool(c.DefaultQuery("sortByNIP", "false"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 40
	}

	// Ambil data dengan pagination dan filter specialization
	students, total, err := teacherRepo.GetAllStudents(context.Background(), page, pageSize, specialization, sortByNIP)
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
