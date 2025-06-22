package exams

import (
	"context"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var examsRepo = repo.ExamRepository{}

// @Summary Get Exams by Class ID
// @Description Fetch exam from a class
// @Tags Exams
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param class_id query int true "Class ID"
// @Success 200 {array} models.Exams
// @Router /api/v1/exams [get]
func ExamsGetByClassHandler(c *gin.Context) {
	classIDStr := c.Query("class_id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil || classID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing class_id"})
		return
	}

	exams, err := examsRepo.GetExamsByClassID(context.Background(), classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

// @Summary Get Exams for Student
// @Description Fetch all exams for a specific class
// @Tags Exams
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param class_id query int true "Class ID"
// @Param number query int true "Number"
// @Success 200 {array} models.Exams
// @Router /api/v1/exams/student [get]
func ExamsGetByClassForStudentHandler(c *gin.Context) {
	classIDStr := c.Query("class_id")
	numberStr := c.Query("number")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil || classID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing class_id"})
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil || number <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing number"})
		return
	}

	exams, err := examsRepo.GetExamsByClassIDForStudent(context.Background(), classID, number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

// @Summary Create Exam
// @Description Create a new exam
// @Tags Exams
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exam body models.CreateExamsRequest true "Exam data"
// @Success 200 {object} models.Exams
// @Router /api/v1/exams [post]
func ExamsPostHandler(c *gin.Context) {
	var req models.CreateExamsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := examsRepo.CreateExam(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

// @Summary Update Exam
// @Description Update an existing exam
// @Tags Exams
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Exam ID"
// @Param exam body models.CreateExamsRequest true "Updated data"
// @Success 200 {object} models.Exams
// @Router /api/v1/exams [patch]
func ExamsUpdateHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exam ID"})
		return
	}

	var req models.CreateExamsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := examsRepo.UpdateExam(context.Background(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

// @Summary Delete Exam
// @Description Delete an exam by ID
// @Tags Exams
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Exam ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/exams [delete]
func ExamsDeleteHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exam ID"})
		return
	}

	err = examsRepo.DeleteExam(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exam deleted successfully"})
}
