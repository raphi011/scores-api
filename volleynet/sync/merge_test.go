package sync

import (
	"testing"

	"github.com/raphi011/scores/volleynet"
)

func TestMergeTournamentTeam(t *testing.T) {
	oldTeam := &volleynet.TournamentTeam{
		Seed:        1,
		TotalPoints: 300,
	}

	newTeam := &volleynet.TournamentTeam{
		PrizeMoney: 200,
		Rank:       2,
		WonPoints:  35,
	}

	result := MergeTournamentTeam(SyncTeamDone, oldTeam, newTeam)

	if result.PrizeMoney != newTeam.PrizeMoney ||
		result.Rank != newTeam.Rank ||
		result.WonPoints != newTeam.WonPoints {
		t.Errorf("MergeTournamentTeam(SyncTeamDone, oldTeam, newTeam) did not update correctly")
	}
}

func TestMergeCanceledTournament(t *testing.T) {
	old := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			Status: volleynet.StatusUpcoming,
		},
	}

	new := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			Status: volleynet.StatusCanceled,
		},
	}

	merged := MergeTournament(SyncTournamentUpcomingToCanceled, old, new)

	if merged.Status != volleynet.StatusCanceled {
		t.Errorf("MergeTournament(SyncTournamentUpcomingToCanceled, old, new) did not update correctly")
	}
}
