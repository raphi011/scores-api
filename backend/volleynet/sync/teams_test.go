package sync

import (
	"testing"

	"github.com/raphi011/scores/repo/sql"
	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/test"
)

func TestDistinctPlayers(t *testing.T) {
	input := []*volleynet.TournamentTeam{
		&volleynet.TournamentTeam{
			Player1: &volleynet.Player{ID: 1},
			Player2: &volleynet.Player{ID: 2},
		},
		&volleynet.TournamentTeam{
			Player1: &volleynet.Player{ID: 3},
			Player2: &volleynet.Player{ID: 1},
		},
	}

	expected := []*volleynet.Player{
		input[0].Player1,
		input[0].Player2,
		input[1].Player1,
	}

	output := distinctPlayers(input)

	test.Compare(t, "distinctPlayers() diff:\n%s", output,expected)
}

func TestAddMissingPlayers(t *testing.T) {
	_, service, db := syncMock(t)

	players := sql.CreatePlayers(t, db,
		sql.P{ ID: 1 },
		sql.P{ ID: 2 },
	)

	teams := []*volleynet.TournamentTeam{
		&volleynet.TournamentTeam{
			Player1: players[0],
			Player2: players[1],
		},
	}

	err := service.addMissingPlayers(teams)

	test.Check(t, "addMissingPlayers() failed: %v", err)
}

func TestSyncTournamentTeams(t *testing.T) {
	_, service, _ := syncMock(t)

	changes := &TeamChanges{}
	tournamentID := 1

	teamDeleted := &volleynet.TournamentTeam{
		TournamentID: tournamentID,
		Seed:         1,
		Player1: &volleynet.Player{ID: 1},
		Player2: &volleynet.Player{ID: 2},
	}

	teamOutdated := &volleynet.TournamentTeam{
		TournamentID: tournamentID,
		Seed:         2,
		Player1: &volleynet.Player{ID: 5},
		Player2: &volleynet.Player{ID: 6},
	}

	teamNew := &volleynet.TournamentTeam{
		TournamentID: tournamentID,
		Seed:         2,
		Player1: &volleynet.Player{ID: 3},
		Player2: &volleynet.Player{ID: 4},
	}

	teamUpdated := *teamOutdated
	teamUpdated.Seed = 3

	old := []*volleynet.TournamentTeam{
		teamDeleted,
		teamOutdated,
	}

	new := []*volleynet.TournamentTeam{
		&teamUpdated,
		teamNew,
	}

	service.syncTournamentTeams(changes, old, new)

	test.Assert(t,"Service.syncTournamentTeam(...) want: len(changes.New) == 1, got: %d",len(changes.New) != 1 , len(changes.New))
	test.Assert(t, "Service.syncTournamentTeam(...) want: len(changes.Update) == 1, got: %d", len(changes.Update) != 1, len(changes.Update) )
	test.Assert(t, "Service.syncTournamentTeam(...) want: len(changes.Delete) == 1, got: %d", len(changes.Delete) != 1,len(changes.Delete))
}
