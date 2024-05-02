package postgre

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitPostgre(dbDriver, dbSource string) *sql.DB {

	pg, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	err = pg.Ping()
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)

	}
	pg.SetMaxIdleConns(10)
	pg.SetMaxOpenConns(100)
	pg.SetConnMaxIdleTime(5 * time.Minute)
	pg.SetConnMaxLifetime(60 * time.Minute)

	return pg
}
