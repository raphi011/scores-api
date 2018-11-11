package scores

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	passwordService := &PBKDF2PasswordService{
		SaltBytes:  16,
		Iterations: 10000,
	}

	pw := []byte("password")

	info, err := passwordService.Hash(pw)

	if err != nil {
		t.Errorf("PBKDF2PasswordService.Hash(\"password\"), err : %s", err)
	}

	if !passwordService.Compare(pw, info) {
		t.Error("PBKDF2PasswordService.Compare(), want true, got false")
	}
}
