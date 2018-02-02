package scores

type Group struct {
	Model
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

type Groups []Group

type GroupService interface {
	GroupsByPlayer(playerID uint) (Groups, error)
	Groups() (Groups, error)
	Group(groupID uint) (*Group, error)
	Create(*Group) (*Group, error)
	AddPlayerToGroup(playerID uint) error
}
