package sql

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func setupDB(t testing.TB) *sqlx.DB {
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
	assert(t, "unable to open db: %v", err)

	err = db.Ping()
	assert(t, "unable to connect to db: %v", err)

	err = migrateAll(dbProvider, db)
	assert(t, "migration failed: %v", err)

	_, err = exec(db, "test/delete-all", make(map[string]interface{}));
	assert(t, "db cleanup failed: %v", err)

	return db
}

func assert(t testing.TB, message string, err error) {
	if err != nil {
		t.Fatalf(message, err)
	}
}