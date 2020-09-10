package scores

import (
	"strings"

	"github.com/google/uuid"
)

// Settings maps Keys to Settings
type Settings map[string]interface{}

// Setting represents a user setting in the repository
type Setting struct {
	Track  `json:"-"`
	UserID uuid.UUID `json:"-" db:"user_id"`
	Key    string    `json:"key" db:"s_key"`
	Value  string    `json:"-" db:"s_value"`
	Type   string    `json:"-" db:"s_type"`
}

const (
	separator = ","
)

// Val converts a string Value to it's real type
func (s Setting) Val() interface{} {
	switch s.Type {
	case "string":
		return s.Value
	case "strings":
		return StringToList(s.Value)
	default:
		return nil
	}
}

// ToSettingsDictionary creates a key-value dictionary out of a list of settings
func ToSettingsDictionary(settings []*Setting) Settings {
	settingsMap := make(Settings)

	for _, setting := range settings {
		settingsMap[setting.Key] = setting.Val()
	}

	return settingsMap
}

// ListToString converts a list to a string
func ListToString(list []string) string {
	b := strings.Builder{}

	for i, s := range list {
		if i != 0 {
			b.WriteString(separator)
		}

		b.WriteString(s)
	}

	return b.String()
}

// StringToList converts a string to a list
func StringToList(listString string) []string {
	return strings.Split(listString, separator)
}
