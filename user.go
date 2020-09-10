package scores

import (
	"github.com/google/uuid"
)

// User represents a user in the repository
type User struct {
	Track
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	ProfileImageURL string    `json:"profileImageUrl" db:"profile_image_url"`
	PlayerID        int       `json:"playerId" db:"player_id"`
	PlayerLogin     string    `json:"playerLogin" db:"player_login"`
	Role            string    `json:"role"`
	Settings        Settings  `json:"settings"`
	PasswordInfo
}

// PasswordInfo contains the passwords hash, it's corresponding salt
// and the amount of PBKDF2 iterations it was hashed with.
type PasswordInfo struct {
	Salt       []byte `db:"pw_salt"`
	Hash       []byte `db:"pw_hash"`
	Iterations int    `db:"pw_iterations"`
}
