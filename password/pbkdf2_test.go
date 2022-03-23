package password

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	password := &PBKDF2{
		SaltBytes:  16,
		Iterations: 10000,
	}

	pw := []byte("password")

	info, err := password.Hash(pw)

	if err != nil {
		t.Errorf("PBKDF2.Hash(\"password\"), err : %s", err)
	}

	if !password.Compare(pw, info) {
		t.Error("PBKDF2.Compare(), want true, got false")
	}
}
