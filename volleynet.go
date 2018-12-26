package scores

import (
	"time"

	"github.com/raphi011/scores/volleynet"
)

// VolleynetRepository stores volleynet data
type VolleynetRepository interface {
	Tournament(tournamentID int) (*volleynet.FullTournament, error)

	AllTournaments() ([]*volleynet.FullTournament, error)
	SeasonTournaments(season int) ([]*volleynet.FullTournament, error)
	GetTournaments(gender, league string, season int) ([]*volleynet.FullTournament, error)
	TournamentsUpdatedSince(since time.Time) ([]*volleynet.FullTournament, error)

	NewTournament(t *volleynet.FullTournament) error
	UpdateTournament(t *volleynet.FullTournament) error

	UpdateTournamentTeam(t *volleynet.TournamentTeam) error
	UpdateTournamentTeams(teams []volleynet.TournamentTeam) error
	NewTeam(t *volleynet.TournamentTeam) error
	NewTeams(teams []volleynet.TournamentTeam) error
	DeleteTeam(t *volleynet.TournamentTeam) error
	TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error)

	SearchPlayers() ([]volleynet.Player, error)
	AllPlayers() ([]volleynet.Player, error)
	NewPlayer(p *volleynet.Player) error
	Player(id int) (*volleynet.Player, error)
	UpdatePlayer(p *volleynet.Player) error
	Ladder(gender string) ([]volleynet.Player, error)
}
