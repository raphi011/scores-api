package sync

import (
	"errors"
	"fmt"

	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/volleynet"
)

// TeamChanges lists the teams that are `New`, `Delete`'d and `Update`'d
// during a sync job.
type TeamChanges struct {
	New    []*volleynet.TournamentTeam
	Delete []*volleynet.TournamentTeam
	Update []*volleynet.TournamentTeam
}

func artificialTeamKey(team *volleynet.TournamentTeam) string {
	return fmt.Sprintf("%d-%d-%d", team.TournamentID, team.Player1.ID, team.Player2.ID)
}

func createTeamMap(teams []*volleynet.TournamentTeam) map[string]*volleynet.TournamentTeam {
	teamMap := make(map[string]*volleynet.TournamentTeam)

	for i := range teams {
		t := teams[i]
		teamMap[artificialTeamKey(t)] = t
	}

	return teamMap
}

func (s *Service) syncTournamentTeams(changes *TeamChanges, oldTeams, newTeams []*volleynet.TournamentTeam) {
	oldTeamMap := createTeamMap(oldTeams)
	newTeamMap := createTeamMap(newTeams)

	for key, newTeam := range newTeamMap {
		if oldTeam, ok := oldTeamMap[key]; !ok {
			changes.New = append(changes.New, newTeam)
		} else {
			mergedTeam := MergeTournamentTeam(oldTeam, newTeam)

			if hasTeamChanged(oldTeam, mergedTeam) {
				changes.Update = append(changes.Update, mergedTeam)
			}
		}
	}

	for key, oldTeam := range oldTeamMap {
		if _, ok := newTeamMap[key]; !ok {
			changes.Delete = append(changes.Delete, oldTeam)
		}
	}
}

func hasTeamChanged(old, new *volleynet.TournamentTeam) bool {
	return *new != *old
}

func (s *Service) persistTeams(changes *TeamChanges) error {
	for _, new := range changes.New {
		_, err := s.TeamRepo.New(new)

		if err != nil {
			return fmt.Errorf("persist new tournament team %+v %w", new, err)
		}
	}

	for _, update := range changes.Update {
		err := s.TeamRepo.Update(update)

		if err != nil {
			return fmt.Errorf("persist updated tournament team %+v %w", update, err)
		}
	}

	for _, delete := range changes.Delete {
		err := s.TeamRepo.Delete(delete)

		if err != nil {
			return fmt.Errorf("persist deleted tournament team %+v %w", delete, err)
		}
	}

	return nil
}

func (s *Service) addMissingPlayers(teams []*volleynet.TournamentTeam) error {
	players := distinctPlayers(teams)

	for _, p := range players {
		err := s.addPlayerIfNeeded(p)

		if err != nil {
			return fmt.Errorf("addMissingPlayers failed: %w", err)
		}
	}

	return nil
}

func distinctPlayers(teams []*volleynet.TournamentTeam) []*volleynet.Player {
	encountered := map[int]bool{}

	distinct := []*volleynet.Player{}

	addIfNotEncountered := func(p *volleynet.Player) {
		if !encountered[p.ID] {
			encountered[p.ID] = true
			distinct = append(distinct, p)
		}
	}

	for _, t := range teams {
		addIfNotEncountered(t.Player1)
		addIfNotEncountered(t.Player2)
	}

	return distinct
}

func (s *Service) addPlayerIfNeeded(player *volleynet.Player) error {
	_, err := s.PlayerRepo.Get(player.ID)

	if errors.Is(err, scores.ErrNotFound) {
		_, err = s.PlayerRepo.New(player)

		return fmt.Errorf("addPlayerIfNeeded: %w", err)
	} else if err != nil {
		return fmt.Errorf("addPlayerIfNeeded: %w", err)
	}

	return nil
}
