package utils

import (
	"github.com/robfig/cron/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
)

func StartCron(db *pgxpool.Pool) {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		if err := UpdateAllExamStatus(db); err != nil {
			fmt.Println("Error updating exam status:", err)
		}
	})
	c.Start()
}
