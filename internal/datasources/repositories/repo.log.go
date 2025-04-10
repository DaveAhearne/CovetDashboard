package repositories

import (
	"context"
	"covet.digital/dashboard/internal/business/domains"
	"covet.digital/dashboard/internal/datasources/drivers"
	"encoding/json"
	"log"
)

type logRepository struct {
	db drivers.PgxIface
}

func NewLogRepository(db drivers.PgxIface) domains.LogRepository {
	return &logRepository{
		db: db,
	}
}

func (l logRepository) Listen(ctx context.Context) (<-chan domains.LogDomain, error) {
	logs := make(chan domains.LogDomain)

	tx, err := l.db.Begin(ctx)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening for database notifications...")
	if _, err := tx.Conn().Exec(ctx, "LISTEN log_event"); err != nil {
		log.Fatalf("Failed to listen for notifications: %v", err)
	}

	go func() {
		for {
			notification, err := tx.Conn().WaitForNotification(ctx)
			if err != nil {
				log.Println("WaitForNotification error:", err)
				continue
			}

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
