package scores

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	pwRepository := &PBKDF2PasswordRepository{
		SaltBytes:  16,
		Iterations: 10000,
	}

	pw := []byte("password")

	info, err := pwRepository.HashPassword(pw)

	if err != nil {
		t.Errorf("PBKDF2PasswordRepository.HashPassword(\"password\"), err : %s", err)
	}

	if !pwRepository.ComparePassword(pw, info) {
		t.Error("PBKDF2PasswordRepository.ComparePassword(), want true, got false")
	}
}
