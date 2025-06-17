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
	sb.Select("classes.id", "classes.name", "classes.description", "classes.grade", "classes.teacher_id", "teachers.name AS teacher_name").
		From("classes").
		Join("teachers", "classes.teacher_id = teachers.id").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var classes []models.Class
	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID, &class.Name, &class.Description, &class.Grade, &class.Teacher_ID, &class.Teacher_Name); err != nil {
			return nil, 0, err
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	countQuery := "SELECT COUNT(*) FROM classes"
	var total int
	err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return classes, total, nil
}

// CreateClass inserts a new class into the database
func (r *ClassRepository) CreateClass(ctx context.Context, name string, description string, teacher_id int, grade int) (models.Class, error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("classes").
		Cols("name", "description", "teacher_id", "grade").
		Values(name, description, teacher_id, grade).
		Returning("id")

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	var classID int
	err := config.DB.QueryRow(ctx, query, args...).Scan(&classID)
	if err != nil {
		return models.Class{}, err
	}

	return models.Class{
		ID:          classID,
		Name:        name,
		Description: description,
		Teacher_ID:  teacher_id,
		Grade:       grade,
	}, nil
}

// UpdateClass updates an existing class
func (r *ClassRepository) UpdateClass(ctx context.Context, id int, name string, description string, teacher_id int, grade int) (models.Class, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("classes").
		Set(
			sb.Assign("name", name),
			sb.Assign("description", description),
			sb.Assign("teacher_id", teacher_id),
			sb.Assign("grade", grade),
		).
		Where(sb.Equal("id", id))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	query += " RETURNING id, name, description, teacher_id, grade"

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	var class models.Class
	err := config.DB.QueryRow(ctx, query, args...).Scan(
		&class.ID,
		&class.Name,
		&class.Description,
		&class.Teacher_ID,
		&class.Grade,
	)
	if err != nil {
		return models.Class{}, err
	}

	return class, nil
}

// DeleteClass deletes a class from the database
func (r *ClassRepository) DeleteClass(ctx context.Context, id int) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("classes").Where(sb.Equal("id", id))
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := config.DB.Exec(ctx, query, args...)
	return err
}

// GetClassById retrieves all classes from the database
func (r *ClassRepository) GetClassById(ctx context.Context, classID, page, pageSize int) (*models.ClassWithStudents, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"c.id", "c.name", "c.description",
		"c.teacher_id", "t.name AS teacher_name",
		"s.id AS student_id", "s.name AS student_name",
		"s.nis AS student_nis", "s.status AS student_status",
		"s.grade AS student_grade",
	).
		From("classes c").
		Join("teachers t", "c.teacher_id = t.id").
		JoinWithOption(sqlbuilder.LeftJoin, "assigned_students_class asc_tbl", "c.id = asc_tbl.class_id").
		JoinWithOption(sqlbuilder.LeftJoin, "students s", "asc_tbl.student_id = s.id").
		Where(sb.Equal("c.id", classID)).
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var class models.ClassWithStudents
	class.Students = []models.Student{}

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(
			&class.ID, &class.Name, &class.Description,
			&class.Teacher_ID, &class.Teacher_Name,
			&student.ID, &student.Name, &student.NIS, &student.Status, &student.Grade,
		); err != nil {
			return nil, 0, err
		}
		if student.ID != 0 { // Pastikan student ada sebelum dimasukkan
			class.Students = append(class.Students, student)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// Hitung total siswa dalam kelas
	countQuery := "SELECT COUNT(*) FROM assigned_students_class WHERE class_id = $1"
	var total int
	err = config.DB.QueryRow(ctx, countQuery, classID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return &class, total, nil
}

// CreateClass inserts a new class into the database
func (r *ClassRepository) AssignStudents(ctx context.Context, classID int, studentIDs []int) error {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("assigned_students_class").
		Cols("student_id", "class_id")

	// Add multiple student_id values in a batch insert
	for _, studentID := range studentIDs {
		sb.Values(studentID, classID)
	}

	// Generate SQL query
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	// Execute the insert query
	_, err := config.DB.Exec(ctx, query, args...)
	return err
}

// UnassignStudents removes multiple students from a class
func (r *ClassRepository) UnassignStudents(ctx context.Context, classID int, studentIDs []int) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("assigned_students_class").
		Where(sb.Equal("class_id", classID))

	// Tambahkan kondisi untuk menghapus hanya student_id yang ada di studentIDs
	if len(studentIDs) > 0 {
		// Konversi []int ke []interface{}
		ids := make([]interface{}, len(studentIDs))
		for i, v := range studentIDs {
			ids[i] = v
		}
		sb.Where(sb.In("student_id", ids...))
	}

	// Generate SQL query
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Args:", args)

	// Execute the delete query
	_, err := config.DB.Exec(ctx, query, args...)
	return err
}

// GettAllClasses retrieves all classes from the database
func (r *ClassRepository) GetClassId(ctx context.Context, grade, teacher_id int) ([]models.Class, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("classes.id").
		From("classes").
		Join("teachers", "classes.teacher_id = teachers.id").
		Where(sb.Equal("grade", grade), sb.Equal("teacher_id", teacher_id))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var classes []models.Class
	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID); err != nil {
			return nil, 0, err
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	countQuery := "SELECT COUNT(*) FROM classes"
	var total int
	err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return classes, total, nil
}

// GetClassesByStudentID fetches all classes assigned to a student with pagination
func (r *ClassRepository) GetClassesByStudentID(ctx context.Context, studentID, page, pageSize int) ([]models.Class, int, error) {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select(
		"c.id", "c.name", "c.description",
		"c.teacher_id", "t.name AS teacher_name",
	).
		From("assigned_students_class asc_tbl").
		Join("classes c", "asc_tbl.class_id = c.id").
		Join("teachers t", "c.teacher_id = t.id").
		Where(sb.Equal("asc_tbl.student_id", studentID)).
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var classes []models.Class
	for rows.Next() {
		var class models.Class
		if err := rows.Scan(
			&class.ID, &class.Name, &class.Description,
			&class.Teacher_ID, &class.Teacher_Name,
		); err != nil {
			return nil, 0, err
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// Hitung total semua kelas untuk student ini
	countSb := sqlbuilder.NewSelectBuilder()
	countSb.Select("COUNT(*)").
		From("assigned_students_class").
		Where(countSb.Equal("student_id", studentID))

	countQuery, countArgs := countSb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	var total int
	err = config.DB.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return classes, total, nil
}
