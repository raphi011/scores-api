package assets

import (
	"io"
	"io/ioutil"

	"github.com/markbates/pkger"
)

const (
	migrationPath = "/repo/sql/assets/migrations"
	queryPath     = "/repo/sql/assets/queries"
)

func init() {
	pkger.Include("/repo/sql/assets/migrations")
	pkger.Include("/repo/sql/assets/queries")
}

func MigrationsFolderPath(provider string) string {
	return migrationPath + "/" + provider
}

func ListDirectoryFileNames(folderName string) ([]string, error) {
	dir, err := pkger.Open(folderName)

	if err != nil {
		return nil, err
	}

	files, err := dir.Readdir(0)

	if err != nil {
		return nil, err
	}

	fileNames := make([]string, len(files))

	for i, f := range files {
		fileNames[i] = dir.Path().Name + "/" + f.Name()
	}

	return fileNames, nil
}

func LoadFile(name string) (io.ReadCloser, error) {
	return pkger.Open(name)
}

func readerToString(reader io.ReadCloser) (string, error) {
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func LoadFileString(name string) (string, error) {
	reader, err := LoadFile(name)

	if err != nil {
		return "", err
	}

	return readerToString(reader)
}

func LoadQueryString(name string) (string, error) {
	return LoadFileString(queryPath + "/" + name)
}
