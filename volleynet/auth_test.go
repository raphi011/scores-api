package volleynet

import (
	"testing"
)

func Test_login(t *testing.T) {
	t.Skip() // Add Testing credentials to env

	c := DefaultClient()
	err := c.Login("", "")

	if err != nil {
		t.Error(err)
	}
}
