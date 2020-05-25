package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/raphi011/scores-backend/repo"
	"github.com/raphi011/scores-backend/repo/sql/crud"
	"github.com/raphi011/scores-backend/volleynet"
)

type playerRepository struct {
	DB *sqlx.DB
}

var _ repo.PlayerRepository = &playerRepository{}

// ByGender gets all players of the passed gender.
func (s *playerRepository) ByGender(gender string) ([]*volleynet.Player, error) {
	players := []*volleynet.Player{}
	err := crud.Read(s.DB, "player/select-by-gender", &players, gender)

	return players, errors.Wrap(err, "by gender")
}

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

	err := crud.ReadNamed(s.DB, "player/select-partners", &players,
		map[string]interface{}{"player_id": playerID})

	return players, errors.Wrap(err, "previousPartners")
}

// Search searches for players that satisfy the passed filter.
func (s *playerRepository) Search(filter repo.PlayerFilter) ([]*volleynet.Player, error) {
	players := []*volleynet.Player{}

	filter.FirstName = startsWith(filter.FirstName)
	filter.LastName = startsWith(filter.LastName)

	err := crud.ReadNamed(s.DB, "player/search", &players, filter)

	return players, errors.Wrap(err, "search")
}

func startsWith(query string) string {
	if len(query) == 0 {
		return query
	}

	return query + "%"
}
