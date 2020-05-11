package scores

import (
	"strings"
	"unicode"
)

// Sluggify transforms a string into a normalized representation.
// e.g. "whitespace and Capitalized" -> "whitespace-and-capitalized".
func Sluggify(text string) string {
	builder := strings.Builder{}
	builder.Grow(len(text))

	var lastChar rune

	for _, char := range text {
		last := lastChar
		lastChar = char

		isSpace := unicode.IsSpace(char)

		if isSpace {
			if !unicode.IsSpace(last) {
				builder.WriteRune('-')
			}

			continue
		}

		if last == '-' && char == '-' {
			continue
		}

		builder.WriteRune(unicode.ToLower(char))
	}

	return builder.String()
}
