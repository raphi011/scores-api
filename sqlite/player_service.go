package sqlite

import (
	"database/sql"
	"errors"
	scores "scores-backend"
)

var _ scores.PlayerService = &PlayerService{}

type PlayerService struct {
	DB *sql.DB
}

func (s *PlayerService) Players() (scores.Players, error) {
	players := []scores.Player{}

	rows, err := s.DB.Query(`
		SELECT
			id,
			name,
			user_id
		FROM players
		WHERE deleted_at is null
	`)

	var userID sql.NullInt64

	for rows.Next() {
		p := scores.Player{}
		err = rows.Scan(&p.ID, &p.Name, &userID)
		if err != nil {
			return nil, err
		}
		if userID.Valid {
			p.UserId = uint(userID.Int64)
		}
		players = append(players, p)
	}

	return players, nil
}

func (s *PlayerService) Player(ID uint) (*scores.Player, error) {
	player := &scores.Player{}

	var userID sql.NullInt64

	err := s.DB.QueryRow(`
		SELECT 
			id,
			name,
			user_id
		FROM players
		WHERE id = $1 and deleted_at is null
	`, ID).Scan(&player.ID, &player.Name, &userID)

	if err != nil {
		return nil, err
	}
	if userID.Valid {
		player.UserId = uint(userID.Int64)
	}

	return player, nil
}

func (s *PlayerService) Create(player *scores.Player) (*scores.Player, error) {
	result, err := s.DB.Exec("INSERT INTO players (created_at, name) VALUES (CURRENT_TIMESTAMP, $1)", player.Name)

	if err != nil {
		return nil, err
	}

	ID, _ := result.LastInsertId()

	player.ID = uint(ID)

	return player, nil
}

func (s *PlayerService) Delete(ID uint) error {
	result, err := s.DB.Exec("UPDATE players SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1", ID)

	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New("Player not found")
	}
	return nil
}
