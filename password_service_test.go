package scores

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	pwRepository := &PBKDF2PasswordService{
		SaltBytes:  16,
		Iterations: 10000,
	}

	pw := []byte("password")

	info, err := pwRepository.HashPassword(pw)

	if err != nil {
		t.Errorf("PBKDF2PasswordService.HashPassword(\"password\"), err : %s", err)
	}

	if !pwRepository.ComparePassword(pw, info) {
		t.Error("PBKDF2PasswordService.ComparePassword(), want true, got false")
	}
}
