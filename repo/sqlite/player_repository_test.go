package sqlite

import (
	"testing"

	"github.com/raphi011/scores"
)

func TestGetPlayersByGroup(t *testing.T) {
	s := createRepositories(t)

	g, _ := s.groupRepository.Create(&scores.Group{Name: "Testgroup"})

	p1, _ := s.playerRepository.Create(&scores.Player{Name: "Test1"})
	p2, _ := s.playerRepository.Create(&scores.Player{Name: "Test2"})
	s.playerRepository.Create(&scores.Player{Name: "Test3"})

	s.groupRepository.AddPlayerToGroup(p1.ID, g.ID, "user")
	s.groupRepository.AddPlayerToGroup(p2.ID, g.ID, "user")

	players, err := s.playerRepository.ByGroup(g.ID)

	if err != nil {
		t.Errorf("PlayerRepository.ByGroup() err: %s", err)
	} else if len(players) != 2 {
		t.Errorf("PlayerRepository.ByGroup(), want 2 players, got %d ", len(players))
	}
}

func TestGetPlayers(t *testing.T) {
	s := createRepositories(t)

	s.playerRepository.Create(&scores.Player{Name: "Test1"})
	s.playerRepository.Create(&scores.Player{Name: "Test2"})
	players, err := s.playerRepository.Players()

	if err != nil {
		t.Errorf("PlayerRepository.Players() err: %s", err)
	} else if len(players) != 2 {
		t.Errorf("PlayerRepository.Players(), want 2 players, got %d ", len(players))
	}
}

func TestCreatePlayer(t *testing.T) {
	s := createRepositories(t)

	player, err := s.playerRepository.Create(&scores.Player{Name: "Test"})

	if err != nil {
		t.Error("Can't create player")
	}
	if player.ID == 0 {
		t.Error("PlayerID not assigned")
	}

	playerID := player.ID

	player, err = s.playerRepository.Player(playerID)

	if err != nil {
		t.Errorf("PlayerRepository.Player() err: %s", err)
	}
	if player.ID != playerID {
		t.Errorf("PlayerRepository.Player(), want ID %d, got %d", playerID, player.ID)
	}
}

func TestDeletePlayer(t *testing.T) {
	s := createRepositories(t)

	player, _ := s.playerRepository.Create(&scores.Player{Name: "Test"})
	s.playerRepository.Create(&scores.Player{Name: "Test2"})

	err := s.playerRepository.Delete(player.ID)

	if err != nil {
		t.Errorf("PlayerRepository.Delete() err: %s", err)
	}

	players, _ := s.playerRepository.Players()
	playerCount := len(players)

	if playerCount != 1 {
		t.Errorf("len(PlayerRepository.Players()), want 1, got %d", playerCount)
	}

}
