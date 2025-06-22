package models

import "time"

// Exams represents the exercises model stored in the database
type Exams struct {
	ID          int    `json:"id" db:"id"`
	Class_ID  int    `json:"class_id" db:"class_id"`
	Title       string `json:"title" db:"title"`
	Content any `json:"content" db:"content"`
	Total_Marks     int    `json:"total_marks" db:"total_marks"`
	Teacher_ID int `json:"teacher_id" db:"teacher_id"`
	Start_Time   time.Time      `json:"start_time" db:"start_time"`
	End_Time     time.Time      `json:"end_time" db:"end_time"`
	Status       string         `json:"status" db:"status"`
}

// CreateExercisesRequest represents the request body for creating an exam
type CreateExamsRequest struct {
	Class_ID  int    `json:"class_id" db:"class_id"`
	Title       string `json:"title" db:"title"`
	Content any `json:"content" db:"content"`
	Total_Marks     int    `json:"total_marks" db:"total_marks"`
	Teacher_ID int `json:"teacher_id" db:"teacher_id"`
	Start_Time   time.Time      `json:"start_time" db:"start_time"`
	End_Time     time.Time      `json:"end_time" db:"end_time"`
}

type ExamAnswers struct {
	ID          int    `json:"id" db:"id"`
	Exam_ID  int    `json:"exam_id" db:"exam_id"`
	Answers any `json:"answers" db:"answers"`
	Student_ID int `json:"student_id" db:"student_id"`
	Status     string    `json:"status" db:"status"`
}

// CreateExamAnswersRequest represents the request body for creating an exam
type CreateExamAnswersRequest struct {
	Exam_ID  int    `json:"exam_id" db:"exam_id"`
	Answers any `json:"answers" db:"answers"`
	Student_ID int `json:"student_id" db:"student_id"`
}

type ExamsDisplay struct {
	Title       string `json:"title" db:"title"`
	Content any `json:"content" db:"content"`
}

type ExamAnswersDisplay struct {
	Answers any `json:"answers" db:"answers"`
}

type ExamGrades struct {
	ID          int    `json:"id" db:"id"`
	Exam_ID  int    `json:"exam_id" db:"exam_id"`
	Student_ID       int `json:"student_id" db:"student_id"`
	Score     float64    `json:"score" db:"score"`
	Detail any `json:"detail" db:"detail"`
}

type CalculateExamGrades struct {
	Exam_ID  int    `json:"exam_id" db:"exam_id"`
	Student_ID       int `json:"student_id" db:"student_id"`
}
