package packr

import (
	"os"
	"path"

	"github.com/gobuffalo/packr/v2"
)

func New(name, relativePath string) *packr.Box {
	p := relativePath
	root := os.Getenv("TEST_FILES")

	if root != "" {
		p = path.Join(root, relativePath)
	}

	return packr.New(name, p)
}