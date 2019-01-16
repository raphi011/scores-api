package scores

// User represents a user in the repository
type User struct {
	TrackedModel
	Email           string       `json:"email"`
	ProfileImageURL string       `json:"profileImageUrl"`
	VolleynetUserID int          `json:"volleynetUserId"`
	VolleynetUser  string       `json:"volleynetLogin"`
	Role            string       `json:"role"`
	PasswordInfo
}

// PasswordInfo contains the passwords hash, it's corresponding salt
// and the amount of PBKDF2 iterations it was hashed with.
type PasswordInfo struct {
	Salt       []byte
	Hash       []byte
	Iterations int
}
