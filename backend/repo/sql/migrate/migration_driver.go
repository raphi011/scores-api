package migrate

import (
	"io/ioutil"
	"bytes"
	"os"
	"log"
	"fmt"
	"io"
	"path"
	"strings"

    "github.com/golang-migrate/migrate/v4/source"
	"github.com/gobuffalo/packr/v2"
)

type packrDriver struct {
	migrations *source.Migrations
	provider string
	box *packr.Box
}

var _ source.Driver = &packrDriver{}

func (p * packrDriver) Open(provider string) (source.Driver, error) {
	migrationSource := source.NewMigrations()
	migrationPaths := migrations.List()

	if len(migrationPaths) == 0 {
		return nil, fmt.Errorf("no migrations available for %s", provider)
	}

	for _, migrationPath := range migrationPaths {
		if !strings.HasPrefix(migrationPath, provider) {
			continue
		}

		m, err := source.DefaultParse(path.Base(migrationPath))

		if err != nil {
			log.Fatalf("invalid migration file name: %s", migrationPath)
		}

		if !migrationSource.Append(m) {
			log.Fatalf("unable to parse file: %s", migrationPath)
		}
	}


	return &packrDriver{
		migrations: migrationSource,
		box: migrations,
		provider: provider,
	}, nil
}

func (p *packrDriver) Close() error {
	return nil
}
 
func (p *packrDriver) First() (version uint, err error) {
	if version, ok := p.migrations.First(); ok {
		return version, nil
	}

	return 0, &os.PathError{Op: fmt.Sprintf("First for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *packrDriver) Prev(version uint) (prevVersion uint, err error) {
	if version, ok := p.migrations.Prev(version); ok {
		return version, nil
	}

	return 0, &os.PathError{Op: fmt.Sprintf("Prev for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *packrDriver) Next(version uint) (nextVersion uint, err error) {
	if version, ok := p.migrations.Next(version); ok {
		return version, nil
	}

	return 0, &os.PathError{Op: fmt.Sprintf("Next for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *packrDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := p.migrations.Up(version); ok {
		migration, _ := p.box.Find(p.provider + "/" + m.Raw)

		return ioutil.NopCloser(bytes.NewReader(migration)), m.Identifier, nil
	}

	return nil, "", &os.PathError{Op: fmt.Sprintf("ReadUp for version %v", version), Path: "", Err: os.ErrNotExist}
}

func (p *packrDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := p.migrations.Down(version); ok {
		migration, _ := p.box.Find(p.provider + "/" + m.Raw)

		return ioutil.NopCloser(bytes.NewReader(migration)), m.Identifier, nil
	}

	return nil, "", &os.PathError{Op: fmt.Sprintf("ReadDown for version %v", version), Path: "", Err: os.ErrNotExist}
}
