package sql

import (
	"testing"

	"github.com/raphi011/scores/volleynet"
)

func TestCreateTeam(t *testing.T) {
	db := setupDB(t)
	teamRepo := &TeamRepository{DB: db}

	ps := createPlayers(t, db,
		P{ Gender: "m", TotalPoints: 5, Rank: 1, ID: 1 },
		P{ Gender: "m", TotalPoints: 4, Rank: 2, ID: 2 },
		P{ Gender: "m", TotalPoints: 0, Rank: 0, ID: 3 },
		P{ Gender: "w", TotalPoints: 4, Rank: 1, ID: 4 },
	)

	ts := createTournaments(t, db,
		T{ ID: 1 },
	)

	_, err := teamRepo.New(&volleynet.TournamentTeam{
		TournamentID: ts[0].ID,
		Player1: ps[0],
		Player2: ps[1],
		PrizeMoney: 4,
		Rank: 1,
		Seed: 3,
		WonPoints: 25,
	})

	if err != nil {
		t.Fatalf("tournamentTeamRepository.New(), err: %s", err)
	}
}
