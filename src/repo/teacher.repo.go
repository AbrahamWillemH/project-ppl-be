package repo

import (
	"context"
	"fmt"
	"project-ppl-be/config"
	"project-ppl-be/src/models"

	"github.com/huandu/go-sqlbuilder"
)

// TeacherRepository struct
type TeacherRepository struct{}

// GetAllTeachers retrieves all teachers from the database
func (r *TeacherRepository) GetAllTeachers(ctx context.Context, page, pageSize int, specialization string, sortByNIP bool) ([]models.Teacher, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "nip", "phone_number", "specialization", "status", "profile_picture_url", "user_id").
		From("teachers").
		Limit(pageSize).
		Offset((page - 1) * pageSize) // OFFSET = (page - 1) * pageSize

	// Add specialization taken from params
	if specialization != "" {
		sb.Where(sb.Equal("specialization", specialization))
	}

	// Sorting berdasarkan NIP jika diaktifkan
	if sortByNIP {
		sb.OrderBy("nip ASC")
	} else {
		sb.OrderBy("nip DESC")
	}

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.NIP, &teacher.Phone_Number, &teacher.Specialization, &teacher.Status, &teacher.Profile_Picture_URL, &teacher.User_ID)
		if err != nil {
			return nil, 0, err
		}
		teachers = append(teachers, teacher)
	}

	// Hitung total jumlah data untuk pagination
	countQuery := "SELECT COUNT(*) FROM teachers"
	if specialization != "" {
		countQuery += " WHERE specialization = $1"
	}
	var total int
	if specialization != "" {
		err = config.DB.QueryRow(ctx, countQuery, specialization).Scan(&total)
	} else {
		err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	}
	if err != nil {
		return nil, 0, err
	}

	return teachers, total, nil
}

// CreateTeacher inserts a new teacher into the database
func (r *TeacherRepository) CreateTeacher(ctx context.Context, name, nip, specialization, status string) (models.Teacher, error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("teachers").
		Cols("name", "nip", "specialization", "status").
		Values(name, nip, specialization, status).
		Returning("id")

	// Convert to PostgreSQL-style placeholders ($1, $2, ...)
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	var studentID int
	err := config.DB.QueryRow(ctx, query, args...).Scan(&studentID)
	if err != nil {
		return models.Teacher{}, err
	}

	// Return the created teacher
	return models.Teacher{
		ID:             studentID,
		Name:           name,
		NIP:            nip,
		Specialization: specialization,
		Status:         status,
	}, nil
}

// UpdateTeacher update an existing teacher
func (r *TeacherRepository) UpdateTeacher(
	ctx context.Context,
	id int,
	name, nip, phone_number, specialization, status, profile_picture_url string,
) (models.Teacher, error) {
	// Build the update query without Returning method
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("teachers").
		Set(
			sb.Assign("name", name),
			sb.Assign("nip", nip),
			sb.Assign("phone_number", phone_number),
			sb.Assign("specialization", specialization),
			sb.Assign("status", status),
			sb.Assign("profile_picture_url", profile_picture_url),
		).
		Where(sb.Equal("id", id))

	// Generate the query and arguments for PostgreSQL
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	// Append the RETURNING clause manually
	query += " RETURNING id, name, nip, phone_number, specialization, status, profile_picture_url"

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	// Prepare the result struct
	var teacher models.Teacher

	// Execute the query and scan the result into the teacher struct
	err := config.DB.QueryRow(ctx, query, args...).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.NIP,
		&teacher.Phone_Number,
		&teacher.Specialization,
		&teacher.Status,
		&teacher.Profile_Picture_URL,
	)
	if err != nil {
		return models.Teacher{}, err
	}

	// Return the updated teacher
	return teacher, nil
}

// DeleteTeacher deletes a teacher from the database
func (r *TeacherRepository) DeleteTeacher(ctx context.Context, id int) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("teachers").Where(sb.Equal("id", id))
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := config.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
