package sql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

func TestCreatePlayer(t *testing.T) {
	db := setupDB(t)
	playerRepo := &PlayerRepository{DB: db}

	player, err := playerRepo.New(&volleynet.Player{
		PlayerInfo: volleynet.PlayerInfo{
			TrackedModel: scores.TrackedModel{ Model: scores.Model{ID: 1 }},
		},
	})

	if err != nil {
		t.Fatalf("playerRepository.New(), err: %v", err)
	}

	persistedPlayer, err := playerRepo.Get(1)
	assert(t, "playerRepo.Get() failed", err)

	if diff := cmp.Diff(persistedPlayer, player); diff != "" {
		t.Fatalf("players are not equal:\n%s", diff)
	}
}

func TestUpdatePlayer(t *testing.T) {
	db := setupDB(t)
	playerRepo := &PlayerRepository{DB: db}

	player, err := playerRepo.New(&volleynet.Player{
		PlayerInfo: volleynet.PlayerInfo{
			TrackedModel: scores.TrackedModel{ Model: scores.Model{ID: 1 }},
		},
	})
	assert(t, "couldn't persist player: %v", err)

	player.FirstName = "test!"

	err = playerRepo.Update(player)
	assert(t, "couldnt update player: %v", err)

	updatedPlayer, err := playerRepo.Get(player.ID)
	assert(t, "couldnt get player: %v", err)

	if diff := cmp.Diff(player, updatedPlayer); diff != "" {
		t.Fatalf("players are not equal:\n%s", diff)
	}
}

func TestLadder(t *testing.T) {
	db := setupDB(t)
	playerRepo := &PlayerRepository{DB: db}

	players, err := playerRepo.Ladder("m")
	assert(t, "playerRepo.Ladder() failed", err)

	if len(players) != 0 {
		t.Fatalf("ladder should be empty: %v", err)
	}

	createPlayers(t, db,
		P{ Gender: "m", TotalPoints: 5, Rank: 1, ID: 1 },
		P{ Gender: "m", TotalPoints: 4, Rank: 2, ID: 2 },
		P{ Gender: "m", TotalPoints: 0, Rank: 0, ID: 3 },
		P{ Gender: "w", TotalPoints: 4, Rank: 1, ID: 4 },
	)


	players, err = playerRepo.Ladder("m")
	assert(t, "playerRepo.Ladder() failed", err)

	if len(players) != 2 {
		t.Fatalf("len(ladder) should be 2 but is: %d", len(players))
	}
}
