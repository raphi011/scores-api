package scores

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

var _ PasswordRepository = &PBKDF2PasswordRepository{}

type PasswordInfo struct {
	Salt       []byte
	Hash       []byte
	Iterations int
}

type PasswordRepository interface {
	ComparePassword([]byte, *PasswordInfo) bool
	HashPassword(password []byte) (*PasswordInfo, error)
}

type PBKDF2PasswordRepository struct {
	SaltBytes  int
	Iterations int
}

func (s *PBKDF2PasswordRepository) newSalt() ([]byte, error) {
	salt := make([]byte, s.SaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func (s *PBKDF2PasswordRepository) ComparePassword(password []byte, info *PasswordInfo) bool {
	hash := s.hashPassword(password, info.Salt, info.Iterations)

	return bytes.Compare(hash, info.Hash) == 0
}

func (s *PBKDF2PasswordRepository) HashPassword(password []byte) (*PasswordInfo, error) {
	salt, err := s.newSalt()

	if err != nil {
		return nil, err
	}

	hash := s.hashPassword(password, salt, s.Iterations)

	return &PasswordInfo{
		Salt:       salt,
		Hash:       hash,
		Iterations: s.Iterations,
	}, nil
}

func (s *PBKDF2PasswordRepository) hashPassword(password, salt []byte, iterations int) []byte {
	hash := pbkdf2.Key(password, salt, iterations, 32, sha256.New)

	return hash
}
