package utils

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
	"time"
	"context"
)

func CheckExamStatus(startTime, endTime time.Time) string {
	now := time.Now()
	switch {
	case now.Before(startTime):
		return "Scheduled"
	case now.After(endTime):
		return "Completed"
	default:
		return "Active"
	}
}

func UpdateAllExamStatus(db *pgxpool.Pool) error {
	ctx := context.Background()

	rows, err := db.Query(ctx, `SELECT id, start_time, end_time FROM exams`)
	if err != nil {
		return fmt.Errorf("failed to fetch exams: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var startTime, endTime time.Time
		if err := rows.Scan(&id, &startTime, &endTime); err != nil {
			return fmt.Errorf("failed to scan exams row: %w", err)
		}

		newStatus := CheckExamStatus(startTime, endTime)

		_, err := db.Exec(ctx, `UPDATE exams SET status = $1 WHERE id = $2`, newStatus, id)
		if err != nil {
			return fmt.Errorf("failed to update status for exams %d: %w", id, err)
		}
	}
	return nil
}
