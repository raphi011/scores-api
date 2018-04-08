package scores

type Group struct {
	Model
	Players  Players `json:"players"`
	Name     string  `json:"name"`
	Role     string  `json:"role"`
	ImageURL string  `json:"imageUrl"`
}

type Groups []Group

type GroupService interface {
	GroupsByPlayer(playerID uint) (Groups, error)
	Groups() (Groups, error)
	Group(groupID uint) (*Group, error)
	Create(*Group) (*Group, error)
	AddPlayerToGroup(playerID, groupID uint, role string) error
}
