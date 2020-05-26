package events

import (
	"testing"

	"github.com/raphi011/scores-api/test"
)

func TestExpandPossibleHandlers(t *testing.T) {
	eventName := "volleynet/sync/new-tournament"

	eventHandlers := []string{
		"volleynet/*",
		"volleynet/sync/*",
		"volleynet/sync/new-tournament",
	}

	output := expandPossibleHandlers(eventName)

	test.Compare(t, "expandPossibleHandlers() err: wrong output:\n%s", eventHandlers, output)
}
