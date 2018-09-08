package sqlite

import (
	"testing"

	"github.com/raphi011/scores"
)

func TestCreateGroup(t *testing.T) {
	s := createServices(t)

	groupName := "Test"
	g, err := s.groupService.Create(&scores.Group{Name: groupName})

	if err != nil {
		t.Errorf("GroupService.Create() err: %s", err)
		return
	}

	gr, err := s.groupService.Group(g.ID)

	if err != nil {
		t.Errorf("GroupService.Group() err: %s", err)
	} else if g.Name != gr.Name {
		t.Errorf("GroupService.Create() want g.Name = %s, got: %s", g.Name, gr.Name)
	}
}

func TestAddPlayerToGroup(t *testing.T) {
	s := createServices(t)

	p, _ := s.playerService.Create(&scores.Player{Name: "Player"})
	g, _ := s.groupService.Create(&scores.Group{Name: "asd"})
	err := s.groupService.AddPlayerToGroup(p.ID, g.ID, "user")

	if err != nil {
		t.Errorf("GroupService.AddPlayerToGroup() err: %s", err)
	}
}
