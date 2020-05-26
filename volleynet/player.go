package volleynet

import (
	"time"

	"github.com/raphi011/scores-api"
)

// Player contains all relevent volleynet player information.
type Player struct {
	ID int `json:"id"`

	scores.Track

	Birthday     *time.Time `json:"birthday"`
	Club         string     `json:"club"`
	CountryUnion string     `json:"countryUnion" db:"country_union"`
	FirstName    string     `json:"firstName" db:"first_name"`
	Gender       string     `json:"gender"`
	LadderRank   int        `json:"ladderRank" db:"ladder_rank"`
	LastName     string     `json:"lastName" db:"last_name"`
	License      string     `json:"license"`
	TotalPoints  int        `json:"totalPoints" db:"total_points"`
}
