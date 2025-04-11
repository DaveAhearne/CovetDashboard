package repositories

import (
	"context"
	"covet.digital/dashboard/internal/business/domains"
	"covet.digital/dashboard/internal/datasources/drivers"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

type logRepository struct {
	db drivers.PgxIface
}

func NewLogRepository(db drivers.PgxIface) domains.LogRepository {
	return &logRepository{
		db: db,
	}
}

func (l logRepository) GetEventsAfter(ctx context.Context, date time.Time) ([]domains.LogDomain, error) {
	tx, err := l.db.Begin(ctx)

	if err != nil {
		return nil, err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Commit(ctx)
		if err != nil {
			log.Printf("Error committing transaction (GetEventsAfter): %s", err)
		}
	}(tx, ctx)

	rows, err := tx.Query(ctx, `
		SELECT id, category, account_id, data, created_at
		FROM log
		WHERE created_at>$1`, date)

	if err != nil {
		return []domains.LogDomain{}, err
	}

	logs, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (domains.LogDomain, error) {
		var log domains.LogDomain
		err := row.Scan(&log.Id, &log.Category, &log.AccountId, &log.Data, &log.CreatedAt)
		return log, err
	})

	if err != nil {
		return []domains.LogDomain{}, err
	}

	return logs, nil
}

func (l logRepository) Listen(ctx context.Context) (<-chan domains.LogDomain, error) {
	logs := make(chan domains.LogDomain)

	poolConn, err := l.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	conn := poolConn.Conn()

	log.Println("Listening for database notifications...")
	if _, err := conn.Exec(ctx, "LISTEN log_event"); err != nil {
		log.Fatalf("Failed to listen for notifications: %v", err)
	}

	go func() {
		for {
			notification, err := conn.WaitForNotification(ctx)
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
