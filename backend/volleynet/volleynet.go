package volleynet

import (
	"time"

	"github.com/raphi011/scores"
)

// PlayerInfo contains all player information that the search player api returns.
type PlayerInfo struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birthday  time.Time `json:"birthday"`
}

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

const (
	// StatusUpcoming represents the state of a tournament not done yet.
	StatusUpcoming = "upcoming"
	// StatusDone represents the state of a completed tournament.
	StatusDone = "done"
	// StatusCanceled represents the state of a canceled tournament.
	StatusCanceled = "canceled"
)

// TournamentInfo is all the information that can be parsed from the tournament list.
type TournamentInfo struct {
	ID               int       `json:"id"`
	Season           int       `json:"season"`
	Start            time.Time `json:"start" db:"start_date"`
	End              time.Time `json:"end" db:"end_date"`
	Name             string    `json:"name" fako:"city"`
	League           string    `json:"league" db:"league"`
	LeagueKey       string    `json:"leagueKey" db:"league_key"`
	SubLeague        string    `json:"subLeague" db:"sub_league"`
	SubLeagueKey    string    `json:"phaseKey" db:"sub_league_key"`
	Link             string    `json:"link" fako:"domain_name"`
	Status           string    `json:"status"` // can be `StatusUpcoming`, `StatusDone` or `StatusCanceled`
	Gender           string    `json:"gender"`
	RegistrationOpen bool      `json:"registrationOpen" db:"registration_open"`
}

// Tournament adds additional information to `TournamentInfo`.
type Tournament struct {
	scores.Track

	TournamentInfo

	EndRegistration *time.Time        `json:"endRegistration" db:"end_registration"`
	Teams           []*TournamentTeam `json:"teams"`
	Location        string            `json:"location"`
	HTMLNotes       string            `json:"htmlNotes" db:"html_notes" fako:"paragraph"`
	Mode            string            `json:"mode"`
	Organiser       string            `json:"organiser" fako:"full_name"`
	Phone           string            `json:"phone" fako:"phone"`
	Email           string            `json:"email" fako:"email_address"`
	Website         string            `json:"website"`
	CurrentPoints   string            `json:"currentPoints" db:"current_points"`
	LivescoringLink string            `json:"livescoringLink" db:"live_scoring_link"`
	SignedupTeams   int               `json:"signedupTeams" db:"signedup_teams"`
	MaxTeams        int               `json:"maxTeams" db:"max_teams"`
	MinTeams        int               `json:"minTeams" db:"min_teams"`
	MaxPoints       int               `json:"maxPoints" db:"max_points"`
	Latitude        float32           `json:"latitude" db:"loc_lat"`
	Longitude       float32           `json:"longitude" db:"loc_lon"`
}

// LoginData contains the data of the form that is shown after a successful login.
type LoginData struct {
	PlayerInfo
	License License
}

// License contains the license data of a player.
type License struct {
	Nr        string
	Type      string
	Requested string
}
