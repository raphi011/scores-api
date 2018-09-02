package sync

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores/volleynet"
)

type TeamChanges struct {
	New    []volleynet.TournamentTeam
	Delete []volleynet.TournamentTeam
	Update []volleynet.TournamentTeam
}

func artificialTeamKey(team volleynet.TournamentTeam) string {
	return fmt.Sprintf("%d-%d-%d", team.TournamentID, team.Player1.ID, team.Player2.ID)
}

func createTeamMap(teams []volleynet.TournamentTeam) map[string]volleynet.TournamentTeam {
	teamMap := make(map[string]volleynet.TournamentTeam)

	for i := range teams {
		t := teams[i]
		teamMap[artificialTeamKey(t)] = t
	}

	return teamMap
}

func (s *SyncService) syncTournamentTeams(changes *TeamChanges, oldTeams, newTeams []volleynet.TournamentTeam) {
	oldTeamMap := createTeamMap(oldTeams)
	newTeamMap := createTeamMap(newTeams)

	for key, newTeam := range newTeamMap {
		if oldTeam, ok := oldTeamMap[key]; !ok {
			changes.New = append(changes.New, newTeam)
		} else {
			mergedTeam := MergeTournamentTeam(&oldTeam, &newTeam)

			if hasTeamChanged(oldTeam, *mergedTeam) {
				changes.Update = append(changes.Update, *mergedTeam)
			}
		}
	}

	for key, oldTeam := range oldTeamMap {
		if _, ok := oldTeamMap[key]; !ok {
			changes.Delete = append(changes.Delete, oldTeam)
		}
	}
}

func hasTeamChanged(old, new volleynet.TournamentTeam) bool {
	return !cmp.Equal(new, old)
}

func (s *SyncService) persistTeams(changes *TeamChanges) error {
	for _, new := range changes.New {
		err := s.VolleynetService.NewTeam(&new)

		if err != nil {
			return err
		}
	}

	for _, update := range changes.Update {
		err := s.VolleynetService.UpdateTournamentTeam(&update)

		if err != nil {
			return err
		}
	}

	// for _, delete := range changes.Delete {
	// 	err := s.VolleynetService.Dele(&update)
	// }

	return nil
}
