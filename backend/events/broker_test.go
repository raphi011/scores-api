package events

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExpandPossibleHandlers(t *testing.T) {
	eventName := "volleynet/sync/new-tournament"

	eventHandlers := []string{
		"volleynet/*",
		"volleynet/sync/*",
		"volleynet/sync/new-tournament",
	}

	output := expandPossibleHandlers(eventName)

	if !cmp.Equal(eventHandlers, output) {
		t.Errorf("expandPossibleHandlers() err: wrong output:\n%s", cmp.Diff(eventHandlers, output))
	}

}
