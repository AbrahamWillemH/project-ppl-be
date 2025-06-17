package classes

import (
	"context"
	"math"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var classesRepo = repo.ClassRepository{}

// ClassGetHandler retrieves a list of classes
// @Summary Get Class
// @Description Fetch all classes from the database with pagination
// @Tags Classes
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/classes [get]
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
	classes, total, err := classesRepo.GettAllClasses(context.Background(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format respons dengan metadata pagination
	c.JSON(http.StatusOK, gin.H{
		"classes": classes,
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
// @Router /api/v1/classes [post]
func ClassPostHandler(c *gin.Context) {
	var req models.CreateClassRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateUser with the extracted values
	user, err := classesRepo.CreateClass(context.Background(), req.Name, req.Description, req.Teacher_ID, req.Grade)
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
// @Router /api/v1/classes [patch]
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
	class, err := classesRepo.UpdateClass(
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
// @Router /api/v1/classes [delete]
func ClassDeleteHandler(c *gin.Context) {
	// Extract `id` from query parameters
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing teacher ID"})
		return
	}

	// Call DeleteTeacher function from repository
	err = classesRepo.DeleteClass(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

// GetClassIDHandler retrieves class id based on grade and logged in teacher
// @Summary Get class id by grade and logged in teacher
// @Description Fetch class id
// @Tags Classes
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param grade query int true "Grade"
// @Param teacher_id query int true "teacher_id"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/classes/class-id [get]
func GetClassIDHandler(c *gin.Context) {
	// Ambil `grade` dari query
	gradeStr := c.Query("grade")
	grade, err := strconv.Atoi(gradeStr)
	if err != nil || grade <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing grade"})
		return
	}

	// Ambil `teacher_id` dari query
	teacherStr := c.Query("teacher_id")
	teacherID, err := strconv.Atoi(teacherStr)
	if err != nil || teacherID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing teacher_id"})
		return
	}

	// Ambil data dari repo
	classes, total, err := classesRepo.GetClassId(context.Background(), grade, teacherID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika tidak ada kelas ditemukan
	if len(classes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No class found"})
		return
	}

	// Ambil hanya ID saja dari kelas pertama
	classID := classes[0].ID

	c.JSON(http.StatusOK, gin.H{
		"class_id": classID,
		"total":    total,
	})
}

// GetClassForStudentHandler retrieves assigned classes for a student by ID
// @Summary Get classes for a student
// @Description Fetch assigned classes for a student using their ID with pagination
// @Tags Classes
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param id query int true "Student ID"
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/classes/assigned [get]
func GetClassForStudentHandler(c *gin.Context) {
	// Ambil student ID dari query
	idStr := c.Query("id")
	studentID, err := strconv.Atoi(idStr)
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student ID"})
		return
	}

	// Ambil parameter pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}

	// Ambil data dari repository
	classes, total, err := classesRepo.GetClassesByStudentID(context.Background(), studentID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika tidak ada kelas ditemukan
	if len(classes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No classes found for the student"})
		return
	}

	// Kirim response
	c.JSON(http.StatusOK, gin.H{
		"classes": classes,
		"meta": gin.H{
			"page":      page,
			"pageSize":  pageSize,
			"total":     total,
			"totalPage": int(math.Ceil(float64(total) / float64(pageSize))),
		},
	})
}
