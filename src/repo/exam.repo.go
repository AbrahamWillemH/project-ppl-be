package repo

import (
    "context"
    "encoding/json"
    "strconv"
    "strings"
		"fmt"

    "project-ppl-be/config"
    "project-ppl-be/src/models"

    "github.com/huandu/go-sqlbuilder"

		utils "project-ppl-be/src/utils"
)


type ExamRepository struct{}

// Get by class_id
func (r *ExamRepository) GetExamsByClassID(ctx context.Context, classID int) ([]models.Exams, error) {
    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("id", "class_id", "title", "content", "total_marks", "teacher_id", "start_time", "end_time", "status").
        From("exams").
        Where(sb.Equal("class_id", classID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.Exams
    for rows.Next() {
        var ex models.Exams
        if err := rows.Scan(&ex.ID, &ex.Class_ID, &ex.Title, &ex.Content, &ex.Total_Marks, &ex.Teacher_ID, &ex.Start_Time, &ex.End_Time, &ex.Status); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExamRepository) GetExamsByClassIDForStudent(ctx context.Context, classID int, number int) ([]models.Exams, error) {
    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("id", "class_id", "title", "content", "total_marks", "teacher_id", "start_time", "end_time", "status").
        From("exams").
        Where(sb.Equal("class_id", classID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    row := config.DB.QueryRow(ctx, query, args...)

    var ex models.Exams
    var contentBytes []byte

    if err := row.Scan(&ex.ID, &ex.Class_ID, &ex.Title, &contentBytes, &ex.Total_Marks, &ex.Teacher_ID, &ex.Start_Time, &ex.End_Time, &ex.Status); err != nil {
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

    return []models.Exams{ex}, nil
}

// Create
func (r *ExamRepository) CreateExam(ctx context.Context, req models.CreateExamsRequest) (models.Exams, error) {
		status := utils.CheckExamStatus(req.Start_Time, req.End_Time)

    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("exams").
        Cols("class_id", "title", "content", "total_marks", "teacher_id", "start_time", "end_time", "status").
        Values(req.Class_ID, req.Title, req.Content, req.Total_Marks, req.Teacher_ID, req.Start_Time, req.End_Time, status).
        Returning("id", "class_id", "title", "content", "total_marks", "teacher_id", "start_time", "end_time", "status")

    query, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)
    var ex models.Exams
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Class_ID, &ex.Title, &ex.Content, &ex.Total_Marks, &ex.Teacher_ID, &ex.Start_Time, &ex.End_Time, &ex.Status,
    )
    if err != nil {
        return models.Exams{}, err
    }
    return ex, nil
}

// Update
func (r *ExamRepository) UpdateExam(ctx context.Context, id int, req models.CreateExamsRequest) (models.Exams, error) {
		status := utils.CheckExamStatus(req.Start_Time, req.End_Time)

    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("exams").
        Set(
            ub.Assign("class_id", req.Class_ID),
            ub.Assign("title", req.Title),
            ub.Assign("content", req.Content),
            ub.Assign("total_marks", req.Total_Marks),
            ub.Assign("teacher_id", req.Teacher_ID),
            ub.Assign("start_time", req.Start_Time),
            ub.Assign("end_time", req.End_Time),
            ub.Assign("status", status),
        ).
        Where(ub.Equal("id", id))

    query, args := ub.BuildWithFlavor(sqlbuilder.PostgreSQL)
    query += " RETURNING id, class_id, title, content, total_marks, teacher_id, start_time, end_time, status"

    var ex models.Exams
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Class_ID, &ex.Title, &ex.Content, &ex.Total_Marks, &ex.Teacher_ID, &ex.Start_Time, &ex.End_Time, &ex.Status,
    )
    if err != nil {
        return models.Exams{}, err
    }
    return ex, nil
}

// Delete
func (r *ExamRepository) DeleteExam(ctx context.Context, id int) error {
    db := sqlbuilder.NewDeleteBuilder()
    db.DeleteFrom("exams").Where(db.Equal("id", id))
    query, args := db.BuildWithFlavor(sqlbuilder.PostgreSQL)
    _, err := config.DB.Exec(ctx, query, args...)
    return err
}

func (r *ExamRepository) GetExamAnswers(ctx context.Context, examID int, studentID int) ([]models.ExamAnswers, error) {
    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("id", "exam_id", "answers", "student_id").
        From("exam_answers").
        Where(sb.Equal("exam_id", examID)).
				Where(sb.Equal("student_id", studentID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.ExamAnswers
    for rows.Next() {
        var ex models.ExamAnswers
        if err := rows.Scan(&ex.ID, &ex.Exam_ID, &ex.Answers, &ex.Student_ID); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExamRepository) CreateExamAnswers(ctx context.Context, req models.CreateExamAnswersRequest) (models.ExamAnswers, error) {
    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("exam_answers").
        Cols("exam_id", "answers", "student_id").
        Values(req.Exam_ID, req.Answers, req.Student_ID, "Active").
        Returning("id", "exam_id", "answers", "student_id")

    query, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)
    var ex models.ExamAnswers
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Exam_ID, &ex.Answers, &ex.Student_ID,
    )
    if err != nil {
        return models.ExamAnswers{}, err
    }
    return ex, nil
}

// Update
func (r *ExamRepository) UpdateExamAnswers(ctx context.Context, id int, req models.CreateExamAnswersRequest) (models.ExamAnswers, error) {
    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("exam_answers").
        Set(
            ub.Assign("exam_id", req.Exam_ID),
            ub.Assign("answers", req.Answers),
            ub.Assign("student_id", req.Student_ID),
        ).
        Where(ub.Equal("id", id))

    query, args := ub.BuildWithFlavor(sqlbuilder.PostgreSQL)
    query += " RETURNING id, exam_id, answers, student_id"

    var ex models.ExamAnswers
    err := config.DB.QueryRow(ctx, query, args...).Scan(
        &ex.ID, &ex.Exam_ID, &ex.Answers, &ex.Student_ID,
    )
    if err != nil {
        return models.ExamAnswers{}, err
    }
    return ex, nil
}

// Delete
func (r *ExamRepository) DeleteExamAnswers(ctx context.Context, id int) error {
    db := sqlbuilder.NewDeleteBuilder()
    db.DeleteFrom("exam_answers").Where(db.Equal("id", id))
    query, args := db.BuildWithFlavor(sqlbuilder.PostgreSQL)
    _, err := config.DB.Exec(ctx, query, args...)
    return err
}

func (r *ExamRepository) GetExamGrades(ctx context.Context, examID int, studentID int) ([]models.ExamGrades, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "student_id", "exam_id", "score").
			From("exam_scores").
			Where(sb.Equal("exam_id", examID)).
			Where(sb.Equal("student_id", studentID)).
			OrderBy("id DESC").
			Limit(1)

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.ExamGrades
    for rows.Next() {
        var ex models.ExamGrades
        if err := rows.Scan(&ex.ID, &ex.Student_ID, &ex.Exam_ID, &ex.Score); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExamRepository) GetAllExamGrades(ctx context.Context, studentID int) ([]models.ExamGrades, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "student_id", "exam_id", "score").
			From("exam_scores").
			Where(sb.Equal("student_id", studentID))

    query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
    rows, err := config.DB.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.ExamGrades
    for rows.Next() {
        var ex models.ExamGrades
        if err := rows.Scan(&ex.ID, &ex.Student_ID, &ex.Exam_ID, &ex.Score); err != nil {
            return nil, err
        }
        list = append(list, ex)
    }
    return list, rows.Err()
}

func (r *ExamRepository) CalculateExamGrades(ctx context.Context, req models.CalculateExamGrades) (models.ExamGrades, error) {
    // ---------------------------------------------------
    // 1️⃣ Ambil jawaban dari exam_answers
    sbAns := sqlbuilder.NewSelectBuilder()
    sbAns.Select("answers").
        From("exam_answers").
        Where(sbAns.Equal("exam_id", req.Exam_ID)).
        Where(sbAns.Equal("student_id", req.Student_ID))
    queryAns, argsAns := sbAns.BuildWithFlavor(sqlbuilder.PostgreSQL)

    var answerBytes []byte
    if err := config.DB.QueryRow(ctx, queryAns, argsAns...).Scan(&answerBytes); err != nil {
        return models.ExamGrades{}, fmt.Errorf("failed to get student answers: %w", err)
    }

    var studentAnswers map[string]string
    if err := json.Unmarshal(answerBytes, &studentAnswers); err != nil {
        return models.ExamGrades{}, fmt.Errorf("failed to unmarshal student answers: %w", err)
    }

    // ---------------------------------------------------
    // 2️⃣ Ambil soal dan total_marks dari exams
    sbEx := sqlbuilder.NewSelectBuilder()
    sbEx.Select("content", "total_marks").
        From("exams").
        Where(sbEx.Equal("id", req.Exam_ID))
    queryEx, argsEx := sbEx.BuildWithFlavor(sqlbuilder.PostgreSQL)

    var contentBytes []byte
    var totalMarks int
    if err := config.DB.QueryRow(ctx, queryEx, argsEx...).Scan(&contentBytes, &totalMarks); err != nil {
        return models.ExamGrades{}, fmt.Errorf("failed to get exam data: %w", err)
    }

    var fullContent map[string]string
    if err := json.Unmarshal(contentBytes, &fullContent); err != nil {
        return models.ExamGrades{}, fmt.Errorf("failed to unmarshal exam content: %w", err)
    }

    // ---------------------------------------------------
    // 3️⃣ Hitung jumlah soal dan nilai per soal
    correctAnswerKeys := make(map[string]string)
    for key, val := range fullContent {
        if strings.HasSuffix(key, "_answer") {
            soalNumber := strings.TrimSuffix(key, "_answer") // contoh: "1_answer" -> "1"
            correctAnswerKeys[soalNumber] = val
        }
    }

    totalQuestions := len(correctAnswerKeys)
    if totalQuestions == 0 {
        return models.ExamGrades{}, fmt.Errorf("no questions found for exam id %d", req.Exam_ID)
    }

    marksPerQuestion := float64(totalMarks) / float64(totalQuestions)

    // ---------------------------------------------------
    // 4️⃣ Hitung total nilai
    var totalScore float64
    for soalNumber, correctAnswer := range correctAnswerKeys {
        if studentAnswers[soalNumber] == correctAnswer {
            totalScore += marksPerQuestion
        }
    }

    // ---------------------------------------------------
    // 5️⃣ Simpan nilai ke exam_scores
    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("exam_scores").
        Cols("student_id", "exam_id", "score").
        Values(req.Student_ID, req.Exam_ID, totalScore).
        Returning("student_id", "exam_id", "score")

    queryInsert, argsInsert := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)

    var savedScore models.ExamGrades
    if err := config.DB.QueryRow(ctx, queryInsert, argsInsert...).Scan(
        &savedScore.Student_ID,
        &savedScore.Exam_ID,
        &savedScore.Score,
    ); err != nil {
        return models.ExamGrades{}, fmt.Errorf("failed to insert exam score: %w", err)
    }

    // ---------------------------------------------------
    // ✅ Return nilai yang berhasil disimpan
    return savedScore, nil
}
