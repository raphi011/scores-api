package scores

// Player represents a player
type Player struct {
	Model
	Name            string  `json:"name"`
	UserID          uint    `json:"userId"`
	ProfileImageURL string  `json:"profileImageUrl"`
	Groups          []Group `json:"groups"`
}

// PlayerRepository persists and retrieves Players
type PlayerRepository interface {
	Get(playerID uint) (*Player, error)
	All() ([]Player, error)
	ByGroup(groupID uint) ([]Player, error)
	Create(*Player) (*Player, error)
	Delete(playerID uint) error
}
