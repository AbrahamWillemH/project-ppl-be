package repo

import (
	"context"
	"fmt"
	"project-ppl-be/config"
	"project-ppl-be/models"

	"github.com/huandu/go-sqlbuilder"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository struct
type UserRepository struct{}

// GetAllUsers retrieves all users from the database
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "username", "email", "password", "role").From("users")

	query, args := sb.Build()
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(ctx context.Context, username, email, password, role string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("users").
		Cols("username", "email", "password", "role").
		Values(username, email, hashedPassword, role).
		Returning("id")

	// Convert to PostgreSQL-style placeholders ($1, $2, $3, ...)
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	var userID int
	err = config.DB.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:       userID,
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}, nil
}
