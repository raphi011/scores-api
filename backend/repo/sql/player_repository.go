package sql

import (
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/volleynet"
)

// PlayerRepository implements the PlayerRepository interface.
type PlayerRepository struct {
	DB *sqlx.DB
}

var _ repo.PlayerRepository = &PlayerRepository{}

// Ladder gets all players of the passed gender that have a rank.
func (s *PlayerRepository) Ladder(gender string) ([]*volleynet.Player, error) {
	return s.scan("player/select-ladder", gender)
}

// Get loads a player.
func (s *PlayerRepository) Get(id int) (*volleynet.Player, error) {
	player, err := s.scanOne("player/select-by-id", id)

	return player, errors.Wrap(err, "get player")
}

// New creates a new player.
func (s *PlayerRepository) New(p *volleynet.Player) (*volleynet.Player, error) {
	_, err := exec(s.DB, "player/insert", p)

	return p, errors.Wrap(err, "new player")
}

// Update updates a player.
func (s *PlayerRepository) Update(p *volleynet.Player) error {
	err := update(s.DB, "player/update", p)

	return errors.Wrap(err, "update player")
}

func (s *PlayerRepository) scan(queryName string, args ...interface{}) (
	[]*volleynet.Player, error) {

	players := []*volleynet.Player{}

	q := query(s.DB, queryName)

	rows, err := s.DB.Query(q, args...)

	if err != nil {
		return players, err
	}

	defer rows.Close()

	for rows.Next() {
		p := &volleynet.Player{}

		err := rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.Birthday,
			&p.Gender,
			&p.TotalPoints,
			&p.Rank,
			&p.Club,
			&p.CountryUnion,
			&p.License,
		)

		if err != nil {
			return players, err
		}

		players = append(players, p)
	}

	return players, nil
}

func (s *PlayerRepository) scanOne(query string, args ...interface{}) (
	*volleynet.Player, error) {

	players, err := s.scan(query, args...)

	if err != nil {
		return nil, err
	}

	if len(players) >= 1 {
		return players[0], nil
	}

	return nil, scores.ErrNotFound
}
