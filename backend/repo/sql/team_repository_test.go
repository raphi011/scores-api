package sql

import (
	"testing"

	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/volleynet"
)

func TestCreateTeam(t *testing.T) {
	db := SetupDB(t)
	teamRepo := &teamRepository{DB: db}

	ps := CreatePlayers(t, db,
		P{ Gender: "m", TotalPoints: 5, LadderRank: 1, ID: 1 },
		P{ Gender: "m", TotalPoints: 4, LadderRank: 2, ID: 2 },
		P{ Gender: "m", TotalPoints: 0, LadderRank: 0, ID: 3 },
		P{ Gender: "w", TotalPoints: 4, LadderRank: 1, ID: 4 },
	)

	ts := CreateTournaments(t, db,
		T{ ID: 1 },
	)

	_, err := teamRepo.New(&volleynet.TournamentTeam{
		TournamentID: ts[0].ID,
		Player1: ps[0],
		Player2: ps[1],
		PrizeMoney: 4,
		Result: 1,
		Seed: 3,
		WonPoints: 25,
	})

	test.Check(t, "teamRepository.New(), err: %v", err)
}

func TestTournamentTeams(t *testing.T) {
	db := SetupDB(t)
	teamRepo := &teamRepository{DB: db}

	ps := CreatePlayers(t, db,
		P{ Gender: "m", ID: 1 },
		P{ Gender: "m", ID: 2 },
		P{ Gender: "m", ID: 3 },
		P{ Gender: "m", ID: 4 },
		P{ Gender: "m", ID: 5 },
		P{ Gender: "m", ID: 6 },
		P{ Gender: "m", ID: 7 },
		P{ Gender: "m", ID: 8 },
	)

	ts := CreateTournaments(t, db,
		T{ ID: 1 },
	)

	CreateTeams(t, db,
		TT{TournamentID: ts[0].ID, Player1: ps[0], Player2: ps[1] },
		TT{TournamentID: ts[0].ID, Player1: ps[2], Player2: ps[3] },
		TT{TournamentID: ts[0].ID, Player1: ps[4], Player2: ps[5] },
	)

	tournamentTeams, err := teamRepo.ByTournament(ts[0].ID)

	test.Check(t, "teamRepo.ByTournament() failed: %v", err)
	test.Assert(t, "teamRepository.ByTournament(), want len(tournaments) == 3, got: %d", len(tournamentTeams) == 3, len(tournamentTeams))
}
