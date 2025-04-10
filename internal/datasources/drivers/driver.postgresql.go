package drivers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

func GetDbConnString(dbusername string, dbpassword string, dbhost string, dbport string, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbusername,
		dbpassword,
		dbhost,
		dbport,
		dbname)
}

func SetupPostgresConnection(connectionString string) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(connectionString)

	if err != nil {
		log.Fatalf("Invalid database connection string: %s\n", err)
	}

	conf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	return pgxpool.NewWithConfig(context.Background(), conf)
}
