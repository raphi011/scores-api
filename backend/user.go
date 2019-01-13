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

// UserRepository stores users
type UserRepository interface {
	New(*User) (*User, error)
	ByID(userID int) (*User, error)
	ByEmail(email string) (*User, error)
	Update(*User) error
	All() ([]*User, error)
}
