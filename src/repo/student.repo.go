package repo

import (
	"context"
	"fmt"
	"project-ppl-be/config"
	"project-ppl-be/src/models"

	"github.com/huandu/go-sqlbuilder"
)

// StudentRepository struct
type StudentRepository struct{}

// GetAllStudents retrieves all students from the database
func (r *StudentRepository) GetAllStudents(ctx context.Context, page, pageSize int, grade string, sortByNIS bool) ([]models.Student, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "nis", "phone_number", "grade", "curr_score", "status", "profile_picture_url", "user_id").
		From("students").
		Limit(pageSize).
		Offset((page - 1) * pageSize) // OFFSET = (page - 1) * pageSize

	// Tambahkan filter berdasarkan grade jika ada
	if grade != "" {
		sb.Where(sb.Equal("grade", grade))
	}

	// Sorting berdasarkan NIS jika diaktifkan
	if sortByNIS {
		sb.OrderBy("nis ASC")
	} else {
		sb.OrderBy("nis DESC")
	}

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.Name, &student.NIS, &student.Phone_Number, &student.Grade, &student.Current_Score, &student.Status, &student.Profile_Picture_URL, &student.User_ID)
		if err != nil {
			return nil, 0, err
		}
		students = append(students, student)
	}

	// Hitung total jumlah data untuk pagination
	countQuery := "SELECT COUNT(*) FROM students"
	if grade != "" {
		countQuery += " WHERE grade = $1"
	}
	var total int
	if grade != "" {
		err = config.DB.QueryRow(ctx, countQuery, grade).Scan(&total)
	} else {
		err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	}
	if err != nil {
		return nil, 0, err
	}

	return students, total, nil
}

// CreateStudent inserts a new student into the database
func (r *StudentRepository) CreateStudent(ctx context.Context, name, nis, grade, status string) (models.Student, error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("students").
		Cols("name", "nis", "grade", "status").
		Values(name, nis, grade, status).
		Returning("id")

	// Convert to PostgreSQL-style placeholders ($1, $2, ...)
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	var studentID int
	err := config.DB.QueryRow(ctx, query, args...).Scan(&studentID)
	if err != nil {
		return models.Student{}, err
	}

	// Return the created student
	return models.Student{
		ID:     studentID,
		Name:   name,
		NIS:    nis,
		Grade:  grade,
		Status: status,
	}, nil
}
