package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sql/crud"
	"github.com/raphi011/scores/volleynet"
)

type playerRepository struct {
	DB *sqlx.DB
}

var _ repo.PlayerRepository = &playerRepository{}

// Ladder gets all players of the passed gender that have a rank.
func (s *playerRepository) Ladder(gender string) ([]*volleynet.Player, error) {

	players := []*volleynet.Player{}
	err := crud.Read(s.DB, "player/select-ladder", &players, gender)

	return players, errors.Wrap(err, "ladder")
}

// Get loads a player.
func (s *playerRepository) Get(id int) (*volleynet.Player, error) {

	player := &volleynet.Player{}
	err := crud.ReadOne(s.DB, "player/select-by-id", player, id)

	return player, errors.Wrap(err, "get player")
}

// New creates a new player.
func (s *playerRepository) New(p *volleynet.Player) (*volleynet.Player, error) {
	err := crud.Create(s.DB, "player/insert", p)

	return p, errors.Wrap(err, "new player")
}

// Update updates a player.
func (s *playerRepository) Update(p *volleynet.Player) error {
	err := crud.Update(s.DB, "player/update", p)

	return errors.Wrap(err, "update player")
}

// PreviousPartners returns a list of all partners a player has played with before.
func (s *playerRepository) PreviousPartners(playerID int) ([]*volleynet.Player, error) {
	players := []*volleynet.Player{}

	err := crud.Read(s.DB, "player/select-partners", &players, playerID)

	return players, errors.Wrap(err, "previousPartners")
}

// Search searches for players that satisfy the passed filter.
func (s *playerRepository) Search(filter repo.PlayerFilter) ([]*volleynet.Player, error) {
	players := []*volleynet.Player{}

	err := crud.Read(s.DB, "player/search", &players,
		startsWith(filter.FirstName),
		startsWith(filter.LastName),
		filter.Gender,
)

	return players, errors.Wrap(err, "search")
}

func startsWith(query string) string {
	if len(query) == 0 {
		return query
	}

	return query + "%"
}