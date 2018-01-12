package sqlite

import (
	"scores-backend"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	playerService := &PlayerService{DB: db}
	playerService.Create(&scores.Player{Name: "Test1"})
	playerService.Create(&scores.Player{Name: "Test2"})
	players, err := playerService.Players()

	if err != nil {
		t.Errorf("PlayerService.Players() err: %s", err)
	} else if len(players) != 2 {
		t.Errorf("PlayerService.Players(), want 2 players, got %d ", len(players))
	}
}

func TestCreatePlayer(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	playerService := &PlayerService{DB: db}
	player, err := playerService.Create(&scores.Player{Name: "Test"})

	if err != nil {
		t.Error("Can't create player")
	}
	if player.ID == 0 {
		t.Error("PlayerID not assigned")
	}

	playerID := player.ID

	player, err = playerService.Player(playerID)

	if err != nil {
		t.Errorf("PlayerService.Player() err: %s", err)
	}
	if player.ID != playerID {
		t.Errorf("PlayerService.Player(), want ID %d, got %d", playerID, player.ID)
	}
}

func TestDeletePlayer(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	playerService := &PlayerService{DB: db}
	player, _ := playerService.Create(&scores.Player{Name: "Test"})
	playerService.Create(&scores.Player{Name: "Test2"})

	err := playerService.Delete(player.ID)

	if err != nil {
		t.Errorf("PlayerService.Delete() err: %s", err)
	}

	players, _ := playerService.Players()
	playerCount := len(players)

	if playerCount != 1 {
		t.Errorf("len(PlayerService.Players()), want 1, got %d", playerCount)
	}

}
