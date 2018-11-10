package sqlite

import (
	"github.com/gobuffalo/packr"
	log "github.com/sirupsen/logrus"
)

var queries packr.Box

func init() {
	queries = packr.NewBox("queries")
}

func query(name string) string {
	q, err := queries.FindString(name + ".sql")

	if err != nil {
		log.Fatalf("could not load sqlite query %s", name)
	}

	return q
}
