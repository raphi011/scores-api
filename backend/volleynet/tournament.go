package volleynet

import (
	"time"

	"github.com/raphi011/scores"
)

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
	Season           string    `json:"season"`
	Start            time.Time `json:"start" db:"start_date"`
	End              time.Time `json:"end" db:"end_date"`
	Name             string    `json:"name" fako:"city"`
	League           string    `json:"league" db:"league"`
	LeagueKey        string    `json:"leagueKey" db:"league_key"`
	SubLeague        string    `json:"subLeague" db:"sub_league"`
	SubLeagueKey     string    `json:"phaseKey" db:"sub_league_key"`
	Link             string    `json:"link" fako:"domain_name"`
	EntryLink        string    `json:"entryLink" db:"entry_link"`
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
