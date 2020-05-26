package migrate

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/raphi011/scores-api/repo/sql/assets"

	"github.com/golang-migrate/migrate/v4/source"
)

type pkgerDriver struct {
	migrations *source.Migrations
	folder     string
}

var _ source.Driver = &pkgerDriver{}

func (p *pkgerDriver) Open(provider string) (source.Driver, error) {
	migrationSource := source.NewMigrations()

	folder := assets.MigrationsFolderPath(provider)

	migrationPaths, err := assets.ListDirectoryFileNames(folder)

	if err != nil {
		return nil, errors.Wrapf(err, "an error occured while loading %s migration files from pkger", provider)
	}

	if len(migrationPaths) == 0 {
		return nil, fmt.Errorf("no migrations available for %s", provider)
	}

	for _, migrationPath := range migrationPaths {
		m, err := source.DefaultParse(path.Base(migrationPath))

		if err != nil {
			log.Fatalf("invalid migration file name: %s", migrationPath)
		}

		if !migrationSource.Append(m) {
			log.Fatalf("unable to parse file: %s", migrationPath)
		}
	}

	return &pkgerDriver{
		migrations: migrationSource,
		folder:     folder,
	}, nil
}

func (p *pkgerDriver) Close() error {
	return nil
}

func (p *pkgerDriver) First() (version uint, err error) {
	if version, ok := p.migrations.First(); ok {
		return version, nil
	}

	return 0, &os.PathError{Op: fmt.Sprintf("First for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *pkgerDriver) Prev(version uint) (prevVersion uint, err error) {
	if version, ok := p.migrations.Prev(version); ok {
		return version, nil
	}

	return 0, &os.PathError{Op: fmt.Sprintf("Prev for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *pkgerDriver) Next(version uint) (nextVersion uint, err error) {
	if version, ok := p.migrations.Next(version); ok {
		return version, nil
	}

	return 0, &os.PathError{Op: fmt.Sprintf("Next for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *pkgerDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := p.migrations.Up(version); ok {
		migration, _ := assets.LoadFile(p.folder + "/" + m.Raw)

		return migration, m.Identifier, nil
	}

	return nil, "", &os.PathError{Op: fmt.Sprintf("ReadUp for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *pkgerDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := p.migrations.Down(version); ok {
		migration, _ := assets.LoadFile(p.folder + "/" + m.Raw)

		return migration, m.Identifier, nil
	}

	return nil, "", &os.PathError{Op: fmt.Sprintf("ReadDown for version %v", version), Path: "", Err: os.ErrNotExist}
}
