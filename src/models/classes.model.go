package models

// Student represents the student model stored in the database
type Class struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Teacher_ID   int    `json:"teacher_id" db:"teacher_id"`
	Teacher_Name string `json:"teacher_name" db:"teacher_name"`
	Grade        int    `json:"grade" db:"grade"`
}

// CreateClassRequest represents the request body for creating a class
type CreateClassRequest struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Teacher_ID  int    `json:"teacher_id" db:"teacher_id"`
	Grade       int    `json:"grade" db:"grade"`
}

// ClassWithStudents represents class with students for get by id
type ClassWithStudents struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	Teacher_ID   int       `json:"teacher_id" db:"teacher_id"`
	Teacher_Name string    `json:"teacher_name" db:"teacher_name"`
	Students     []Student `json:"students"`
}

type ClassAssignStudents struct {
	ID         int   `json:"id" db:"id"`
	Student_ID []int `json:"student_id" db:"student_id"`
}
