package scores

type User struct {
	Model
	Email           string `json:"email"`
	Player          Player `json:"player"`
	PlayerID        uint   `json:"playerId"`
	ProfileImageURL string `json:"profileImageUrl"`
}

type Users []User

type UserService interface {
	Create(*User) (*User, error)
	User(userID uint) (*User, error)
	ByEmail(email string) (*User, error)
	Update(*User) error
}
