package repo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"project-ppl-be/config"
	"project-ppl-be/src/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AuthRepository struct
type AuthRepository struct{}

// LoginUser finds a user by username and verifies password
func (r *AuthRepository) LoginUser(ctx context.Context, username, password string) (string, error) {
	// Query untuk menggabungkan tabel users dengan teachers dan students
	query := `
		SELECT u.id, u.username, u.email, u.password, u.role, u.display_name,
		       t.id AS teacher_id, s.id AS student_id
		FROM users u
		LEFT JOIN teachers t ON t.user_id = u.id
		LEFT JOIN students s ON s.user_id = u.id
		WHERE u.username = $1
	`

	// Eksekusi query
	row := config.DB.QueryRow(ctx, query, username)

	var user models.User
	var teacherID, studentID *int

	// Scan hasil query ke variabel user dan ID untuk teacher dan student
	err := row.Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Display_Name,
		&teacherID, &studentID,
	)
	if err != nil {
		fmt.Println("User not found error:", err)
		return "", errors.New("user not found")
	}

	// Cek kecocokan password dengan bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Password mismatch:", err)
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := generateJWT(user, teacherID, studentID)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}

	return token, nil
}

// GenerateJWT creates a JWT token for authentication
func generateJWT(user models.User, teacherID, studentID *int) (string, error) {
	// Ambil secret key dari environment variable
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "defaultsecret"
	}

	// Buat klaim untuk JWT
	claims := jwt.MapClaims{
		"user_id":      user.ID,
		"email":        user.Email,
		"role":         user.Role,
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
		"display_name": user.Display_Name,
	}

	// Jika role teacher, tambahkan teacher_id ke klaim
	if user.Role == "teacher" && teacherID != nil {
		claims["teacher_id"] = *teacherID
	}

	// Jika role student, tambahkan student_id ke klaim
	if user.Role == "student" && studentID != nil {
		claims["student_id"] = *studentID
	}

	// Buat token dengan klaim yang sudah dibuat
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
