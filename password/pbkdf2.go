package password

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"github.com/raphi011/scores-api"
	"golang.org/x/crypto/pbkdf2"
)

// PBKDF2 contains parameters for the PBKDF2 algorithm
type PBKDF2 struct {
	SaltBytes  int
	Iterations int
}

func (s *PBKDF2) newSalt() ([]byte, error) {
	salt := make([]byte, s.SaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Compare takes a password and compares it to the hashed version and returns true
// if they are equal
func (s *PBKDF2) Compare(password []byte, info *scores.PasswordInfo) bool {
	hash := s.hash(password, info.Salt, info.Iterations)

	return bytes.Compare(hash, info.Hash) == 0
}

// Hash generates and new salt and hashes the passed password with it
func (s *PBKDF2) Hash(password []byte) (*scores.PasswordInfo, error) {
	salt, err := s.newSalt()

	if err != nil {
		return nil, err
	}

	hash := s.hash(password, salt, s.Iterations)

	return &scores.PasswordInfo{
		Salt:       salt,
		Hash:       hash,
		Iterations: s.Iterations,
	}, nil
}

func (s *PBKDF2) hash(password, salt []byte, iterations int) []byte {
	hash := pbkdf2.Key(password, salt, iterations, 32, sha256.New)

	return hash
}
