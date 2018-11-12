package scores

// Group represents a group of players and the matches they've played
type Group struct {
	Model
	Players          []Player          `json:"players"`
	Matches          []Match           `json:"matches"`
	PlayerStatistics []PlayerStatistic `json:"playerStatistics"`
	Name             string            `json:"name"`
	Role             string            `json:"role"`
	ImageURL         string            `json:"imageUrl"`
}

// GroupRepository persists and retrieves Groups
type GroupRepository interface {
	ByPlayer(playerID uint) ([]Group, error)
	All() ([]Group, error)
	Get(groupID uint) (*Group, error)
	Create(*Group) (*Group, error)
	AddPlayer(playerID, groupID uint, role string) error
}
