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
	s := createRepositories(t)

	player1, _ := s.playerRepository.Create(&scores.Player{Name: "Player1"})
	player2, _ := s.playerRepository.Create(&scores.Player{Name: "Player2"})

	_, err := s.teamRepository.Create(&scores.Team{
		Name:      "Team1",
		Player1ID: player1.ID,
		Player2ID: player2.ID,
	})

	if err != nil {
		t.Error("Can't create team")
	}

	newTeam, err := s.teamRepository.ByPlayers(player1.ID, player2.ID)

	if err != nil {
		t.Errorf("TeamRepository.ByPlayers() err: %s", err)
	}
	if newTeam.Player1ID != player1.ID || newTeam.Player2ID != player2.ID {
		t.Error("TeamRepository.Create(), team not created")
	}
}
