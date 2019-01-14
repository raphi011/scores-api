package volleynet

import (
	"time"

	"github.com/raphi011/scores"
)

// PlayerInfo contains all player information that the search player api returns
type PlayerInfo struct {
	scores.TrackedModel
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birthday  time.Time `json:"birthday"`
}

// Player adds additional information to `PlayerInfo`
type Player struct {
	PlayerInfo

	Gender       string `json:"gender"`
	TotalPoints  int    `json:"totalPoints"`
	Rank         int    `json:"rank"`
	Club         string `json:"club"`
	CountryUnion string `json:"countryUnion"`
	License      string `json:"license"`
}

// TournamentTeam is the current status of the team entry in a
// tournament, if the tournament is finished it may also contain
// the seed
type TournamentTeam struct {
	TournamentID int     `json:"tournamentId"`
	TotalPoints  int     `json:"totalPoints"`
	Seed         int     `json:"seed"`
	Rank         int     `json:"rank"`
	WonPoints    int     `json:"wonPoints"`
	Player1      *Player `json:"player1"`
	Player2      *Player `json:"player2"`
	PrizeMoney   float32 `json:"prizeMoney"`
	Deregistered bool    `json:"deregistered"`
}

const (
	// StatusUpcoming represents the state of a tournament not done yet
	StatusUpcoming = "upcoming"
	// StatusDone represents the state of a completed tournament
	StatusDone = "done"
	// StatusCanceled represents the state of a canceled tournament
	StatusCanceled = "canceled"
)

// Tournament is all the information that can be parsed from the tournament list
type Tournament struct {
	Start            time.Time `json:"start"`
	End              time.Time `json:"end"`
	Name             string    `json:"name" fako:"city"`
	Season           int       `json:"season"`
	League           string    `json:"league"`
	Phase            string    `json:"phase"`
	Link             string    `json:"link" fako:"domain_name"` 
	EntryLink        string    `json:"entryLink"`
	ID               int       `json:"id"`
	Status           string    `json:"status"` // can be `StatusUpcoming`, `StatusDone` or `StatusCanceled`
	RegistrationOpen bool      `json:"registrationOpen"`
	Format 			 string    `json:"format"`
}

// FullTournament adds additional information to `Tournament`
type FullTournament struct {
	Tournament
	CreatedAt       time.Time        `json:"createdAt"`
	UpdatedAt       time.Time        `json:"updatedAt"`
	EndRegistration time.Time        `json:"endRegistration"`
	Teams           []TournamentTeam `json:"teams"`
	Location        string           `json:"location"`
	HTMLNotes       string           `json:"htmlNotes" fako:"paragraph"`
	Mode            string           `json:"mode"`
	Organiser       string           `json:"organiser" fako:"full_name"`
	Phone           string           `json:"phone" fako:"phone"`
	Email           string           `json:"email" fako:"email_address"`
	Website         string           `json:"website"`
	CurrentPoints   string           `json:"currentPoints"`
	LivescoringLink string           `json:"livescoringLink"`
	SignedupTeams   int              `json:"signedupTeams"`
	MaxTeams        int              `json:"maxTeams"`
	MinTeams        int              `json:"minTeams"`
	MaxPoints       int              `json:"maxPoints"`
	Latitude        float32          `json:"latitude"`
	Longitude       float32          `json:"longitude"`
}

// LoginData contains the data of the form that is shown after a successful login
type LoginData struct {
	PlayerInfo
	License License
}

// License contains the license data of a player
type License struct {
	Nr        string
	Type      string
	Requested string
}
