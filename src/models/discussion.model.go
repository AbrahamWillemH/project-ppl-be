package models

// Student represents the student model stored in the database
type Discussion struct {
	ID          int    `json:"id" db:"id"`
	Student_ID  string `json:"student_id" db:"student_id"`
	Topic       string `json:"topic" db:"topic"`
	Description string `json:"description" db:"description"`
	Replies     string `json:"replies" db:"replies"`
}

// CreateDiscussionRequest represents the request body for creating a class
type CreateDiscussionRequest struct {
	ID          int    `json:"id" db:"id"`
	Student_ID  string `json:"student_id" db:"student_id"`
	Topic       string `json:"topic" db:"topic"`
	Description string `json:"description" db:"description"`
	Replies     string `json:"replies" db:"replies"`
}
