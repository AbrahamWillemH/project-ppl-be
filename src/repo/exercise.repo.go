package repo

import (
    "context"
    "project-ppl-be/config"
    "project-ppl-be/src/models"

    "github.com/huandu/go-sqlbuilder"
)

type ExerciseRepository struct{}

// Get by class_id
func (r *ExerciseRepository) GetExercisesByMaterialID(ctx context.Context, materialID int) ([]models.Exercises, error) {
    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("id", "material_id", "title", "content", "total_marks", "teacher_id").
        From("exercises").
        Where(sb.Equal("material_id", materialID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.Exercises
    for rows.Next() {
        var ex models.Exercises
        if err := rows.Scan(&ex.ID, &ex.Material_ID, &ex.Title, &ex.Content, &ex.Total_Marks, &ex.Teacher_ID); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

// Create
func (r *ExerciseRepository) CreateExercise(ctx context.Context, req models.CreateExercisesRequest) (models.Exercises, error) {
    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("exercises").
        Cols("material_id", "title", "content", "total_marks", "teacher_id").
        Values(req.Material_ID, req.Title, req.Content, req.Total_Marks, req.Teacher_ID).
        Returning("id", "material_id", "title", "content", "total_marks", "teacher_id")

    query, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)
    var ex models.Exercises
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Material_ID, &ex.Title, &ex.Content, &ex.Total_Marks, &ex.Teacher_ID,
    )
    if err != nil {
        return models.Exercises{}, err
    }
    return ex, nil
}

// Update
func (r *ExerciseRepository) UpdateExercise(ctx context.Context, id int, req models.CreateExercisesRequest) (models.Exercises, error) {
    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("exercises").
        Set(
            ub.Assign("material_id", req.Material_ID),
            ub.Assign("title", req.Title),
            ub.Assign("content", req.Content),
            ub.Assign("total_marks", req.Total_Marks),
            ub.Assign("teacher_id", req.Teacher_ID),
        ).
        Where(ub.Equal("id", id))

    query, args := ub.BuildWithFlavor(sqlbuilder.PostgreSQL)
    query += " RETURNING id, material_id, title, content, total_marks, teacher_id"

    var ex models.Exercises
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Material_ID, &ex.Title, &ex.Content, &ex.Total_Marks, &ex.Teacher_ID,
    )
    if err != nil {
        return models.Exercises{}, err
    }
    return ex, nil
}

// Delete
func (r *ExerciseRepository) DeleteExercise(ctx context.Context, id int) error {
    db := sqlbuilder.NewDeleteBuilder()
    db.DeleteFrom("exercises").Where(db.Equal("id", id))
    query, args := db.BuildWithFlavor(sqlbuilder.PostgreSQL)
    _, err := config.DB.Exec(ctx, query, args...)
    return err
}

func (r *ExerciseRepository) GetExerciseAnswers(ctx context.Context, exerciseID int) ([]models.ExerciseAnswers, error) {
    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("id", "exercise_id", "answers", "teacher_id").
        From("exercise_answers").
        Where(sb.Equal("exercise_id", exerciseID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.ExerciseAnswers
    for rows.Next() {
        var ex models.ExerciseAnswers
        if err := rows.Scan(&ex.ID, &ex.Exercise_ID, &ex.Answers, &ex.Teacher_ID); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExerciseRepository) CreateExerciseAnswers(ctx context.Context, req models.CreateExerciseAnswersRequest) (models.ExerciseAnswers, error) {
    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("exercise_answers").
        Cols("exercise_id", "answers", "teacher_id").
        Values(req.Exercise_ID, req.Answers, req.Teacher_ID).
        Returning("id", "exercise_id", "answers", "teacher_id")

    query, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)
    var ex models.ExerciseAnswers
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Exercise_ID, &ex.Answers, &ex.Teacher_ID,
    )
    if err != nil {
        return models.ExerciseAnswers{}, err
    }
    return ex, nil
}
