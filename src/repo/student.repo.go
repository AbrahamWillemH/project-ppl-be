package repo

import (
	"context"
	"fmt"
	"project-ppl-be/config"
	"project-ppl-be/src/models"
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

// StudentRepository struct
type StudentRepository struct{}

// GetAllStudents retrieves all students from the database
func (r *StudentRepository) GetAllStudents(ctx context.Context, page, pageSize int, grade string, sortByNIS bool, search string) ([]models.Student, int, error) {
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

	if search != "" {
		lowerSearch := strings.ToLower(search) // Ubah input menjadi lowercase
		sb.Where(sb.Or(
			sb.Like("LOWER(name)", "%"+lowerSearch+"%"),
			sb.Like("LOWER(nis)", "%"+lowerSearch+"%"),
		))
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
func (r *StudentRepository) CreateStudent(ctx context.Context, name string, nis string, grade int, status string) (models.Student, error) {
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

// UpdateStudent update an existing student
func (r *StudentRepository) UpdateStudent(
	ctx context.Context,
	id int,
	name string, nis string, phone_number string, grade int, status string, curr_score *int, profile_picture_url string,
) (models.Student, error) {
	// Build the update query without Returning method
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("students").
		Set(
			sb.Assign("name", name),
			sb.Assign("nis", nis),
			sb.Assign("phone_number", phone_number),
			sb.Assign("grade", grade),
			sb.Assign("status", status),
			sb.Assign("curr_score", curr_score),
			sb.Assign("profile_picture_url", profile_picture_url),
		).
		Where(sb.Equal("id", id))

	// Generate the query and arguments for PostgreSQL
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	// Append the RETURNING clause manually
	query += " RETURNING id, name, nis, phone_number, grade, status, curr_score, profile_picture_url"

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	// Prepare the result struct
	var student models.Student

	// Execute the query and scan the result into the student struct
	err := config.DB.QueryRow(ctx, query, args...).Scan(
		&student.ID,
		&student.Name,
		&student.NIS,
		&student.Phone_Number,
		&student.Grade,
		&student.Status,
		&student.Current_Score,
		&student.Profile_Picture_URL,
	)
	if err != nil {
		return models.Student{}, err
	}

	// Return the updated student
	return student, nil
}

// DeleteSTudent deletes a student from the database
func (r *StudentRepository) DeleteStudent(ctx context.Context, id int) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("students").Where(sb.Equal("id", id))
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := config.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// MigrateStudentGrade inserts a new student into the database
func (r *StudentRepository) MigrateStudentGrade(ctx context.Context, migrate string) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("students")

	switch migrate {
	case "up":
		ub.Set("grade = grade + 1")
	case "down":
		ub.Set("grade = grade - 1")
	default:
		return fmt.Errorf("invalid migration type: %s", migrate)
	}

	query, args := ub.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := config.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
