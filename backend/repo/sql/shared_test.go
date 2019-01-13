package sql

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func setupDB(t *testing.T) *sqlx.DB {
	t.Helper()

	dbProvider := "sqlite3"
	connectionString := "file::memory:?_busy_timeout=5000&mode=memory"

	p := os.Getenv("TEST_DB_PROVIDER")
	c := os.Getenv("TEST_DB_CONNECTION")

	if p != "" && c != "" {
		dbProvider = p
		connectionString = c
	}

	db, err := sqlx.Open(dbProvider, connectionString)

	if err != nil {
		t.Fatalf("unable to open db: %v", err)
	}

	err = db.Ping()

	if err != nil {
		t.Fatalf("unable to connect to db: %v", err)
	}

	if err = migrateAll(dbProvider, db); err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	if _, err := exec(db, "test/delete-all", make(map[string]interface{})); err != nil {
		t.Fatalf("db cleanup failed: %v", err)
	}
	
	return db
}
