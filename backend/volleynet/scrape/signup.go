package scrape

import (
	"io"

	"github.com/pkg/errors"
)

// UniqueWriteCode extracts the XSRF token from the tournament login page
func UniqueWriteCode(html io.Reader) (string, error) {
	doc, err := parseHTML(html)

	if err != nil {
		return "", errors.Wrap(err, "UniqueWriteCode failed")
	}

	input := doc.Find("input[name='XX_unique_write_XXBeach/Profile/TurnierAnmeldung']")

	val, exists := input.Attr("value")

	if !exists {
		return "", errors.New("UniqueWriteCode code not found")
	}

	return val, nil
}
