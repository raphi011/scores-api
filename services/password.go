package services

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"github.com/raphi011/scores-backend"
	"golang.org/x/crypto/pbkdf2"
)

var _ Password = &PBKDF2Password{}

// Password allows comparing hashed passwords and creating new ones
type Password interface {
	Compare([]byte, *scores.PasswordInfo) bool
	Hash(password []byte) (*scores.PasswordInfo, error)
}

// PBKDF2Password contains parameters for the PBKDF2 algorithm
type PBKDF2Password struct {
	SaltBytes  int
	Iterations int
}

func (s *PBKDF2Password) newSalt() ([]byte, error) {
	salt := make([]byte, s.SaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Compare takes a password and compares it to the hashed version and returns true
// if they are equal
func (s *PBKDF2Password) Compare(password []byte, info *scores.PasswordInfo) bool {
	hash := s.hash(password, info.Salt, info.Iterations)

	return bytes.Compare(hash, info.Hash) == 0
}

// Hash generates and new salt and hashes the passed password with it
func (s *PBKDF2Password) Hash(password []byte) (*scores.PasswordInfo, error) {
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

func (s *PBKDF2Password) hash(password, salt []byte, iterations int) []byte {
	hash := pbkdf2.Key(password, salt, iterations, 32, sha256.New)

	return hash
}
