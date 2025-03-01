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

var studentRepo = repo.StudentRepository{}

// StudentsGetHandler retrieves a list of students
// @Summary Get Students
// @Description Fetch all students from the database with pagination, filtering by grade, and sorting by NIS
// @Tags Students
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 40)"
// @Param grade query string false "Filter by grade (e.g., '10')"
// @Param sortByNIS query bool false "Sort by NIS (true for ascending, false for descending)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/students [get]
func StudentsGetHandler(c *gin.Context) {
	// Ambil parameter query dari request
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "40"))
	grade := c.DefaultQuery("grade", "")
	sortByNIS, _ := strconv.ParseBool(c.DefaultQuery("sortByNIS", "false"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 40
	}

	// Ambil data dengan pagination dan filter grade
	students, total, err := studentRepo.GetAllStudents(context.Background(), page, pageSize, grade, sortByNIS)
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

// StudentPostHandler creates a new student
// @Summary Create Student
// @Description Create a new student in the database
// @Tags Students
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.CreateStudentRequest true "Student data"
// @Success 200 {object} models.Student
// @Router /api/v1/students [post]
func StudentPostHandler(c *gin.Context) {
	var req models.CreateStudentRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateUser with the extracted values
	user, err := studentRepo.CreateStudent(context.Background(), req.Name, req.NIS, req.Grade, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, user)
}
