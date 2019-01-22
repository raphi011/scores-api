package repo

import (
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

// PlayerRepository exposes CRUD operations on players.
type PlayerRepository interface {
	Get(id int) (*volleynet.Player, error)
	Ladder(gender string) ([]*volleynet.Player, error)
	New(p *volleynet.Player) (*volleynet.Player, error)
	// NewBatch(p ...*volleynet.Player) error
	Update(p *volleynet.Player) error
	// UpdateBatch(p ...*volleynet.Player) error
}

// TeamRepository exposes CRUD operations on teams.
type TeamRepository interface {
	ByTournament(tournamentID int) ([]*volleynet.TournamentTeam, error)
	Delete(t *volleynet.TournamentTeam) error
	New(t *volleynet.TournamentTeam) (*volleynet.TournamentTeam, error)
	NewBatch(t ...*volleynet.TournamentTeam) error
	Update(t *volleynet.TournamentTeam) error
	UpdateBatch(t ...*volleynet.TournamentTeam) error
}

// TournamentRepository exposes CRUD operations on tournaments.
type TournamentRepository interface {
	Filter(seasons []int, leagues []string, gender []string) (
		[]*volleynet.Tournament, error)
	Get(tournamentID int) (*volleynet.Tournament, error)
	New(t *volleynet.Tournament) (*volleynet.Tournament, error)
	NewBatch(t ...*volleynet.Tournament) error
	Update(t *volleynet.Tournament) error
	UpdateBatch(t ...*volleynet.Tournament) error

	Leagues() ([]scores.NameValue, error)
	SubLeagues() ([]scores.NameValue, error)
	Seasons() ([]int, error)
}

// UserRepository exposes CRUD operations on users.
type UserRepository interface {
	All() ([]*scores.User, error)
	ByEmail(email string) (*scores.User, error)
	ByID(userID int) (*scores.User, error)
	New(user *scores.User) (*scores.User, error)
	Update(user *scores.User) error
}

// Repositories is a collection of instances of all available repositories.
type Repositories struct {
	PlayerRepo     PlayerRepository
	TeamRepo       TeamRepository
	TournamentRepo TournamentRepository
	UserRepo       UserRepository
}
