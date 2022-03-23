package sync

import (
	"errors"
	"fmt"
	"time"

	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/events"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/volleynet"
	"github.com/raphi011/scores-api/volleynet/client"
)

// Changes contains metrics of a scrape job
type Changes struct {
	TournamentInfo TournamentChanges
	Team           TeamChanges
	ScrapeDuration time.Duration
	Success        bool
}

// Service allows loading and synchronizing of the volleynetpage.
type Service struct {
	TeamRepo       repo.TeamRepository
	TournamentRepo repo.TournamentRepository
	PlayerRepo     repo.PlayerRepository

	Client        client.Client
	Subscriptions events.Publisher
}

// Tournaments loads tournaments of a certain `gender`, `league` and `season` and
// synchronizes + updates them (if necessary) in the repository.
func (s *Service) Tournaments(gender, league string, season int) error {
	report := &Changes{TournamentInfo: TournamentChanges{}, Team: TeamChanges{}}
	s.publishStartScrapeEvent("tournaments", time.Now())

	current, err := s.Client.Tournaments(gender, league, season)

	if err != nil {
		return fmt.Errorf("loading the client tournament list failed: %w", err)
	}

	persistedTournaments := []*volleynet.Tournament{}
	toDownload := []*volleynet.TournamentInfo{}

	for _, t := range current {
		persisted, err := s.TournamentRepo.Get(t.ID)

		if errors.Is(err, scores.ErrNotFound) {
			persisted = nil
		} else if err != nil {
			return fmt.Errorf("loading the persisted tournament failed: %w", err)
		}

		syncInfo := Tournaments(persisted, t)

		if syncInfo.Type == SyncTournamentNoUpdate {
			continue
		} else if syncInfo.Type != SyncTournamentNew {
			persisted.Teams, err = s.TeamRepo.ByTournament(t.ID)

			if err != nil {
				return fmt.Errorf("loading the persisted tournament teams failed: %w", err)
			}

			persistedTournaments = append(persistedTournaments, persisted)
		}

		toDownload = append(toDownload, t)
	}

	if len(toDownload) == 0 {
		return nil
	}

	currentTournaments := make([]*volleynet.Tournament, len(toDownload))

	for i, t := range toDownload {
		currentTournaments[i], err = s.Client.ComplementTournament(t)

		if err != nil {
			// remove it from the tournaments for now
			currentTournaments = append(currentTournaments[:i], currentTournaments[i+1:]...)
		}
	}

	s.syncTournaments(report, persistedTournaments, currentTournaments)

	err = s.persistChanges(report)

	s.publishEndScrapeEvent(report, time.Now())

	return fmt.Errorf("sync failed: %w", err)
}

func (s *Service) persistChanges(report *Changes) error {
	err := s.addMissingPlayers(report.Team.New)

	if err != nil {
		return err
	}

	err = s.persistTournaments(&report.TournamentInfo)

	if err != nil {
		return err
	}

	return s.persistTeams(&report.Team)
}
