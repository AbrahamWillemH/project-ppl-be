package exams

import (
	"context"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var examAnswersRepo = repo.ExamRepository{}

// @Summary Get Exam Answers by Exam ID
// @Description Fetch all exams for a specific material
// @Tags Exam Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exam_id query int true "Exam ID"
// @Param student_id query int true "Student ID"
// @Success 200 {array} models.ExamAnswers
// @Router /api/v1/exams-answers [get]
func ExamAnswersGetHandler(c *gin.Context) {
	examIDStr := c.Query("exam_id")
	examID, err := strconv.Atoi(examIDStr)
	if err != nil || examID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exam_id"})
		return
	}

	studentIDStr := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student_id"})
		return
	}

	exams, err := examAnswersRepo.GetExamAnswers(context.Background(), examID, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

// @Summary Create Exam Answers
// @Description Create a new exam
// @Tags Exam Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exam body models.CreateExamAnswersRequest true "Exam answers data"
// @Success 200 {object} models.ExamAnswers
// @Router /api/v1/exams-answers [post]
func ExamAnswersPostHandler(c *gin.Context) {
	var req models.CreateExamAnswersRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := examsRepo.CreateExamAnswers(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

// @Summary Update Exam Answer
// @Description Update an existing exam answer
// @Tags Exam Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Exam Answer ID (Student's Answer ID)"
// @Param exam body models.CreateExamAnswersRequest true "Updated data"
// @Success 200 {object} models.ExamAnswers
// @Router /api/v1/exams-answers [patch]
func ExamAnswersUpdateHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exam ID"})
		return
	}

	var req models.CreateExamAnswersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := examsRepo.UpdateExamAnswers(context.Background(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

// @Summary Delete Exam Answer
// @Description Delete an exam answer by ID
// @Tags Exam Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Exam ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/exams-answers [delete]
func ExamAnswersDeleteHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exam ID"})
		return
	}

	err = examsRepo.DeleteExamAnswers(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exam deleted successfully"})
}

// @Summary Calculate Grade
// @Description Calculate grade after submitting exam
// @Tags Exam Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exam body models.CalculateExamGrades true "Exam grades data"
// @Success 200 {object} models.ExamGrades
// @Router /api/v1/exams/calculate-grade [post]
func CalculateGradePostHandler(c *gin.Context) {
	var req models.CalculateExamGrades

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := examsRepo.CalculateExamGrades(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

// @Summary Get Grade
// @Description Get calculated grade
// @Tags Exam Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exam_id query int true "Exam ID"
// @Param student_id query int true "Student ID"
// @Success 200 {object} models.ExamGrades
// @Router /api/v1/exams/get-grade [get]
func ExamGradesGetHandler(c *gin.Context) {
	examIDStr := c.Query("exam_id")
	examID, err := strconv.Atoi(examIDStr)
	if err != nil || examID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exam_id"})
		return
	}

	studentIDStr := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student_id"})
		return
	}

	exams, err := examAnswersRepo.GetExamGrades(context.Background(), examID, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

// @Summary Get ALl Grades
// @Description Get calculated grades for all exams
// @Tags Exam Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param student_id query int true "Student ID"
// @Success 200 {object} models.ExamGrades
// @Router /api/v1/exams/get-all-grade [get]
func ExamsAllGradesGetHandler(c *gin.Context) {
	studentIDStr := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student_id"})
		return
	}

	exams, err := examAnswersRepo.GetAllExamGrades(context.Background(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}
