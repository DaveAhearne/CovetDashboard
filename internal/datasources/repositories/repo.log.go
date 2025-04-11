package repositories

import (
	"context"
	"covet.digital/dashboard/internal/business/domains"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"log"
)

type logRepository struct {
	db *pgx.Conn
}

func NewLogRepository(db *pgx.Conn) domains.LogRepository {
	return &logRepository{
		db: db,
	}
}

func (l logRepository) Listen(ctx context.Context) (<-chan domains.LogDomain, error) {
	logs := make(chan domains.LogDomain)

	log.Println("Listening for database notifications...")
	if _, err := l.db.Exec(ctx, "LISTEN log_event"); err != nil {
		log.Fatalf("Failed to listen for notifications: %v", err)
	}

	go func() {
		for {
			notification, err := l.db.WaitForNotification(ctx)
			if err != nil {
				log.Println("WaitForNotification error:", err)
				continue
			}

			println("Got a notification!")

			var logEvent domains.LogDomain
			if err := json.Unmarshal([]byte(notification.Payload), &logEvent); err != nil {
				log.Println("Error deserializing log payload", err)
				continue
			}

			logs <- logEvent
		}
	}()

	return logs, nil
}
