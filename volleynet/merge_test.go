package volleynet

import "testing"

func TestMergeTournamentTeam(t *testing.T) {
	oldTeam := &TournamentTeam{
		Seed:        1,
		TotalPoints: 300,
	}

	newTeam := &TournamentTeam{
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
	old := &FullTournament{
		Tournament: Tournament{
			Status: StatusUpcoming,
		},
	}

	new := &FullTournament{
		Tournament: Tournament{
			Status: StatusCanceled,
		},
	}

	merged := MergeTournament(SyncTournamentUpcomingToCanceled, old, new)

	if merged.Status != StatusCanceled {
		t.Errorf("MergeTournament(SyncTournamentUpcomingToCanceled, old, new) did not update correctly")
	}
}
