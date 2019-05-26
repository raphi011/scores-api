package sync

import (
	"github.com/pkg/errors"
	"github.com/raphi011/scores/volleynet"
)

// LadderSyncReport contains metrics of a Ladder sync job
type LadderSyncReport struct {
	NewPlayers     int
	UpdatedPlayers int
}

// Ladder synchronizes player and rank data of all players of a certain `gender`
func (s *Service) Ladder(gender string) (*LadderSyncReport, error) {
	ranks, err := s.Client.Ladder(gender)
	report := &LadderSyncReport{}

	if err != nil {
		return nil, errors.Wrap(err, "loading the ladder failed")
	}

	persisted, err := s.PlayerRepo.ByGender(gender)

	if err != nil {
		return nil, errors.Wrap(err, "loading persisted players failed")
	}

	syncInfos := Players(persisted, ranks...)

	for _, info := range syncInfos {
		if info.IsNew {
			s.Log.Debugf("adding player id: %v, name: %s %s",
				info.NewPlayer.ID,
				info.NewPlayer.FirstName,
				info.NewPlayer.LastName)

			_, err = s.PlayerRepo.New(info.NewPlayer)
			report.NewPlayers++

		} else {
			merged := MergePlayer(info.OldPlayer, info.NewPlayer)

			s.Log.Debugf("updating player id: %d, name: %s %s",
				info.NewPlayer.ID,
				merged.FirstName,
				merged.LastName)

			err = s.PlayerRepo.Update(merged)
			report.UpdatedPlayers++

		}

		if err != nil {
			return nil, errors.Wrap(err, "sync player failed")
		}
	}

	return report, nil
}

// PlayerSyncInformation contains sync information for two `Player`s
type PlayerSyncInformation struct {
	IsNew     bool
	OldPlayer *volleynet.Player
	NewPlayer *volleynet.Player
}

// Players takes a slice of current and old `Player`s and finds out which
// one is new and which needs to get updated
func Players(persisted []*volleynet.Player, current ...*volleynet.Player) []PlayerSyncInformation {
	ps := []PlayerSyncInformation{}
	for i := range current {
		newPlayer := current[i]
		oldPlayer := FindPlayer(persisted, newPlayer.ID)

		ps = append(ps, PlayerSyncInformation{
			NewPlayer: newPlayer,
			OldPlayer: oldPlayer,
			IsNew:     oldPlayer == nil,
		})
	}

	return ps
}
