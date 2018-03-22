package sqlite

import (
	"database/sql"

	"github.com/raphi011/scores"
)

var _ scores.GroupService = &GroupService{}

func scanGroup(scanner scan) (*scores.Group, error) {
	g := &scores.Group{}

	err := scanner.Scan(
		&g.ID,
		&g.CreatedAt,
		&g.Name,
		&g.ImageURL,
	)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func scanGroups(db *sql.DB, query string, args ...interface{}) (scores.Groups, error) {
	groups := scores.Groups{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		g, err := scanGroup(rows)
		if err != nil {
			return nil, err
		}

		groups = append(groups, *g)
	}

	return groups, nil
}

const (
	groupsBaseSelectSQL = `
	SELECT
		g.id,
		g.created_at,	
		g.name,
		COALESCE(g.image_url, "") as image_url
	FROM groups g
	`

	groupsWhereSQL = " WHERE g.deleted_at is null "

	groupsByPlayerSelectSQL = groupsBaseSelectSQL + `
		JOIN groupPlayers gp on g.id = gp.group_id ` +
		groupsWhereSQL +
		" AND gp.player_id = $1 "

	groupsSelectSQL = groupsBaseSelectSQL + groupsWhereSQL
	groupSelectSQL  = groupsSelectSQL + " AND g.id = $1"
)

type GroupService struct {
	DB *sql.DB
}

func (s *GroupService) GroupsByPlayer(playerID uint) (scores.Groups, error) {
	return scanGroups(s.DB, groupsByPlayerSelectSQL, playerID)
}

func (s *GroupService) Groups() (scores.Groups, error) {
	return scanGroups(s.DB, groupsSelectSQL)
}

func (s *GroupService) Group(groupID uint) (*scores.Group, error) {
	p := &scores.Group{}

	row := s.DB.QueryRow(groupSelectSQL, groupID)

	p, err := scanGroup(row)

	if err != nil {
		return nil, err
	}

	return p, nil
}

const (
	groupsInsertSQL = `
		INSERT INTO groups
		(
			created_at,
			name,
			image_url
		)
		VALUES
		(
			CURRENT_TIMESTAMP,
			$1,
			$2
		)
	`
)

func (s *GroupService) Create(group *scores.Group) (*scores.Group, error) {
	result, err := s.DB.Exec(groupsInsertSQL,
		group.Name,
		group.ImageURL,
	)

	if err != nil {
		return nil, err
	}

	ID, _ := result.LastInsertId()

	group.ID = uint(ID)

	return group, nil
}

const (
	addPlayerToGroupSQL = `
		INSERT INTO groupPlayers (player_id, group_id, role)
		VALUES ($1, $2, $3)
	`
)

func (s *GroupService) AddPlayerToGroup(playerID, groupID uint, role string) error {
	_, err := s.DB.Exec(addPlayerToGroupSQL, playerID, groupID, role)

	return err
}
