package sqlite

import (
	"testing"

	"github.com/raphi011/scores"
)

func TestGetPlayersByGroup(t *testing.T) {
	s := createRepositories(t)

	g, _ := s.Group.Create(&scores.Group{Name: "Testgroup"})

	p1, _ := s.Player.Create(&scores.Player{Name: "Test1"})
	p2, _ := s.Player.Create(&scores.Player{Name: "Test2"})
	s.Player.Create(&scores.Player{Name: "Test3"})

	s.Group.AddPlayer(p1.ID, g.ID, "user")
	s.Group.AddPlayer(p2.ID, g.ID, "user")

	players, err := s.Player.ByGroup(g.ID)

	if err != nil {
		t.Errorf("PlayerRepository.ByGroup() err: %s", err)
	} else if len(players) != 2 {
		t.Errorf("PlayerRepository.ByGroup(), want 2 players, got %d ", len(players))
	}
}

func TestGetPlayers(t *testing.T) {
	s := createRepositories(t)

	s.Player.Create(&scores.Player{Name: "Test1"})
	s.Player.Create(&scores.Player{Name: "Test2"})
	players, err := s.Player.All()

	if err != nil {
		t.Errorf("PlayerRepository.All() err: %s", err)
	} else if len(players) != 2 {
		t.Errorf("PlayerRepository.All(), want 2 players, got %d ", len(players))
	}
}

func TestCreatePlayer(t *testing.T) {
	s := createRepositories(t)

	player, err := s.Player.Create(&scores.Player{Name: "Test"})

	if err != nil {
		t.Error("Can't create player")
	}
	if player.ID == 0 {
		t.Error("PlayerID not assigned")
	}

	playerID := player.ID

	player, err = s.Player.Get(playerID)

	if err != nil {
		t.Errorf("PlayerRepository.Get() err: %s", err)
	}
	if player.ID != playerID {
		t.Errorf("PlayerRepository.Get(), want ID %d, got %d", playerID, player.ID)
	}
}

func TestDeletePlayer(t *testing.T) {
	s := createRepositories(t)

	player, _ := s.Player.Create(&scores.Player{Name: "Test"})
	s.Player.Create(&scores.Player{Name: "Test2"})

	err := s.Player.Delete(player.ID)

	if err != nil {
		t.Errorf("PlayerRepository.Delete() err: %s", err)
	}

	players, _ := s.Player.All()
	playerCount := len(players)

	if playerCount != 1 {
		t.Errorf("len(PlayerRepository.All()), want 1, got %d", playerCount)
	}

}
