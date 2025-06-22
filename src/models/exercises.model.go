package models

// Exercises represents the exercises model stored in the database
type Exercises struct {
	ID          int    `json:"id" db:"id"`
	Material_ID  int    `json:"material_id" db:"material_id"`
	Title       string `json:"title" db:"title"`
	Content any `json:"content" db:"content"`
	Total_Marks     int    `json:"total_marks" db:"total_marks"`
	Teacher_ID int `json:"teacher_id" db:"teacher_id"`
}

// CreateExercisesRequest represents the request body for creating an exercise
type CreateExercisesRequest struct {
	Material_ID  int    `json:"material_id" db:"material_id"`
	Title       string `json:"title" db:"title"`
	Content any `json:"content" db:"content"`
	Total_Marks     int    `json:"total_marks" db:"total_marks"`
	Teacher_ID int `json:"teacher_id" db:"teacher_id"`
}

type ExerciseAnswers struct {
	ID          int    `json:"id" db:"id"`
	Exercise_ID  int    `json:"exercise_id" db:"exercise_id"`
	Answers any `json:"answers" db:"answers"`
	Student_ID int `json:"student_id" db:"student_id"`
	Status     string    `json:"status" db:"status"`
}

// CreateExerciseAnswersRequest represents the request body for creating an exercise
type CreateExerciseAnswersRequest struct {
	Exercise_ID  int    `json:"exercise_id" db:"exercise_id"`
	Answers any `json:"answers" db:"answers"`
	Student_ID int `json:"student_id" db:"student_id"`
}

type ExercisesDisplay struct {
	Title       string `json:"title" db:"title"`
	Content any `json:"content" db:"content"`
}

type ExerciseAnswersDisplay struct {
	Answers any `json:"answers" db:"answers"`
}

type ExerciseGrades struct {
	ID          int    `json:"id" db:"id"`
	Exercise_ID  int    `json:"exercise_id" db:"exercise_id"`
	Student_ID       int `json:"student_id" db:"student_id"`
	Score     float64    `json:"score" db:"score"`
	Detail any `json:"detail" db:"detail"`
}

type CalculateExerciseGrades struct {
	Exercise_ID  int    `json:"exercise_id" db:"exercise_id"`
	Student_ID       int `json:"student_id" db:"student_id"`
}
