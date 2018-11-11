package sqlite

import (
	"testing"

	"github.com/raphi011/scores"
)

func TestCreateGroup(t *testing.T) {
	s := createRepositories(t)

	groupName := "Test"
	g, err := s.Group.Create(&scores.Group{Name: groupName})

	if err != nil {
		t.Errorf("GroupRepository.Create() err: %s", err)
		return
	}

	gr, err := s.Group.Group(g.ID)

	if err != nil {
		t.Errorf("GroupRepository.Group() err: %s", err)
	} else if g.Name != gr.Name {
		t.Errorf("GroupRepository.Create() want g.Name = %s, got: %s", g.Name, gr.Name)
	}
}

func TestAddPlayerToGroup(t *testing.T) {
	s := createRepositories(t)

	p, _ := s.Player.Create(&scores.Player{Name: "Player"})
	g, _ := s.Group.Create(&scores.Group{Name: "asd"})
	err := s.Group.AddPlayerToGroup(p.ID, g.ID, "user")

	if err != nil {
		t.Errorf("GroupRepository.AddPlayerToGroup() err: %s", err)
	}
}
