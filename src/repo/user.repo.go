package repo

import (
	"context"
	"fmt"
	"project-ppl-be/config"
	"project-ppl-be/src/models"

	"github.com/huandu/go-sqlbuilder"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository struct
type UserRepository struct{}

// GetAllUsers retrieves all users from the database
func (r *UserRepository) GetAllUsers(ctx context.Context, page, pageSize int, role string, sortByUsername bool) ([]models.User, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "username", "email", "role").
		From("users").
		Limit(pageSize).
		Offset((page - 1) * pageSize) // OFFSET = (page - 1) * pageSize

	// Tambahkan filter berdasarkan role jika ada
	if role != "" {
		sb.Where(sb.Equal("role", role))
	}

	// Sorting berdasarkan username jika diaktifkan
	if sortByUsername {
		sb.OrderBy("username ASC")
	} else {
		sb.OrderBy("username DESC")
	}

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	// Hitung total jumlah data untuk pagination
	countQuery := "SELECT COUNT(*) FROM users"
	if role != "" {
		countQuery += " WHERE role = $1"
	}
	var total int
	if role != "" {
		err = config.DB.QueryRow(ctx, countQuery, role).Scan(&total)
	} else {
		err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	}
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
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
