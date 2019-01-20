package scores

// User represents a user in the repository
type User struct {
	Model
	Tracked
	Email           string       `json:"email"`
	ProfileImageURL string       `json:"profileImageUrl" db:"profile_image_url"`
	VolleynetUserID int          `json:"volleynetUserId" db:"volleynet_user_id"`
	VolleynetUser  string        `json:"volleynetUser" db:"volleynet_user"`
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
