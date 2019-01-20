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

// func (s *playerRepository) scan(queryName string, args ...interface{}) (
// 	[]*volleynet.Player, error) {

// 	players := []*volleynet.Player{}

// 	q := query(s.DB, queryName)

// 	rows, err := s.DB.Query(q, args...)

// 	if err != nil {
// 		return players, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		p := &volleynet.Player{}

// 		err := rows.Scan(
// 			&p.ID,
// 			&p.FirstName,
// 			&p.LastName,
// 			&p.Birthday,
// 			&p.Gender,
// 			&p.TotalPoints,
// 			&p.Rank,
// 			&p.Club,
// 			&p.CountryUnion,
// 			&p.License,
// 		)

// 		if err != nil {
// 			return players, err
// 		}

// 		players = append(players, p)
// 	}

// 	return players, nil
// }

// func (s *playerRepository) scanOne(query string, args ...interface{}) (
// 	*volleynet.Player, error) {

// 	players, err := s.scan(query, args...)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(players) >= 1 {
// 		return players[0], nil
// 	}

// 	return nil, scores.ErrNotFound
// }
