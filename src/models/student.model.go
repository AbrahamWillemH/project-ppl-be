package models

// Student represents the student model stored in the database
type Student struct {
	ID                  int     `json:"id" db:"id"`
	Name                string  `json:"name" db:"name"`
	NIS                 string  `json:"nis" db:"nis"`
	Phone_Number        *string `json:"phone_number" db:"phone_number"`
	Grade               int     `json:"grade" db:"grade"`
	Current_Score       *int    `json:"curr_score" db:"curr_score"`
	Status              string  `json:"status" db:"status"`
	Profile_Picture_URL *string `json:"profile_picture_url" db:"profile_picture_url"`
	User_ID             string  `json:"user_id" db:"user_id"`
}

// CreateStudentRequest represents the request body for creating a student (without ID, Profile picture, user_id, phone number)
type CreateStudentRequest struct {
	Name   string `json:"name" db:"name"`
	NIS    string `json:"nis" db:"nis"`
	Grade  int    `json:"grade" db:"grade"`
	Status string `json:"status" db:"status"`
}

type UpdateStudentRequest struct {
	Name                string  `json:"name" db:"name"`
	NIS                 string  `json:"nis" db:"nis"`
	Phone_Number        *string `json:"phone_number" db:"phone_number"`
	Grade               int     `json:"grade" db:"grade"`
	Status              string  `json:"status" db:"status"`
	Current_Score       *int    `json:"curr_score" db:"curr_score"`
	Profile_Picture_URL *string `json:"profile_picture_url" db:"profile_picture_url"`
}
