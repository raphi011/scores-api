package volleynet

import "github.com/raphi011/scores"

// TournamentTeam is the current status of the team entry in a
// tournament, if the tournament is finished it may also contain
// the seed.
type TournamentTeam struct {
	scores.Track

	TournamentID int     `json:"tournamentId" db:"tournament_id"`
	Player1      *Player `json:"player1"`
	Player2      *Player `json:"player2"`

	Deregistered bool    `json:"deregistered"`
	PrizeMoney   float32 `json:"prizeMoney" db:"prize_money"`
	Result       int     `json:"result"`
	Seed         int     `json:"seed"`
	TotalPoints  int     `json:"totalPoints" db:"total_points"`
	WonPoints    int     `json:"wonPoints" db:"won_points"`
}
