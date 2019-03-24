package repo

import (
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

// PlayerFilter exposes search fields for a player.
type PlayerFilter struct {
	FirstName string
	LastName  string
	Gender    string
}

// PlayerRepository exposes CRUD operations on players.
type PlayerRepository interface {
	Get(id int) (*volleynet.Player, error)
	New(p *volleynet.Player) (*volleynet.Player, error)
	Update(p *volleynet.Player) error
	Ladder(gender string) ([]*volleynet.Player, error)
	PreviousPartners(playerID int) ([]*volleynet.Player, error)
	Search(filter PlayerFilter) ([]*volleynet.Player, error)
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

// TournamentFilter contains all available Tournament filters.
type TournamentFilter struct {
	Seasons []string
	Leagues []string
	Genders []string
}

// TournamentRepository exposes CRUD operations on tournaments.
type TournamentRepository interface {
	Search(filter TournamentFilter) (
		[]*volleynet.Tournament, error)
	Get(tournamentID int) (*volleynet.Tournament, error)
	New(t *volleynet.Tournament) (*volleynet.Tournament, error)
	NewBatch(t ...*volleynet.Tournament) error
	Update(t *volleynet.Tournament) error
	UpdateBatch(t ...*volleynet.Tournament) error

	Leagues() ([]string, error)
	SubLeagues() ([]string, error)
	Seasons() ([]string, error)
}

// UserRepository exposes CRUD operations on users.
type UserRepository interface {
	All() ([]*scores.User, error)
	ByEmail(email string) (*scores.User, error)
	ByID(userID int) (*scores.User, error)
	New(user *scores.User) (*scores.User, error)
	Update(user *scores.User) error
}

// SettingRepository exposes CRUD operations on settings.
type SettingRepository interface {
	Create(setting *scores.Setting) (*scores.Setting, error)
	Update(setting *scores.Setting) error
	ByUserID(userID int) ([]*scores.Setting, error)
}

// Repositories is a collection of instances of all available repositories.
type Repositories struct {
	PlayerRepo     PlayerRepository
	TeamRepo       TeamRepository
	TournamentRepo TournamentRepository
	UserRepo       UserRepository
	SettingRepo    SettingRepository
}
