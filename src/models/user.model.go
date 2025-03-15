package models

// User represents the user model stored in the database
type User struct {
	ID           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	Password     string `json:"password" db:"password"`
	Role         string `json:"role" db:"role"`
	Display_Name string `json:"display_name" db:"display_name"`
}

// CreateUserRequest represents the request body for creating a user (without ID)
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
