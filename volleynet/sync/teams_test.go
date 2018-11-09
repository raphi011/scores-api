package sync

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores/volleynet"
)

func TestDistinctPlayers(t *testing.T) {
	input := []volleynet.TournamentTeam{
		volleynet.TournamentTeam{
			Player1: &volleynet.Player{
				PlayerInfo: volleynet.PlayerInfo{ID: 1},
			},
			Player2: &volleynet.Player{
				PlayerInfo: volleynet.PlayerInfo{ID: 2},
			},
		},
		volleynet.TournamentTeam{
			Player1: &volleynet.Player{
				PlayerInfo: volleynet.PlayerInfo{ID: 3},
			},
			Player2: &volleynet.Player{
				PlayerInfo: volleynet.PlayerInfo{ID: 1},
			},
		},
	}

	expected := []*volleynet.Player{
		input[0].Player1,
		input[0].Player2,
		input[1].Player1,
	}

	output := distinctPlayers(input)

	if !cmp.Equal(output, expected) {
		t.Errorf("distinctPlayers() diff:\n%s", cmp.Diff(output, expected))
	}
}

func TestAddMissingPlayers(t *testing.T) {
	_, volleynetMock, syncRepository := syncMock()

	teams := []volleynet.TournamentTeam{
		volleynet.TournamentTeam{
			Player1: &volleynet.Player{
				PlayerInfo: volleynet.PlayerInfo{ID: 1},
			},
			Player2: &volleynet.Player{
				PlayerInfo: volleynet.PlayerInfo{ID: 2},
			},
		},
	}

	var nilPlayer *volleynet.Player

	volleynetMock.On("Player", 1).Return(nilPlayer, nil)
	volleynetMock.On("Player", 2).Return(teams[0].Player2, nil)

	volleynetMock.On("NewPlayer", teams[0].Player1).Return(nil)

	err := syncRepository.addMissingPlayers(teams)

	if err != nil {
		t.Fatal(err)
	}

	volleynetMock.AssertExpectations(t)
}

func TestSyncTournamentTeams(t *testing.T) {
	_, _, syncRepository := syncMock()

	changes := &TeamChanges{}
	tournamentID := 1

	teamDeleted := volleynet.TournamentTeam{
		TournamentID: tournamentID,
		Seed:         1,
		Player1: &volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{ID: 1},
		},
		Player2: &volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{ID: 2},
		},
	}

	teamOutdated := volleynet.TournamentTeam{
		TournamentID: tournamentID,
		Seed:         2,
		Player1: &volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{ID: 5},
		},
		Player2: &volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{ID: 6},
		},
	}

	teamNew := volleynet.TournamentTeam{
		TournamentID: tournamentID,
		Seed:         2,
		Player1: &volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{ID: 3},
		},
		Player2: &volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{ID: 4},
		},
	}

	teamUpdated := teamOutdated
	teamUpdated.Seed = 3

	old := []volleynet.TournamentTeam{
		teamDeleted,
		teamOutdated,
	}

	new := []volleynet.TournamentTeam{
		teamUpdated,
		teamNew,
	}

	syncRepository.syncTournamentTeams(changes, old, new)

	if len(changes.New) != 1 {
		t.Errorf("SyncRepository.syncTournamentTeam(...) want: len(changes.New) == 1, got: %d", len(changes.New))
	}
	if len(changes.Update) != 1 {
		t.Errorf("SyncRepository.syncTournamentTeam(...) want: len(changes.Update) == 1, got: %d", len(changes.Update))
	}
	if len(changes.Delete) != 1 {
		t.Errorf("SyncRepository.syncTournamentTeam(...) want: len(changes.Delete) == 1, got: %d", len(changes.Delete))
	}
}
