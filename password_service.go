package scores

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

var _ PasswordService = &PBKDF2PasswordService{}

type PasswordInfo struct {
	Salt       []byte
	Hash       []byte
	Iterations int
}

type PasswordService interface {
	Compare([]byte, *PasswordInfo) bool
	Hash(password []byte) (*PasswordInfo, error)
}

type PBKDF2PasswordService struct {
	SaltBytes  int
	Iterations int
}

func (s *PBKDF2PasswordService) newSalt() ([]byte, error) {
	salt := make([]byte, s.SaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Compare takes a password and compares it to the hashed version and returns true
// if they are equal
func (s *PBKDF2PasswordService) Compare(password []byte, info *PasswordInfo) bool {
	hash := s.hash(password, info.Salt, info.Iterations)

	return bytes.Compare(hash, info.Hash) == 0
}

// Hash generates and new salt and hashes the passed password with it
func (s *PBKDF2PasswordService) Hash(password []byte) (*PasswordInfo, error) {
	salt, err := s.newSalt()

	if err != nil {
		return nil, err
	}

	hash := s.hash(password, salt, s.Iterations)

	return &PasswordInfo{
		Salt:       salt,
		Hash:       hash,
		Iterations: s.Iterations,
	}, nil
}

func (s *PBKDF2PasswordService) hash(password, salt []byte, iterations int) []byte {
	hash := pbkdf2.Key(password, salt, iterations, 32, sha256.New)

	return hash
}
