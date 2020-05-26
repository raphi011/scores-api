// +build repository

package sql

import (
	"testing"

	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/test"
	"github.com/raphi011/scores-api/volleynet"
)

func TestCreatePlayer(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}

	player := &volleynet.Player{ID: 1}

	_, err := playerRepo.New(player)
	test.Check(t, "playerRepository.New(), err: %v", err)

	persistedPlayer, err := playerRepo.Get(1)

	test.Check(t, "playerRepo.Get() failed: %v", err)
	test.Compare(t, "players are not equal:\n%s", persistedPlayer, player)
}

func TestUpdatePlayer(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}

	player, err := playerRepo.New(&volleynet.Player{ID: 1})
	test.Check(t, "couldn't persist player: %v", err)

	player.FirstName = "test!"

	err = playerRepo.Update(player)
	test.Check(t, "couldnt update player: %v", err)

	updatedPlayer, err := playerRepo.Get(player.ID)

	test.Check(t, "couldnt get player: %v", err)
	test.Compare(t, "players are not equal:\n%s", player, updatedPlayer)
}

func TestLadder(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}

	players, err := playerRepo.Ladder("m")

	test.Check(t, "playerRepo.Ladder() failed: %v", err)
	test.Assert(t, "ladder should be empty", len(players) == 0)

	CreatePlayers(t, db,
		P{Gender: "m", TotalPoints: 5, LadderRank: 1, ID: 1},
		P{Gender: "m", TotalPoints: 4, LadderRank: 2, ID: 2},
		P{Gender: "m", TotalPoints: 0, LadderRank: 0, ID: 3},
		P{Gender: "w", TotalPoints: 4, LadderRank: 1, ID: 4},
	)

	players, err = playerRepo.Ladder("m")

	test.Check(t, "playerRepo.Ladder() failed: %v", err)
	test.Assert(t, "len(ladder) should be 2 but is: %d", len(players) == 2, len(players))
}

func TestPreviousPartners(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}

	p := CreatePlayers(t, db,
		P{Gender: "m", TotalPoints: 5, LadderRank: 1, ID: 1},
		P{Gender: "m", TotalPoints: 4, LadderRank: 2, ID: 2},
		P{Gender: "m", TotalPoints: 0, LadderRank: 0, ID: 3},
		P{Gender: "w", TotalPoints: 4, LadderRank: 1, ID: 4},
	)

	CreateTournaments(t, db,
		T{ID: 1},
		T{ID: 2},
	)

	CreateTeams(t, db,
		TT{TournamentID: 1, Player1: p[0], Player2: p[1]},
		TT{TournamentID: 1, Player1: p[2], Player2: p[3]},
		TT{TournamentID: 2, Player1: p[0], Player2: p[2]},
	)

	players, err := playerRepo.PreviousPartners(1)

	test.Check(t, "playerRepo.PreviousPartners() failed: %v", err)
	test.Assert(t, "len(PreviousPartners) should be 2", len(players) == 2)
}

func TestSearchPlayers(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}

	CreatePlayers(t, db,
		P{Gender: "m", FirstName: "Richard", LastName: "Bosse", ID: 1},
		P{Gender: "m", FirstName: "Dominik", LastName: "Rieder", ID: 2},
		P{Gender: "w", FirstName: "Romana", LastName: "Musterfrau", ID: 3},
		P{Gender: "m", FirstName: "Roman", LastName: "Gutleber", ID: 4},
	)

	players, err := playerRepo.Search(repo.PlayerFilter{
		Gender:    "m",
		FirstName: "R",
	})

	test.Check(t, "playerRepo.Search() failed: %v", err)
	test.Assert(t, "len(Search) should be 2 but is %d", len(players) == 2, len(players))
}
