package drivers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
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

func SetupPostgresConnection(connectionString string) (*pgx.Conn, error) {
	conf, err := pgx.ParseConfig(connectionString)

	println(conf.ConnString())

	//if err != nil {
	//	log.Fatalf("Invalid database connection string: %s\n", err)
	//}

	//conf.AfterConnect = func(ctx context.Context, conn *pgconn.PgConn) error {
	//	pgxuuid.Register(conn.TypeMap())
	//	return nil
	//}

	println("before connect")

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	println("afer connect")

	return conn, err

	//return pgxpool.NewWithConfig(context.Background(), conf)
}
