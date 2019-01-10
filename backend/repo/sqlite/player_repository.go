package sqlite

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

// VolleynetPlayerRepository implements VolleynetPlayerRepository interface
type VolleynetPlayerRepository struct {
	DB *sql.DB
}


// Ladder gets all players of the passed gender that have a rank
func (s *VolleynetPlayerRepository) Ladder(gender string) ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, query("volleynet/select-player-ladder"), gender)
}

// All loads all players
// Note: should only be used for debugging
func (s *VolleynetPlayerRepository) All() ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, query("volleynet/select-player-all"))
}

// Get loads a player
func (s *VolleynetPlayerRepository) Get(id int) (*volleynet.Player, error) {
	row := s.DB.QueryRow(
		query("volleynet/select-player-by-id"),
		id,
	)

	return scanVolleynetPlayer(row)
}

// New creates a new player
func (s *VolleynetPlayerRepository) New(p *volleynet.Player) error {
	_, err := s.DB.Exec(query("volleynet/insert-player"),
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

// Update updates a player
func (s *VolleynetPlayerRepository) Update(p *volleynet.Player) error {
	result, err := s.DB.Exec(
		query("volleynet/update-player"),
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

func scanVolleynetPlayers(db *sql.DB, query string, args ...interface{}) ([]volleynet.Player, error) {
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