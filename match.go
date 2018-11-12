package scores

import (
	"time"
)

// Match represents a played match
type Match struct {
	Model
	Team1           *Team  `json:"team1"`
	Team2           *Team  `json:"team2"`
	ScoreTeam1      int    `json:"scoreTeam1"`
	ScoreTeam2      int    `json:"scoreTeam2"`
	TargetScore     int    `json:"targetScore"`
	CreatedByUserID uint   `json:"createdBy"`
	Group           *Group `json:"group"`
	GroupID         uint   `json:"groupId"`
}

// MatchRepository persists and retrieves Matches
type MatchRepository interface {
	ByGroup(groupID uint, after time.Time, count uint) ([]Match, error)
	Get(matchID uint) (*Match, error)
	ByPlayer(playerID uint, after time.Time, count uint) ([]Match, error)
	After(after time.Time, count uint) ([]Match, error)
	Create(*Match) (*Match, error)
	Delete(matchID uint) error
}
