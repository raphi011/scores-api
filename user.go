package scores

type User struct {
	Model
	Email           string       `json:"email"`
	PlayerID        uint         `json:"playerId"`
	ProfileImageURL string       `json:"profileImageUrl"`
	PasswordInfo    PasswordInfo `json:"-"`
}

type Users []User

type UserService interface {
	UpdatePasswordAuthentication(uint, *PasswordInfo) error
	Create(*User) (*User, error)
	User(userID uint) (*User, error)
	ByEmail(email string) (*User, error)
	Update(*User) error
}
