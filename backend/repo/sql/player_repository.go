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
