package sqlite

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/raphi011/scores"
)

var _ scores.PlayerRepository = &PlayerRepository{}

type PlayerRepository struct {
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
	playersGroupJoinSQL = " JOIN group_players gp on gp.player_id = p.id"
	playersByGroupSQL   = playersBaseSelectSQL + playersGroupJoinSQL + playersWhereSQL + `
	 AND gp.group_id = ?	
	`

	playerSelectSQL = playersBaseSelectSQL + playersWhereSQL + " and p.id = ?"
)

func (s *PlayerRepository) ByGroup(groupID uint) ([]scores.Player, error) {
	return scanPlayers(s.DB, query("player/select-by-group"), groupID)
}

func (s *PlayerRepository) All() ([]scores.Player, error) {
	return scanPlayers(s.DB, query("player/select-all"))
}

// Get retrieves a player by an ID
func (s *PlayerRepository) Get(ID uint) (*scores.Player, error) {
	groupRepository := GroupRepository{DB: s.DB}

	p := &scores.Player{}

	row := s.DB.QueryRow(query("player/select-by-id"), ID)

	p, err := scanPlayer(row)

	if err != nil {
		return nil, err
	}

	groups, err := groupRepository.ByPlayer(p.ID)

	if err != nil {
		return nil, err
	}

	p.Groups = groups

	return p, nil
}

func (s *PlayerRepository) Create(player *scores.Player) (*scores.Player, error) {
	var userID sql.NullInt64

	if player.UserID > 0 {
		userID.Scan(player.UserID)
	}

	result, err := s.DB.Exec(query("player/insert"), player.Name, userID)

	if err != nil {
		return nil, errors.Wrap(err, "creating player failed")
	}

	ID, _ := result.LastInsertId()

	player.ID = uint(ID)

	return player, nil
}

func (s *PlayerRepository) Delete(ID uint) error {
	result, err := s.DB.Exec(query("player/update-delete"), ID)

	if err != nil {
		return errors.Wrap(err, "deleting player failed")
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New("Player not found")
	}
	return nil
}

func scanPlayer(scanner scan) (*scores.Player, error) {
	var userID sql.NullInt64
	var profileImageURL sql.NullString

	p := &scores.Player{}
	err := scanner.Scan(&p.ID, &p.Name, &userID, &profileImageURL, &p.CreatedAt)

	if err != nil {
		return nil, errors.Wrap(err, "scanning player failed")
	}

	if profileImageURL.Valid {
		p.ProfileImageURL = profileImageURL.String
	}
	if userID.Valid {
		p.UserID = uint(userID.Int64)
	}

	return p, nil
}

func scanPlayers(db *sql.DB, query string, args ...interface{}) ([]scores.Player, error) {
	players := []scores.Player{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "players query failed")
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
