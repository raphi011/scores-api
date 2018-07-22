package volleynet

import (
	"io"

	"github.com/pkg/errors"
)

func parseUniqueWriteCode(html io.Reader) (string, error) {
	doc, err := parseHTML(html)

	if err != nil {
		return "", errors.Wrap(err, "parseUniqueWriteCode failed")
	}

	input := doc.Find("input[name='XX_unique_write_XXBeach/Profile/TurnierAnmeldung']")

	val, exists := input.Attr("value")

	if !exists {
		return "", errors.New("parseUniqueWriteCode failed, code not found")
	}

	return val, nil
}
