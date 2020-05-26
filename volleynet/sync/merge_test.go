package sync

import (
	"testing"

	"github.com/raphi011/scores-api/volleynet"
)

func TestMergeTournamentTeam(t *testing.T) {
	oldTeam := &volleynet.TournamentTeam{
		Seed:        1,
		TotalPoints: 300,
	}

	newTeam := &volleynet.TournamentTeam{
		PrizeMoney: 200,
		Result:     2,
		WonPoints:  35,
	}

	result := MergeTournamentTeam(oldTeam, newTeam)

	if result.PrizeMoney != newTeam.PrizeMoney ||
		result.Result != newTeam.Result ||
		result.WonPoints != newTeam.WonPoints {
		t.Errorf("MergeTournamentTeam(oldTeam, newTeam) did not update correctly")
	}
}

func TestMergeCanceledTournament(t *testing.T) {
	old := &volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			Status: volleynet.StatusUpcoming,
		},
	}

	new := &volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			Status: volleynet.StatusCanceled,
		},
	}

	merged := MergeTournament(old, new)

	if merged.Status != volleynet.StatusCanceled {
		t.Errorf("MergeTournament(old, new) did not update correctly")
	}
}
