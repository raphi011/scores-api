package scores

type User struct {
	Model
	Email           string       `json:"email"`
	ProfileImageURL string       `json:"profileImageUrl"`
	VolleynetUserID int          `json:"volleynetUserId"`
	VolleynetLogin  string       `json:"volleynetLogin"`
	Role            string       `json:"role"`
	PlayerID        uint         `json:"playerId"`
	Player          *Player      `json:"player"`
	PasswordInfo    PasswordInfo `json:"-"`
}

type Users []User

type UserRepository interface {
	New(*User) (*User, error)
	ByID(userID uint) (*User, error)
	ByEmail(email string) (*User, error)
	Update(*User) error
	All() (Users, error)
}
