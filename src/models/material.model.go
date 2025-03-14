package models

// Student represents the student model stored in the database
type Material struct {
	ID          int    `json:"id" db:"id"`
	Class_ID    int    `json:"class_id" db:"class_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Content     int    `json:"content" db:"content"`
	Teacher_ID  int    `json:"teacher_id" db:"teacher_id"`
}

// CreateMaterialRequest represents the request body for creating a student (without ID, Profile picture, user_id, phone number)
type CreateMaterialRequest struct {
	Class_ID    int    `json:"class_id" db:"class_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Content     int    `json:"content" db:"content"`
	Teacher_ID  int    `json:"teacher_id" db:"teacher_id"`
}

type UpdateMaterialRequest struct {
	Class_ID    int    `json:"class_id" db:"class_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Content     int    `json:"content" db:"content"`
	Teacher_ID  int    `json:"teacher_id" db:"teacher_id"`
}
