package scores

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	password := &PBKDF2Password{
		SaltBytes:  16,
		Iterations: 10000,
	}

	pw := []byte("password")

	info, err := password.Hash(pw)

	if err != nil {
		t.Errorf("PBKDF2Password.Hash(\"password\"), err : %s", err)
	}

	if !password.Compare(pw, info) {
		t.Error("PBKDF2Password.Compare(), want true, got false")
	}
}
