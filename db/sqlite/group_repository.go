package sqlite

import (
	"database/sql"

	"github.com/raphi011/scores"
)

var _ scores.GroupRepository = &GroupRepository{}

type GroupRepository struct {
	DB *sql.DB
}

func (s *GroupRepository) GroupsByPlayer(playerID uint) (scores.Groups, error) {
	return scanGroups(s.DB, query("group/select-by-player"), playerID)
}

func (s *GroupRepository) Groups() (scores.Groups, error) {
	return scanGroups(s.DB, query("group/select-all-groups"))
}

func (s *GroupRepository) Group(groupID uint) (*scores.Group, error) {
	return scanGroup(s.DB.QueryRow(query("group/select-by-id"), groupID))
}

func (s *GroupRepository) Create(group *scores.Group) (*scores.Group, error) {
	result, err := s.DB.Exec(query("group/insert"),
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

func (s *GroupRepository) AddPlayerToGroup(playerID, groupID uint, role string) error {
	_, err := s.DB.Exec(query("group/insert-group-players"), playerID, groupID, role)

	return err
}

func scanGroup(scanner scan) (*scores.Group, error) {
	g := &scores.Group{}

	err := scanner.Scan(
		&g.ID,
		&g.CreatedAt,
		&g.Name,
		&g.ImageURL,
	)

	return g, err
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
