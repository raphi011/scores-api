package scores

// User represents a user in the repository
type User struct {
	Model
	Email           string       `json:"email"`
	ProfileImageURL string       `json:"profileImageUrl"`
	VolleynetUserID int          `json:"volleynetUserId"`
	VolleynetLogin  string       `json:"volleynetLogin"`
	Role            string       `json:"role"`
	PasswordInfo    PasswordInfo `json:"-"`
}

// Users is a slice of multiple users
type Users []User

// UserRepository stores users
type UserRepository interface {
	New(*User) (*User, error)
	ByID(userID uint) (*User, error)
	ByEmail(email string) (*User, error)
	Update(*User) error
	All() (Users, error)
}
