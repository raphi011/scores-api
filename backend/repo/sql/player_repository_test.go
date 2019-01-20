package sql

import (
	"time"
	"testing"

	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/volleynet"
)

func TestCreatePlayer(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}
	now := time.Now()

	player, err := playerRepo.New(&volleynet.Player{ ID: 1 })
	player.SetTestTime(&now)

	if err != nil {
		t.Fatalf("playerRepository.New(), err: %v", err)
	}

	persistedPlayer, err := playerRepo.Get(1)

	test.Check(t, "playerRepo.Get() failed: %v", err)
	test.Compare(t, "players are not equal:\n%s", persistedPlayer, player)
}

func TestUpdatePlayer(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}

	player, err := playerRepo.New(&volleynet.Player{ ID: 1 })
	test.Check(t, "couldn't persist player: %v", err)

	player.FirstName = "test!"

	err = playerRepo.Update(player)
	test.Check(t, "couldnt update player: %v", err)

	updatedPlayer, err := playerRepo.Get(player.ID)

	test.Check(t, "couldnt get player: %v", err)
	test.Compare(t,"players are not equal:\n%s", player, updatedPlayer)
}

func TestLadder(t *testing.T) {
	db := SetupDB(t)
	playerRepo := &playerRepository{DB: db}

	players, err := playerRepo.Ladder("m")

	test.Check(t, "playerRepo.Ladder() failed: %v", err)
	test.Assert(t, "ladder should be empty", len(players) == 0)

	CreatePlayers(t, db,
		P{ Gender: "m", TotalPoints: 5, Rank: 1, ID: 1 },
		P{ Gender: "m", TotalPoints: 4, Rank: 2, ID: 2 },
		P{ Gender: "m", TotalPoints: 0, Rank: 0, ID: 3 },
		P{ Gender: "w", TotalPoints: 4, Rank: 1, ID: 4 },
	)

	players, err = playerRepo.Ladder("m")

	test.Check(t, "playerRepo.Ladder() failed: %v", err)
	test.Assert(t, "len(ladder) should be 2 but is: %d", len(players) == 2, len(players))
}
