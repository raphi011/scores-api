package repo

import (
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

type PlayerRepository interface {
	Get(id int) (*volleynet.Player, error)
	Ladder(gender string) ([]*volleynet.Player, error)
	New(p *volleynet.Player) (*volleynet.Player, error)
	Update(p *volleynet.Player) error
}


type TeamRepository interface {
	ByTournament(tournamentID int) ([]*volleynet.TournamentTeam, error)
	Delete(t *volleynet.TournamentTeam) error
	New(t *volleynet.TournamentTeam) (*volleynet.TournamentTeam, error)
	NewBatch(teams []*volleynet.TournamentTeam) error
	Update(t *volleynet.TournamentTeam) error
	UpdateBatch(teams []*volleynet.TournamentTeam) error
}

type TournamentRepository interface {
	Filter(seasons []int, leagues []string, formats []string) (
		[]*volleynet.FullTournament, error)
	Get(tournamentID int) (*volleynet.FullTournament, error)
	New(t *volleynet.FullTournament) (*volleynet.FullTournament, error)
	NewBatch(t ...*volleynet.FullTournament) error
	Update(t *volleynet.FullTournament) error
}

type UserRepository interface {
	All() ([]*scores.User, error)
	ByEmail(email string) (*scores.User, error)
	ByID(userID int) (*scores.User, error)
	New(user *scores.User) (*scores.User, error)
	Update(user *scores.User) error
}

type Repositories struct {
	PlayerRepo PlayerRepository
	TeamRepo TeamRepository
	TournamentRepo TournamentRepository
	UserRepo UserRepository
}