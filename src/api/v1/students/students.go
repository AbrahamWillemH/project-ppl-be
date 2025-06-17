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

var studentRepo = repo.StudentRepository{}

// StudentsGetHandler retrieves a list of students
// @Summary Get Students
// @Description Fetch all students from the database with pagination, filtering by grade, and sorting by NIS
// @Tags Students
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Param grade query string false "Filter by grade (e.g., '10')"
// @Param sortByNIS query bool false "Sort by NIS (true for ascending, false for descending)"
// @Param search query string false "Search by Name or NIS"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/students [get]
func StudentsGetHandler(c *gin.Context) {
	// Ambil parameter query dari request
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	grade := c.DefaultQuery("grade", "")
	sortByNIS, _ := strconv.ParseBool(c.DefaultQuery("sortByNIS", "false"))
	search := c.DefaultQuery("search", "")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	// Ambil data dengan pagination dan filter grade
	students, total, err := studentRepo.GetAllStudents(context.Background(), page, pageSize, grade, sortByNIS, search)
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

// StudentUpdateHandler updates an existing student
// @Summary Update Student
// @Description Updates an existing student in the database
// @Tags Students
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Student ID"
// @Param student body models.UpdateStudentRequest true "Updated Student Data"
// @Success 200 {object} models.Student
// @Router /api/v1/students [patch]
func StudentUpdateHandler(c *gin.Context) {
	var req models.UpdateStudentRequest

	// Extract `id` from query parameters
	idStr := c.Query("id") // This will get the ID from query params
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student ID"})
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

	// Call UpdateStudent with the correct parameters
	student, err := studentRepo.UpdateStudent(
		context.Background(),
		id,
		req.Name,
		req.NIS,
		phoneNumber, // Pass the dereferenced phone number
		req.Grade,
		req.Status,
		req.Current_Score,
		profilePictureURL, // Pass the dereferenced profile picture URL
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated student
	c.JSON(http.StatusOK, student)
}

// StudentDeleteHandler deletes a student
// @Summary Delete Student
// @Description Deletes a student from the database by ID
// @Tags Students
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Student ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/students [delete]
func StudentDeleteHandler(c *gin.Context) {
	// Extract `id` from query parameters
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student ID"})
		return
	}

	// Call DeleteStudent function from repository
	err = studentRepo.DeleteStudent(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

// StudentGetByIDHandler retrieves a student by their ID
// @Summary Get Student by ID
// @Description Fetch a single student from the database by their ID
// @Tags Students
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id query int true "Student ID"
// @Success 200 {object} models.Student
// @Router /api/v1/students/details [get]
func StudentGetByIDHandler(c *gin.Context) {
	// Ambil parameter ID dari path
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student ID"})
		return
	}

	// Ambil data student berdasarkan ID
	student, err := studentRepo.GetStudentByID(context.Background(), id)
	if err != nil {
		// Misal repositori return error kalau student tidak ditemukan
		if err.Error() == "student not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respon dengan data student
	c.JSON(http.StatusOK, student)
}
