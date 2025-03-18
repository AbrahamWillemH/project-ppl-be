package models

// Student represents the student model stored in the database
type Class struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Teacher_ID   int    `json:"teacher_id" db:"teacher_id"`
	Teacher_Name string `json:"teacher_name" db:"teacher_name"`
}

// CreateClassRequest represents the request body for creating a class
type CreateClassRequest struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Teacher_ID  int    `json:"teacher_id" db:"teacher_id"`
}
