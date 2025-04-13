package classes

import (
	"context"
	"math"
	"net/http"
	"project-ppl-be/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ClassGetByIdHandler retrieves a list of assigned students and teachers
// @Summary Get Class by Id
// @Description Fetch assigned students and teachers in the class
// @Tags Classes
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 15)"
// @Param id query int true "Class ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/classes/details [get]
func ClassGetByIdHandler(c *gin.Context) {
	idStr := c.Query("id") // This will get the ID from query params
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing class ID"})
		return
	}
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
	classes, total, err := classesRepo.GetClassById(context.Background(), id, page, pageSize)
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

// ClassAssignStudentsHandler assigns multiple students to a class
// @Summary Assign Students to Class
// @Description Assign multiple students to a class in the database
// @Tags Classes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param classAssignStudents body models.ClassAssignStudents true "Assign students to class"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Bad Request - Invalid Input"
// @Failure 500 {object} map[string]string "Internal Server Error - Database Issue"
// @Router /api/v1/classes/assign-students [post]
func ClassAssignStudentsHandler(c *gin.Context) {
	var req models.ClassAssignStudents

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call AssignStudents with the extracted values
	err := classesRepo.AssignStudents(context.Background(), req.ID, req.Student_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Students assigned successfully"})
}

// ClassUnassignStudentsHandler unassigns a student from a class
// @Summary Unassign Student from a Class
// @Description Unassigns a student from a class in the database
// @Tags Classes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param classAssignStudents body models.ClassAssignStudents true "Unassigns students from a class"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Bad Request - Invalid Input"
// @Failure 500 {object} map[string]string "Internal Server Error - Database Issue"
// @Router /api/v1/classes/unassign-students [delete]
func ClassUnassignStudentsHandler(c *gin.Context) {
	var req models.ClassAssignStudents

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call UnassignStudents with the extracted values
	err := classesRepo.UnassignStudents(context.Background(), req.ID, req.Student_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Students unassigned successfully"})
}
