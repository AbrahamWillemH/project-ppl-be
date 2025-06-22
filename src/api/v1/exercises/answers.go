package exercises

import (
	"context"
	"net/http"
	"project-ppl-be/src/models"
	"project-ppl-be/src/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

var exerciseAnswersRepo = repo.ExerciseRepository{}

// @Summary Get Exercise Answers by Exercise ID
// @Description Fetch all exercises for a specific material
// @Tags Exercise Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exercise_id query int true "Exercise ID"
// @Param student_id query int true "Student ID"
// @Success 200 {array} models.ExerciseAnswers
// @Router /api/v1/exercises-answers [get]
func ExerciseAnswersGetHandler(c *gin.Context) {
	exerciseIDStr := c.Query("exercise_id")
	exerciseID, err := strconv.Atoi(exerciseIDStr)
	if err != nil || exerciseID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exercise_id"})
		return
	}

	studentIDStr := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student_id"})
		return
	}

	exercises, err := exerciseAnswersRepo.GetExerciseAnswers(context.Background(), exerciseID, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercises)
}

// @Summary Create Exercise Answers
// @Description Create a new exercise
// @Tags Exercise Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exercise body models.CreateExerciseAnswersRequest true "Exercise answers data"
// @Success 200 {object} models.ExerciseAnswers
// @Router /api/v1/exercises-answers [post]
func ExerciseAnswersPostHandler(c *gin.Context) {
	var req models.CreateExerciseAnswersRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := exercisesRepo.CreateExerciseAnswers(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

// @Summary Update Exercise Answer
// @Description Update an existing exercise answer
// @Tags Exercise Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Exercise Answer ID (Student's Answer ID)"
// @Param exercise body models.CreateExerciseAnswersRequest true "Updated data"
// @Success 200 {object} models.ExerciseAnswers
// @Router /api/v1/exercises-answers [patch]
func ExerciseAnswersUpdateHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exercise ID"})
		return
	}

	var req models.CreateExerciseAnswersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := exercisesRepo.UpdateExerciseAnswers(context.Background(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

// @Summary Delete Exercise Answer
// @Description Delete an exercise answer by ID
// @Tags Exercise Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id query int true "Exercise ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/exercises-answers [delete]
func ExerciseAnswersDeleteHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exercise ID"})
		return
	}

	err = exercisesRepo.DeleteExerciseAnswers(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exercise deleted successfully"})
}

// @Summary Calculate Grade
// @Description Calculate grade after submitting exercise
// @Tags Exercise Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exercise body models.CalculateExerciseGrades true "Exercise grades data"
// @Success 200 {object} models.ExerciseGrades
// @Router /api/v1/exercises/calculate-grade [post]
func CalculateGradePostHandler(c *gin.Context) {
	var req models.CalculateExerciseGrades

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := exercisesRepo.CalculateExerciseGrades(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

// @Summary Get Grade
// @Description Get calculated grade
// @Tags Exercise Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exercise_id query int true "Exercise ID"
// @Param student_id query int true "Student ID"
// @Success 200 {object} models.ExerciseGrades
// @Router /api/v1/exercises/get-grade [get]
func ExerciseGradesGetHandler(c *gin.Context) {
	exerciseIDStr := c.Query("exercise_id")
	exerciseID, err := strconv.Atoi(exerciseIDStr)
	if err != nil || exerciseID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exercise_id"})
		return
	}

	studentIDStr := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student_id"})
		return
	}

	exercises, err := exerciseAnswersRepo.GetExerciseGrades(context.Background(), exerciseID, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercises)
}

// @Summary Get All Grades
// @Description Get calculated grades for all materials
// @Tags Exercise Answers (Student Answers)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param student_id query int true "Student ID"
// @Success 200 {object} models.ExerciseGrades
// @Router /api/v1/exercises/get-all-grade [get]
func ExerciseAllGradesGetHandler(c *gin.Context) {
	studentIDStr := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing student_id"})
		return
	}

	exercises, err := exerciseAnswersRepo.GetAllExerciseGrades(context.Background(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercises)
}
