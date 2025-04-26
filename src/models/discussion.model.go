package models

// Student represents the student model stored in the database
type Discussion struct {
	ID          int    `json:"id" db:"id"`
	Student_ID  int    `json:"student_id" db:"student_id"`
	Topic       string `json:"topic" db:"topic"`
	Description string `json:"description" db:"description"`
	Replies     any    `json:"replies" db:"replies"`
}

// CreateDiscussionRequest represents the request body for creating a class
type CreateDiscussionRequest struct {
	Student_ID  int    `json:"student_id" db:"student_id"`
	Topic       string `json:"topic" db:"topic"`
	Description string `json:"description" db:"description"`
	Replies     any    `json:"replies" db:"replies"`
}

type ReplyDiscussion struct {
	Student_ID   int    `json:"student_id"`
	Student_Name string `json:"student_name"`
	Replies      string `json:"replies"`
}
