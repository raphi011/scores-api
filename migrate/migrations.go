package migrate

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
	"strings"
)

const (
	versionTable = `
		CREATE TABLE "dbVersion" (
			"version" integer NOT NULL
		)
	`

	insertVersion1SQL = "INSERT INTO dbVersion VALUES (1)"
)

type Migration []string
type TableNames []string

func Reset(db *sql.DB, resets []TableNames) error {
	dbVersion, err := GetDBVersion(db)

	if err != nil {
		return err
	}

	if dbVersion == 0 {
		return nil
	}

	for _, tableName := range resets[dbVersion-1] {
		_, err := db.Exec("DELETE FROM " + tableName)

		if err != nil {
			return err
		}
	}

	return nil
}

func Migrate(db *sql.DB, migrations []Migration) error {
	dbVersion, err := GetDBVersion(db)

	if err != nil {
		return err
	}

	newestVersion := uint16(len(migrations))

	if dbVersion > newestVersion {
		return errors.New("Unknown version number")
	}

	if newestVersion != dbVersion {
		err = runMigrations(db, migrations, dbVersion)
	}

	return err
}

func GetDBVersion(db *sql.DB) (uint16, error) {
	var dbVersion uint16

	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='dbVersion'").Scan()

	if err == sql.ErrNoRows {
		// version table doesn't exist
		return 0, nil
	}

	row := db.QueryRow("SELECT version FROM dbVersion")

	err = row.Scan(&dbVersion)

	if err == sql.ErrNoRows {
		// version row somehow is missing, don't assume anything ...
		return 0, errors.New("DB version missing")
	} else if err != nil {
		return 0, err
	}

	return dbVersion, nil
}

func runMigrations(db *sql.DB, migrations []Migration, fromVersion uint16) error {
	migrationQueue := migrations[fromVersion:]
	currentVersion := fromVersion + 1
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	for _, m := range migrationQueue {
		log.Printf("Applying migration v%d", currentVersion)

		if currentVersion == 1 {
			m = append(m, versionTable, insertVersion1SQL)
		}

		err = execMultiple(tx, m...)

		if err != nil {
			log.Fatalf("An error occured, rolling back: %v", err)
			_ = tx.Rollback()
			return err
		}

		err := setDBVersion(tx, currentVersion)

		if err != nil {
			log.Fatalf("An error occured, rolling back: %v", err)
		}

		currentVersion++
	}

	err = tx.Commit()

	return err
}

func execMultiple(db *sql.Tx, statements ...string) error {
	for _, statement := range statements {
		innerStatements := strings.Split(statement, ";")

		for _, s := range innerStatements {
			_, err := db.Exec(s)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func setDBVersion(db *sql.Tx, version uint16) error {
	_, err := db.Exec("UPDATE dbVersion SET version=$1", version)

	return err
}
