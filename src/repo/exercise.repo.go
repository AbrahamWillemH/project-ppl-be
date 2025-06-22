package repo

import (
    "context"
    "encoding/json"
    "strconv"
    "strings"
		"fmt"
		utils "project-ppl-be/src/utils"

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

func (r *ExerciseRepository) GetExercisesByMaterialIDForStudent(ctx context.Context, materialID int, number int) ([]models.Exercises, error) {
    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("id", "material_id", "title", "content", "total_marks", "teacher_id").
        From("exercises").
        Where(sb.Equal("material_id", materialID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    row := config.DB.QueryRow(ctx, query, args...)

    var ex models.Exercises
    var contentBytes []byte

    if err := row.Scan(&ex.ID, &ex.Material_ID, &ex.Title, &contentBytes, &ex.Total_Marks, &ex.Teacher_ID); err != nil {
        return nil, err
    }

    // Unmarshal content (JSON) ke map
    var fullContent map[string]string
    if err := json.Unmarshal(contentBytes, &fullContent); err != nil {
        return nil, err
    }

    // Konversi number ke string
    numberStr := strconv.Itoa(number)

    // Filter content yang sesuai nomor soal dan bukan _answer
    filtered := make(map[string]string)
    for key, val := range fullContent {
        if (key == numberStr || strings.HasPrefix(key, numberStr+"_")) && !strings.HasSuffix(key, "_answer") {
            filtered[key] = val
        }
    }

    // Assign content hasil filter ke ex.Content
    ex.Content = filtered

    return []models.Exercises{ex}, nil
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

func (r *ExerciseRepository) GetExerciseAnswers(ctx context.Context, exerciseID int, studentID int) ([]models.ExerciseAnswers, error) {
    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("id", "exercise_id", "answers", "student_id", "status").
        From("exercise_answers").
        Where(sb.Equal("exercise_id", exerciseID)).
				Where(sb.Equal("student_id", studentID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.ExerciseAnswers
    for rows.Next() {
        var ex models.ExerciseAnswers
        if err := rows.Scan(&ex.ID, &ex.Exercise_ID, &ex.Answers, &ex.Student_ID, &ex.Status); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExerciseRepository) CreateExerciseAnswers(ctx context.Context, req models.CreateExerciseAnswersRequest) (models.ExerciseAnswers, error) {
    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("exercise_answers").
        Cols("exercise_id", "answers", "student_id", "status").
        Values(req.Exercise_ID, req.Answers, req.Student_ID, "Active").
        Returning("id", "exercise_id", "answers", "student_id", "status")

    query, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)
    var ex models.ExerciseAnswers
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Exercise_ID, &ex.Answers, &ex.Student_ID, &ex.Status,
    )
    if err != nil {
        return models.ExerciseAnswers{}, err
    }
    return ex, nil
}

// Update
func (r *ExerciseRepository) UpdateExerciseAnswers(ctx context.Context, id int, req models.CreateExerciseAnswersRequest) (models.ExerciseAnswers, error) {
    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("exercise_answers").
        Set(
            ub.Assign("exercise_id", req.Exercise_ID),
            ub.Assign("answers", req.Answers),
            ub.Assign("student_id", req.Student_ID),
        ).
        Where(ub.Equal("id", id))

    query, args := ub.BuildWithFlavor(sqlbuilder.PostgreSQL)
    query += " RETURNING id, exercise_id, answers, student_id, status"

    var ex models.ExerciseAnswers
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Exercise_ID, &ex.Answers, &ex.Student_ID, &ex.Status,
    )
    if err != nil {
        return models.ExerciseAnswers{}, err
    }
    return ex, nil
}

// Delete
func (r *ExerciseRepository) DeleteExerciseAnswers(ctx context.Context, id int) error {
    db := sqlbuilder.NewDeleteBuilder()
    db.DeleteFrom("exercise_answers").Where(db.Equal("id", id))
    query, args := db.BuildWithFlavor(sqlbuilder.PostgreSQL)
    _, err := config.DB.Exec(ctx, query, args...)
    return err
}

func (r *ExerciseRepository) GetExerciseGrades(ctx context.Context, exerciseID int, studentID int) ([]models.ExerciseGrades, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "student_id", "exercise_id", "score", "detail").
			From("exercise_scores").
			Where(sb.Equal("exercise_id", exerciseID)).
			Where(sb.Equal("student_id", studentID)).
			OrderBy("id DESC").
			Limit(1)

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.ExerciseGrades
    for rows.Next() {
        var ex models.ExerciseGrades
        if err := rows.Scan(&ex.ID, &ex.Student_ID, &ex.Exercise_ID, &ex.Score, &ex.Detail); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExerciseRepository) GetAllExerciseGrades(ctx context.Context, studentID int) ([]models.ExerciseGrades, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "student_id", "exercise_id", "score", "detail").
			From("exercise_scores").
			Where(sb.Equal("student_id", studentID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.ExerciseGrades
    for rows.Next() {
        var ex models.ExerciseGrades
        if err := rows.Scan(&ex.ID, &ex.Student_ID, &ex.Exercise_ID, &ex.Score, &ex.Detail); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExerciseRepository) CalculateExerciseGrades(ctx context.Context, req models.CalculateExerciseGrades) (models.ExerciseGrades, error) {
    // ---------------------------------------------------
    // 1️⃣ Get student answers
    sbAns := sqlbuilder.NewSelectBuilder()
    sbAns.Select("answers").
        From("exercise_answers").
        Where(sbAns.Equal("exercise_id", req.Exercise_ID)).
        Where(sbAns.Equal("student_id", req.Student_ID))
    queryAns, argsAns := sbAns.BuildWithFlavor(sqlbuilder.PostgreSQL)

    var answerBytes []byte
    if err := config.DB.QueryRow(ctx, queryAns, argsAns...).Scan(&answerBytes); err != nil {
        return models.ExerciseGrades{}, fmt.Errorf("failed to get student answers: %w", err)
    }

    var studentAnswers map[string]string
    if err := json.Unmarshal(answerBytes, &studentAnswers); err != nil {
        return models.ExerciseGrades{}, fmt.Errorf("failed to unmarshal student answers: %w", err)
    }

    // ---------------------------------------------------
    // 2️⃣ Get exercise content and total marks
    sbEx := sqlbuilder.NewSelectBuilder()
    sbEx.Select("content", "total_marks").
        From("exercises").
        Where(sbEx.Equal("id", req.Exercise_ID))
    queryEx, argsEx := sbEx.BuildWithFlavor(sqlbuilder.PostgreSQL)

    var contentBytes []byte
    var totalMarks int
    if err := config.DB.QueryRow(ctx, queryEx, argsEx...).Scan(&contentBytes, &totalMarks); err != nil {
        return models.ExerciseGrades{}, fmt.Errorf("failed to get exercise data: %w", err)
    }

    var fullContent map[string]string
    if err := json.Unmarshal(contentBytes, &fullContent); err != nil {
        return models.ExerciseGrades{}, fmt.Errorf("failed to unmarshal exercise content: %w", err)
    }

    // ---------------------------------------------------
    // 3️⃣ Identify questions
    correctAnswerKeys := make(map[string]string)
    essayQuestions := make([]string, 0)

    for key, val := range fullContent {
        if strings.HasSuffix(key, "_answer") {
            soalNumber := strings.TrimSuffix(key, "_answer")
            correctAnswerKeys[soalNumber] = val
        } else if strings.HasSuffix(key, "_essay") {
            soalNumber := strings.TrimSuffix(key, "_essay")
            essayQuestions = append(essayQuestions, soalNumber)
        }
    }

    totalQuestions := len(correctAnswerKeys) + len(essayQuestions)
    if totalQuestions == 0 {
        return models.ExerciseGrades{}, fmt.Errorf("no questions found for exercise id %d", req.Exercise_ID)
    }

    marksPerQuestion := float64(totalMarks) / float64(totalQuestions)

    // ---------------------------------------------------
    // 4️⃣ Calculate total score and build detail
    var totalScore float64
    detail := make(map[string]interface{})

    // MCQs
    for soalNumber, correctAnswer := range correctAnswerKeys {
        if studentAnswers[soalNumber] == correctAnswer {
            totalScore += marksPerQuestion
            detail[soalNumber] = true
        } else {
            detail[soalNumber] = false
        }
    }

    // Essays
    for _, soalNumber := range essayQuestions {
        essayKey := soalNumber + "_essay"
        questionText := fullContent[essayKey]
        essayScore, err := utils.EvaluateEssayWithGemini(questionText, studentAnswers[soalNumber], marksPerQuestion)
        if err != nil {
            return models.ExerciseGrades{}, fmt.Errorf("failed to evaluate essay: %w", err)
        }
        totalScore += essayScore
        detail[soalNumber] = fmt.Sprintf("%.2f/%.2f", essayScore, marksPerQuestion)
    }

    // ---------------------------------------------------
    // 5️⃣ Save score and detail
    detailBytes, err := json.Marshal(detail)
    if err != nil {
        return models.ExerciseGrades{}, fmt.Errorf("failed to marshal detail: %w", err)
    }

    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("exercise_scores").
        Cols("student_id", "exercise_id", "score", "detail").
        Values(req.Student_ID, req.Exercise_ID, totalScore, detailBytes).
        Returning("id", "student_id", "exercise_id", "score", "detail")

    queryInsert, argsInsert := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)

    var savedScore models.ExerciseGrades
    if err := config.DB.QueryRow(ctx, queryInsert, argsInsert...).Scan(
        &savedScore.ID,
        &savedScore.Student_ID,
        &savedScore.Exercise_ID,
        &savedScore.Score,
        &savedScore.Detail,
    ); err != nil {
        return models.ExerciseGrades{}, fmt.Errorf("failed to insert exercise score: %w", err)
    }

    // ---------------------------------------------------
    // 6️⃣ Mark the answer status as Inactive
    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("exercise_answers").
        Set(ub.Assign("status", "Inactive")).
        Where(ub.Equal("exercise_id", req.Exercise_ID)).
        Where(ub.Equal("student_id", req.Student_ID))

    queryUpdate, argsUpdate := ub.BuildWithFlavor(sqlbuilder.PostgreSQL)

    if _, err := config.DB.Exec(ctx, queryUpdate, argsUpdate...); err != nil {
        return models.ExerciseGrades{}, fmt.Errorf("failed to update exercise_answer status: %w", err)
    }

    // ---------------------------------------------------
    // ✅ Done
    return savedScore, nil
}
