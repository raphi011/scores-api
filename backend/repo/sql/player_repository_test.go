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

