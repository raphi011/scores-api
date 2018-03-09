package sqlite

import (
	"database/sql"

	"github.com/raphi011/scores"
)

var _ scores.GroupService = &GroupService{}

type GroupService struct {
	DB *sql.DB
}

func (s *GroupService) GroupsByPlayer(playerID uint) (scores.Groups, error) {

	return nil, nil
}

func (s *GroupService) Groups() (scores.Groups, error) {

	return nil, nil
}

func (s *GroupService) Group(groupID uint) (*scores.Group, error) {

	return nil, nil
}

func (s *GroupService) Create(group *scores.Group) (*scores.Group, error) {

	return nil, nil
}

func (s *GroupService) AddPlayerToGroup(playerID, groupID uint) error {

	return nil
}
