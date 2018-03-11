package scores

type Player struct {
	Model
	Name            string `json:"name"`
	UserID          uint   `json:"userId"`
	ProfileImageURL string `json:"profileImageUrl"`
}

type Players []Player

type PlayerService interface {
	Player(playerID uint) (*Player, error)
	Players() (Players, error)
	ByGroup(groupID uint) (Players, error)
	Create(*Player) (*Player, error)
	Delete(playerID uint) error
}
