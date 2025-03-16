package repo

import (
	"context"
	"fmt"
	"project-ppl-be/config"
	"project-ppl-be/src/models"

	"github.com/huandu/go-sqlbuilder"
)

// StudentRepository struct
type MaterialRepository struct{}

// GetAllStudents retrieves all materials from the database
func (r *MaterialRepository) GetAllMaterials(ctx context.Context, page, pageSize int) ([]models.Material, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "class_id", "title", "description", "teacher_id").
		From("materials").
		Limit(pageSize).
		Offset((page - 1) * pageSize) // OFFSET = (page - 1) * pageSize

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var materials []models.Material
	for rows.Next() {
		var material models.Material
		err := rows.Scan(&material.ID, &material.Class_ID, &material.Title, &material.Description, &material.Content, &material.Teacher_ID)
		if err != nil {
			return nil, 0, err
		}
		materials = append(materials, material)
	}

	// Hitung total jumlah data untuk pagination
	countQuery := "SELECT COUNT(*) FROM materials"

	var total int

	err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return materials, total, nil
}

// CreateStudent inserts a new material into the database
func (r *MaterialRepository) CreateMaterial(ctx context.Context, class_id int, title string, description string, content string, teacher_id int) (models.Material, error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("materials").
		Cols("class_id", "title", "description", "content", "teacher_id").
		Values(class_id, title, description, content, teacher_id).
		Returning("id")

	// Convert to PostgreSQL-style placeholders ($1, $2, ...)
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	var materialID int
	err := config.DB.QueryRow(ctx, query, args...).Scan(&materialID)
	if err != nil {
		return models.Material{}, err
	}

	// Return the created material
	return models.Material{
		ID:          materialID,
		Class_ID:    class_id,
		Title:       title,
		Description: description,
		Content:     content,
		Teacher_ID:  teacher_id,
	}, nil
}

// UpdateStudent update an existing material
func (r *MaterialRepository) UpdateMaterial(ctx context.Context, id int, class_id int, title string, description string, content string, teacher_id int) (models.Material, error) {
	// Build the update query without Returning method
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("materials").
		Set(
			sb.Assign("class_id", class_id),
			sb.Assign("title", title),
			sb.Assign("description", description),
			sb.Assign("content", content),
			sb.Assign("teacher_id", teacher_id),
		).
		Where(sb.Equal("id", id))

	// Generate the query and arguments for PostgreSQL
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	// Append the RETURNING clause manually
	query += " RETURNING id, class_id, title, description, content, teacher_id"

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	// Prepare the result struct
	var material models.Material

	// Execute the query and scan the result into the material struct
	err := config.DB.QueryRow(ctx, query, args...).Scan(
		&material.ID,
		&material.Class_ID,
		&material.Title,
		&material.Description,
		&material.Content,
		&material.Teacher_ID,
	)
	if err != nil {
		return models.Material{}, err
	}

	// Return the updated material
	return material, nil
}

// DeleteTeacher deletes a material from the database
func (r *MaterialRepository) DeleteMaterial(ctx context.Context, id int) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("materials").Where(sb.Equal("id", id))
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := config.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
