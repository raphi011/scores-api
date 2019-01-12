package sql

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

// PlayerRepository implements the PlayerRepository interface.
type PlayerRepository struct {
	DB *sqlx.DB
}

// Ladder gets all players of the passed gender that have a rank.
func (s *PlayerRepository) Ladder(gender string) ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, query("player/select-ladder"), gender)
}

// All loads all players.
// Note: should only be used for debugging.
func (s *PlayerRepository) All() ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, query("player/select-all"))
}

// Get loads a player.
func (s *PlayerRepository) Get(id int) (*volleynet.Player, error) {
	row := s.DB.QueryRow(
		query("player/select-by-id"),
		id,
	)

	return scanVolleynetPlayer(row)
}

// New creates a new player.
func (s *PlayerRepository) New(p *volleynet.Player) error {
	_, err := s.DB.Exec(query("player/insert"),
		p.ID,
		p.FirstName,
		p.LastName,
		p.Birthday,
		p.Gender,
		p.TotalPoints,
		p.Rank,
		p.Club,
		p.CountryUnion,
		p.License,
	)

	return err
}

// Update updates a player.
func (s *PlayerRepository) Update(p *volleynet.Player) error {
	result, err := s.DB.Exec(
		query("player/update"),
		p.FirstName,
		p.LastName,
		p.Birthday,
		p.Gender,
		p.TotalPoints,
		p.Rank,
		p.Club,
		p.CountryUnion,
		p.License,
		p.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Player not found")
	}

	return nil
}

func scanVolleynetPlayers(db *sqlx.DB, query string, args ...interface{}) ([]volleynet.Player, error) {
	players := []volleynet.Player{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		player, err := scanVolleynetPlayer(rows)

		if err != nil {
			return nil, err
		}

		players = append(players, *player)
	}

	return players, nil
}

func scanVolleynetPlayer(scanner scan) (*volleynet.Player, error) {
	p := &volleynet.Player{}

	err := scanner.Scan(
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
		return nil, err
	}

	return p, nil
}