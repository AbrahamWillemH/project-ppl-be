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
// @Tags Exercises
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param exercise_id query int true "Exercise ID"
// @Param number query int true "Number"
// @Success 200 {array} models.ExerciseAnswers
// @Router /api/v1/exercises-answers [get]
func ExerciseAnswersGetHandler(c *gin.Context) {
	exerciseIDStr := c.Query("exercise_id")
	exerciseID, err := strconv.Atoi(exerciseIDStr)
	if err != nil || exerciseID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exercise_id"})
		return
	}

	exercises, err := exerciseAnswersRepo.GetExerciseAnswers(context.Background(), exerciseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercises)
}

// @Summary Create Exercise Answers
// @Description Create a new exercise
// @Tags Exercises
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

// // @Summary Update Exercise
// // @Description Update an existing exercise
// // @Tags Exercises
// // @Security BearerAuth
// // @Accept json
// // @Produce json
// // @Param id query int true "Exercise ID"
// // @Param exercise body models.CreateExercisesRequest true "Updated data"
// // @Success 200 {object} models.Exercises
// // @Router /api/v1/exercises [patch]
// func ExercisesUpdateHandler(c *gin.Context) {
// 	idStr := c.Query("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil || id <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exercise ID"})
// 		return
// 	}
//
// 	var req models.CreateExercisesRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	exercise, err := exercisesRepo.UpdateExercise(context.Background(), id, req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, exercise)
// }
//
// // @Summary Delete Exercise
// // @Description Delete an exercise by ID
// // @Tags Exercises
// // @Security BearerAuth
// // @Accept json
// // @Produce json
// // @Param id query int true "Exercise ID"
// // @Success 200 {object} map[string]string
// // @Router /api/v1/exercises [delete]
// func ExercisesDeleteHandler(c *gin.Context) {
// 	idStr := c.Query("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil || id <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing exercise ID"})
// 		return
// 	}
//
// 	err = exercisesRepo.DeleteExercise(context.Background(), id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, gin.H{"message": "Exercise deleted successfully"})
// }
