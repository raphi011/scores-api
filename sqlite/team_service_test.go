package sqlite

import (
	"testing"

	"github.com/raphi011/scores"
)

func TestTeamPlayerOrder(t *testing.T) {
	var player1ID, player2ID uint = 1, 2

	new1ID, new2ID := TeamPlayerOrder(player1ID, player2ID)

	if player1ID != new1ID || player2ID != new2ID {
		t.Errorf("TeamPlayerOrder(%[1]d, %[2]d), want %[1]d, %[2]d, got %[3]d, %[4]d",
			player1ID, player2ID, new1ID, new2ID)
	}

	player1ID, player2ID = 2, 1
	new1ID, new2ID = TeamPlayerOrder(player1ID, player2ID)

	if player1ID != new2ID || player2ID != new1ID {
		t.Errorf("TeamPlayerOrder(%[1]d, %[2]d), want %[2]d, %[1]d, got %[3]d, %[4]d",
			player1ID, player2ID, new1ID, new2ID)
	}
}

func TestCreate(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	playerService := &PlayerService{DB: db}
	player1, _ := playerService.Create(&scores.Player{Name: "Player1"})
	player2, _ := playerService.Create(&scores.Player{Name: "Player2"})
	teamService := &TeamService{DB: db}

	_, err := teamService.Create(&scores.Team{
		Name:      "Team1",
		Player1ID: player1.ID,
		Player2ID: player2.ID,
	})

	if err != nil {
		t.Error("Can't create team")
	}

	newTeam, err := teamService.ByPlayers(player1.ID, player2.ID)

	if err != nil {
		t.Errorf("TeamService.ByPlayers() err: %s", err)
	}
	if newTeam.Player1ID != player1.ID || newTeam.Player2ID != player2.ID {
		t.Error("TeamService.Create(), team not created")
	}
}
