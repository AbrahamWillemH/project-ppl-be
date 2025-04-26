package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"project-ppl-be/config"
	"project-ppl-be/src/models"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

// DiscussionRepository struct
type DiscussionRepository struct{}

// GetAllDiscussions retrieves all discussions from the database
func (r *DiscussionRepository) GetAllDiscussions(ctx context.Context, page, pageSize int) ([]models.Discussion, int, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "student_id", "topic", "description", "replies").
		From("general_forum").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := config.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var discussions []models.Discussion
	for rows.Next() {
		var discussion models.Discussion
		if err := rows.Scan(&discussion.ID, &discussion.Student_ID, &discussion.Topic, &discussion.Description, &discussion.Replies); err != nil {
			return nil, 0, err
		}
		discussions = append(discussions, discussion)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	countQuery := "SELECT COUNT(*) FROM general_forum"
	var total int
	err = config.DB.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return discussions, total, nil
}

// CreateDiscussion inserts a new discussion into the database
func (r *DiscussionRepository) CreateDiscussion(ctx context.Context, studentID int, topic, description string, replies any) (models.Discussion, error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("general_forum").
		Cols("student_id", "topic", "description", "replies").
		Values(studentID, topic, description, replies).
		Returning("id")

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	var discussionID int
	err := config.DB.QueryRow(ctx, query, args...).Scan(&discussionID)
	if err != nil {
		return models.Discussion{}, err
	}

	return models.Discussion{
		ID:          discussionID,
		Student_ID:  studentID,
		Topic:       topic,
		Description: description,
		Replies:     replies,
	}, nil
}

// UpdateDiscussion updates an existing discussion
func (r *DiscussionRepository) UpdateDiscussion(ctx context.Context, id int, studentID int, topic, description string, replies any) (models.Discussion, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("general_forum").
		Set(
			sb.Assign("student_id", studentID),
			sb.Assign("topic", topic),
			sb.Assign("description", description),
			sb.Assign("replies", replies),
		).
		Where(sb.Equal("id", id))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	query += " RETURNING id, student_id, topic, description, replies"

	var discussion models.Discussion
	err := config.DB.QueryRow(ctx, query, args...).Scan(
		&discussion.ID,
		&discussion.Student_ID,
		&discussion.Topic,
		&discussion.Description,
		&discussion.Replies,
	)
	if err != nil {
		return models.Discussion{}, err
	}

	return discussion, nil
}

// DeleteDiscussion deletes a discussion from the database
func (r *DiscussionRepository) DeleteDiscussion(ctx context.Context, id int) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("general_forum").Where(sb.Equal("id", id))
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := config.DB.Exec(ctx, query, args...)
	return err
}

// GetDiscussionById retrieves a discussion by ID
func (r *DiscussionRepository) GetDiscussionById(ctx context.Context, id int) (*models.Discussion, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "student_id", "topic", "description", "replies").
		From("general_forum").
		Where(sb.Equal("id", id))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	var discussion models.Discussion
	err := config.DB.QueryRow(ctx, query, args...).Scan(
		&discussion.ID,
		&discussion.Student_ID,
		&discussion.Topic,
		&discussion.Description,
		&discussion.Replies,
	)
	if err != nil {
		return nil, err
	}

	return &discussion, nil
}

func (r *DiscussionRepository) ReplyDiscussion(ctx context.Context, discussionID int, replies string, studentID int, studentName string) (models.Discussion, error) {
	// Ambil existing replies dari discussion
	var existingRepliesJSON []byte
	err := config.DB.QueryRow(ctx, "SELECT replies FROM general_forum WHERE id = $1", discussionID).Scan(&existingRepliesJSON)

	// Cek apakah terjadi error lainnya selain tidak ditemukan baris
	if err != nil && err != sql.ErrNoRows {
		return models.Discussion{}, err
	}

	// Cek apakah existingRepliesJSON kosong (NULL atau empty)
	var existingReplies []map[string]interface{}
	if err == sql.ErrNoRows || len(existingRepliesJSON) == 0 {
		// Jika baris tidak ditemukan atau kosong, inisialisasi dengan array kosong
		existingReplies = []map[string]interface{}{}
	} else {
		// Decode existing replies jika ada
		if err := json.Unmarshal(existingRepliesJSON, &existingReplies); err != nil {
			return models.Discussion{}, err
		}
	}

	// Buat reply baru
	newReply := map[string]interface{}{
		"id":    studentID,
		"name":  studentName,
		"time":  time.Now().Format(time.RFC3339),
		"reply": replies,
	}

	// Append reply baru ke existing replies
	existingReplies = append(existingReplies, newReply)

	// Encode kembali ke JSON
	updatedRepliesJSON, err := json.Marshal(existingReplies)
	if err != nil {
		return models.Discussion{}, err
	}

	// Update discussion dengan replies terbaru
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("general_forum").
		Set(
			sb.Assign("replies", updatedRepliesJSON),
		).
		Where(sb.Equal("id", discussionID))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	query += " RETURNING id, student_id, topic, description, replies"

	var discussion models.Discussion
	err = config.DB.QueryRow(ctx, query, args...).Scan(
		&discussion.ID,
		&discussion.Student_ID,
		&discussion.Topic,
		&discussion.Description,
		&discussion.Replies,
	)
	if err != nil {
		return models.Discussion{}, err
	}

	return discussion, nil
}
