package scores

type Player struct {
	Model
	Name   string `json:"name"`
	UserId uint
}

type Players []Player

type PlayerService interface {
	Player(playerID uint) (*Player, error)
	Players() (Players, error)
	Create(*Player) (*Player, error)
	Delete(playerID uint) error
}
