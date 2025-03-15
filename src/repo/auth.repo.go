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
	"github.com/huandu/go-sqlbuilder"
	"golang.org/x/crypto/bcrypt"
)

// AuthRepository struct
type AuthRepository struct{}

// LoginUser finds a user by email and verifies password
func (r *AuthRepository) LoginUser(ctx context.Context, username, password string) (string, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "username", "email", "password", "role", "display_name").From("users")
	sb.Where(sb.Equal("username", username))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	fmt.Println("Executing SQL:", query, "with args:", args) // Debugging log

	row := config.DB.QueryRow(ctx, query, args...)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Display_Name)
	if err != nil {
		fmt.Println("User not found error:", err) // Debugging log
		return "", errors.New("user not found")
	}

	// Pastikan password di database sudah dalam bentuk bcrypt hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Password mismatch:", err) // Debugging log
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := generateJWT(user)
	if err != nil {
		fmt.Println("Error generating token:", err) // Debugging log
		return "", err
	}

	return token, nil
}

// GenerateJWT creates a JWT token for authentication
func generateJWT(user models.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "defaultsecret" // Gunakan env var untuk keamanan
	}

	claims := jwt.MapClaims{
		"user_id":      user.ID,
		"email":        user.Email,
		"role":         user.Role,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // Expire dalam 24 jam
		"display_name": user.Display_Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
