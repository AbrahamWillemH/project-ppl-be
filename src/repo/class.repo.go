package repo

import (
	"context"
	"fmt"
	"project-ppl-be/config"
	"project-ppl-be/src/models"

	"github.com/huandu/go-sqlbuilder"
)

// ClassRepository struct
type ClassRepository struct{}

// GettAllClasses retrieves all classes from the database
func (r *ClassRepository) GettAllClasses(ctx context.Context, page, pageSize int) ([]models.Class, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "description", "teacher_id").
		From("classes").
		Limit(pageSize).
		Offset((page - 1) * pageSize) // OFFSET = (page - 1) * pageSize

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var classes []models.Class
	for rows.Next() {
		var class models.Class
		err := rows.Scan(&class.ID, &class.Name, &class.Description, &class.Teacher_ID)
		if err != nil {
			return nil, 0, err
		}
		classes = append(classes, class)
	}

	// Hitung total jumlah data untuk pagination
	countQuery := "SELECT COUNT(*) FROM classes"

	var total int

	err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return classes, total, nil
}

// CreateClass inserts a new class into the database
func (r *ClassRepository) CreateClass(ctx context.Context, name string, description string, teacher_id int) (models.Class, error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("classes").
		Cols("name", "description", "teacher_id").
		Values(name, description, teacher_id).
		Returning("id")

	// Convert to PostgreSQL-style placeholders ($1, $2, ...)
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	var classID int
	err := config.DB.QueryRow(ctx, query, args...).Scan(&classID)
	if err != nil {
		return models.Class{}, err
	}

	// Return the created class
	return models.Class{
		ID:          classID,
		Name:        name,
		Description: description,
		Teacher_ID:  teacher_id,
	}, nil
}

// UpdateClass update an existing class
func (r *ClassRepository) UpdateClass(ctx context.Context, id int, name string, description string, teacher_id int) (models.Class, error) {
	// Build the update query without Returning method
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("classes").
		Set(
			sb.Assign("name", name),
			sb.Assign("description", description),
			sb.Assign("teacher_id", teacher_id),
		).
		Where(sb.Equal("id", id))

	// Generate the query and arguments for PostgreSQL
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	// Append the RETURNING clause manually
	query += " RETURNING id, name, description, teacher_id"

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	// Prepare the result struct
	var class models.Class

	// Execute the query and scan the result into the class struct
	err := config.DB.QueryRow(ctx, query, args...).Scan(
		&class.ID,
		&class.Name,
		&class.Description,
		&class.Teacher_ID,
	)
	if err != nil {
		return models.Class{}, err
	}

	// Return the updated class
	return class, nil
}

// DeleteClass deletes a class from the database
func (r *ClassRepository) DeleteClass(ctx context.Context, id int) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("classes").Where(sb.Equal("id", id))
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := config.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
