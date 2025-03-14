package models

// Teacher represents the user model stored in the database
type Teacher struct {
	ID                  int     `json:"id" db:"id"`
	Name                string  `json:"name" db:"name"`
	NIP                 string  `json:"nip" db:"nip"`
	Phone_Number        *string `json:"phone_number" db:"phone_number"`
	Specialization      string  `json:"specialization" db:"specialization"`
	Status              string  `json:"status" db:"status"`
	Profile_Picture_URL *string `json:"profile_picture_url" db:"profile_picture_url"`
	User_ID             string  `json:"user_id" db:"user_id"`
}

// CreateTeacherRequest represents the request body for creating a teacher (without ID)
type CreateTeacherRequest struct {
	Name           string `json:"name" db:"name"`
	NIP            string `json:"nip" db:"nip"`
	Specialization string `json:"specialization" db:"specialization"`
	Status         string `json:"status" db:"status"`
}

// UpdateTeacherRequest represents the request body for updating a teacher
type UpdateTeacherRequest struct {
	Name                string  `json:"name" db:"name"`
	NIP                 string  `json:"nip" db:"nip"`
	Phone_Number        *string `json:"phone_number" db:"phone_number"`
	Specialization      string  `json:"specialization" db:"specialization"`
	Status              string  `json:"status" db:"status"`
	Profile_Picture_URL *string `json:"profile_picture_url" db:"profile_picture_url"`
}
