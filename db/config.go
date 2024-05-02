package db

import (
	"database/sql"
	"golang-base/db/postgre"

	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Database(dbDriver, dbSource string) *sql.DB {
	return postgre.InitPostgre(dbDriver, dbSource)
}

func RunDBMigration(migrationURL, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalf("cannot create new migration instance: %v", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrate up : %v", err)
	}

	log.Print("db migration successfully")
}
