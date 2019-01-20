package sync

import (
	"github.com/raphi011/scores"
	"github.com/google/go-cmp/cmp/cmpopts"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

// TournamentChanges lists the tournaments that are `New`, `Delete`'d and `Update`'d
// during a sync job
type TournamentChanges struct {
	New    []*volleynet.Tournament
	Delete []*volleynet.Tournament
	Update []*volleynet.Tournament
}

// TournamentSyncInformation contains sync information for two `TournamentInfo`s
type TournamentSyncInformation struct {
	IsNew         bool
	Type          string
	OldTournament *volleynet.Tournament
	NewTournament *volleynet.TournamentInfo
}

// represents the various tournament sync states.
const (
	SyncTournamentNoUpdate           = "SyncTournamentNoUpdate"
	SyncTournamentUpcomingToCanceled = "SyncTournamentUpcomingToCanceled"
	SyncTournamentUpcoming           = "SyncTournamentUpcoming"
	SyncTournamentUpcomingToDone     = "SyncTournamentUpcomingToDone"
	SyncTournamentNew                = "SyncTournamentNew"
)

func (s *Service) syncTournaments(changes *Changes, oldTournaments, newTournaments []*volleynet.Tournament) {
	oldTournamentMap := createTournamentMap(oldTournaments)
	newTournamentMap := createTournamentMap(newTournaments)

	for key, newTournament := range newTournamentMap {
		oldTournament, ok := oldTournamentMap[key]
		if !ok {
			changes.TournamentInfo.New = append(changes.TournamentInfo.New, newTournament)
		} else {
			mergedTournament := MergeTournament(oldTournament, newTournament)

			if hasTournamentChanged(oldTournament, mergedTournament) {
				changes.TournamentInfo.Update = append(changes.TournamentInfo.Update, mergedTournament)
			}
		}

		oldTeams := []*volleynet.TournamentTeam{}
		if oldTournament != nil {
			oldTeams = oldTournament.Teams
		}

		s.syncTournamentTeams(&changes.Team, oldTeams, newTournament.Teams)
	}

	for key, oldTournament := range oldTournamentMap {
		if _, ok := oldTournamentMap[key]; !ok {
			changes.TournamentInfo.Delete = append(changes.TournamentInfo.Delete, oldTournament)
		}
	}
}

func createTournamentMap(tournaments []*volleynet.Tournament) map[int]*volleynet.Tournament {
	tournamentMap := make(map[int]*volleynet.Tournament)

	for i := range tournaments {
		t := tournaments[i]
		tournamentMap[t.ID] = t
	}

	return tournamentMap
}

func tournamentSyncType(persisted *volleynet.Tournament, current *volleynet.TournamentInfo) string {
	if persisted == nil {
		return SyncTournamentNew
	}
	if persisted.Status != volleynet.StatusUpcoming {
		return SyncTournamentNoUpdate
	}
	if current.Status == volleynet.StatusCanceled {
		return SyncTournamentUpcomingToCanceled
	}
	if current.Status == volleynet.StatusUpcoming {
		return SyncTournamentUpcoming
	}
	if current.Status == volleynet.StatusDone {
		return SyncTournamentUpcomingToDone
	}

	return ""
}

func hasTournamentChanged(old, new *volleynet.Tournament) bool {
	new.UpdatedAt = time.Time{}
	old.UpdatedAt = time.Time{}

	return !cmp.Equal(new, old, cmp.Options{ cmpopts.IgnoreUnexported(scores.Tracked{}) })
}

// Tournaments figures out if and how a tournament needs to be synchronized
func Tournaments(persisted *volleynet.Tournament, current *volleynet.TournamentInfo) TournamentSyncInformation {
	syncType := tournamentSyncType(persisted, current)

	return TournamentSyncInformation{
		OldTournament: persisted,
		NewTournament: current,
		IsNew:         persisted == nil,
		Type:          syncType,
	}
}

func (s *Service) persistTournaments(changes *TournamentChanges) error {
	for _, new := range changes.New {
		_, err := s.TournamentRepo.New(new)

		if err != nil {
			return errors.Wrap(err, "persisting new tournament failed")
		}
	}

	for _, update := range changes.Update {
		err := s.TournamentRepo.Update(update)

		if err != nil {
			return errors.Wrap(err, "persisting updated tournament failed")
		}
	}

	return nil
}
