package sqlite

import (
	"testing"
)

func TestOpen(t *testing.T) {
	_, err := Open("file::memory:", "&mode=memory&cache=shared")

	if err != nil {
		t.Errorf("Error opening db: %s", err)
	}
}
