package scores

import (
	"time"
)

type Match struct {
	Model
	Team1       *Team  `json:"team1"`
	Team2       *Team  `json:"team2"`
	ScoreTeam1  int    `json:"scoreTeam1"`
	ScoreTeam2  int    `json:"scoreTeam2"`
	TargetScore int    `json:"targetScore"`
	CreatedBy   *User  `json:"createdBy"`
	Group       *Group `json:"group"`
	GroupID     uint   `json:"groupId"`
}

type Matches []Match

type MatchRepository interface {
	GroupMatches(groupID uint, after time.Time, count uint) (Matches, error)
	Match(matchID uint) (*Match, error)
	PlayerMatches(playerID uint, after time.Time, count uint) (Matches, error)
	Matches(after time.Time, count uint) (Matches, error)
	Create(*Match) (*Match, error)
	Delete(matchID uint) error
}
