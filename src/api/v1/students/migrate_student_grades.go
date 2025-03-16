package students

import (
	"context"
	"net/http"
	"project-ppl-be/src/models"

	"github.com/gin-gonic/gin"
)

// StudentPostHandler migrates all student grades by 1 level
// @Summary Migrates Student Grades by 1.
// @Description Accepts "up" to increase grade, and "down" to decrease grade
// @Tags Students
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param student body models.MigrateStudentGradeRequest true "Student grade"
// @Success 200 {object} models.Student
// @Router /api/v1/students/grade-migrate [post]
func StudentGradeMigrateHandler(c *gin.Context) {
	var req models.MigrateStudentGradeRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default to "up" if not provided
	if req.Migrate == "" {
		req.Migrate = "up"
	}

	// Call MigrateStudentGrade to update all students' grades
	err := studentRepo.MigrateStudentGrade(context.Background(), req.Migrate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "All student grades migrated successfully"})
}
