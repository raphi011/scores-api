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
	scores.Tracked

	ID           int        `json:"id"`
	FirstName    string     `json:"firstName" db:"first_name"`
	LastName     string     `json:"lastName" db:"last_name"`
	Birthday     *time.Time `json:"birthday"`
	Gender       string     `json:"gender"`
	TotalPoints  int        `json:"totalPoints" db:"total_points"`
	LadderRank   int        `json:"ladderRank" db:"ladder_rank"`
	Club         string     `json:"club"`
	CountryUnion string     `json:"countryUnion" db:"country_union"`
	License      string     `json:"license"`
}

// TournamentTeam is the current status of the team entry in a
// tournament, if the tournament is finished it may also contain
// the seed.
type TournamentTeam struct {
	scores.Tracked

	TournamentID int     `json:"tournamentId" db:"tournament_id"`
	TotalPoints  int     `json:"totalPoints" db:"total_points"`
	Seed         int     `json:"seed"`
	Result       int     `json:"result"`
	WonPoints    int     `json:"wonPoints" db:"won_points"`
	Player1      *Player `json:"player1"`
	Player2      *Player `json:"player2"`
	PrizeMoney   float32 `json:"prizeMoney" db:"prize_money"`
	Deregistered bool    `json:"deregistered"`
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
	League           string    `json:"league"`
	Phase            string    `json:"phase"`
	Link             string    `json:"link" fako:"domain_name"`
	EntryLink        string    `json:"entryLink" db:"entry_link"`
	Status           string    `json:"status"` // can be `StatusUpcoming`, `StatusDone` or `StatusCanceled`
	Gender           string    `json:"gender"`
	RegistrationOpen bool      `json:"registrationOpen" db:"registration_open"`
}

// Tournament adds additional information to `TournamentInfo`.
type Tournament struct {
	scores.Tracked

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
