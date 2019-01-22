package scores

import (
	"strings"
	"unicode"
)

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
