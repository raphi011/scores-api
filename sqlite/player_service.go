package sqlite

import (
	"database/sql"
	"errors"
	"log"

	"github.com/raphi011/scores"
)

var _ scores.PlayerService = &PlayerService{}

type PlayerService struct {
	DB *sql.DB
}

const (
	playersBaseSelectSQL = `
		SELECT
			p.id,
			p.name,
			p.user_id,
			u.profile_image_url,
			p.created_at
		FROM players p
		LEFT JOIN users u on u.id = p.user_id 
		`
	playersSelectSQL    = playersBaseSelectSQL + playersWhereSQL
	playersWhereSQL     = " WHERE p.deleted_at is null"
	playersGroupJoinSQL = " JOIN groupPlayers gp on gp.player_id = p.id"
	playersByGroupSQL   = playersBaseSelectSQL + playersGroupJoinSQL + playersWhereSQL + `
	 AND gp.group_id = $1	
	`

	playerSelectSQL = playersBaseSelectSQL + playersWhereSQL + " and p.id = $1"
)

func scanPlayer(scanner scan) (*scores.Player, error) {
	var userID sql.NullInt64
	var profileImageURL sql.NullString

	p := &scores.Player{}
	err := scanner.Scan(&p.ID, &p.Name, &userID, &profileImageURL, &p.CreatedAt)

	if err != nil {
		return nil, err
	}

	if profileImageURL.Valid {
		p.ProfileImageURL = profileImageURL.String
	}
	if userID.Valid {
		p.UserID = uint(userID.Int64)
	}

	return p, nil
}

func scanPlayers(db *sql.DB, query string, args ...interface{}) (scores.Players, error) {
	players := scores.Players{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		p, err := scanPlayer(rows)
		if err != nil {
			return nil, err
		}

		players = append(players, *p)
	}

	return players, nil
}

func (s *PlayerService) ByGroup(groupID uint) (scores.Players, error) {
	log.Print(playersByGroupSQL)
	return scanPlayers(s.DB, playersByGroupSQL, groupID)
}

func (s *PlayerService) Players() (scores.Players, error) {
	return scanPlayers(s.DB, playersSelectSQL)
}

func (s *PlayerService) Player(ID uint) (*scores.Player, error) {
	p := &scores.Player{}

	row := s.DB.QueryRow(playerSelectSQL, ID)

	p, err := scanPlayer(row)

	if err != nil {
		return nil, err
	}

	return p, nil
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
