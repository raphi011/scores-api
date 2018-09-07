package sync

import (
	"time"

	"github.com/pkg/errors"
	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/volleynet/client"
)

// PersistanceService is a way to persist and retrieve tournaments
type PersistanceService interface {
	Tournament(tournamentID int) (*volleynet.FullTournament, error)
	UpdateTournament(t *volleynet.FullTournament) error
	NewTournament(t *volleynet.FullTournament) error
	SeasonTournaments(season int) ([]volleynet.FullTournament, error)

	AllPlayers() ([]volleynet.Player, error)
	NewPlayer(p *volleynet.Player) error
	UpdatePlayer(p *volleynet.Player) error
	Player(id int) (*volleynet.Player, error)

	TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error)
	NewTeam(t *volleynet.TournamentTeam) error
	DeleteTeam(t *volleynet.TournamentTeam) error
	UpdateTournamentTeam(t *volleynet.TournamentTeam) error
}

type Changes struct {
	Tournament     *TournamentChanges
	Team           *TeamChanges
	ScrapeDuration time.Duration
	Success        bool
}

type SyncService struct {
	VolleynetService PersistanceService
	Client           client.Client
}

func (s *SyncService) Tournaments(gender, league string, season int) (*Changes, error) {
	report := &Changes{Tournament: &TournamentChanges{}, Team: &TeamChanges{}}

	start := time.Now()
	current, err := s.Client.AllTournaments(gender, league, season)

	if err != nil {
		return nil, errors.Wrap(err, "loading the client tournament list failed")
	}

	persistedTournaments := []volleynet.FullTournament{}
	toDownload := []volleynet.Tournament{}

	for _, t := range current {
		persisted, err := s.VolleynetService.Tournament(t.ID)

		if err != nil {
			return report, errors.Wrap(err, "loading the persisted tournament failed")
		}

		syncInfo := SyncTournaments(persisted, &t)

		if syncInfo.Type == SyncTournamentNoUpdate {
			continue
		} else if syncInfo.Type != SyncTournamentNew {
			persisted.Teams, err = s.VolleynetService.TournamentTeams(t.ID)

			if err != nil {
				return report, errors.Wrap(err, "loading the persisted tournament teams failed")
			}

			persistedTournaments = append(persistedTournaments, *persisted)
		}

		toDownload = append(toDownload, t)
	}

	if len(toDownload) == 0 {
		return report, nil
	}

	currentTournaments, err := s.Client.ComplementMultipleTournaments(toDownload)

	s.syncTournaments(report, persistedTournaments, currentTournaments)

	err = s.persistChanges(report)

	report.ScrapeDuration = time.Since(start) / time.Millisecond

	return report, errors.Wrap(err, "persisting tournament scrape changes failed")
}

func (s *SyncService) persistChanges(report *Changes) error {
	err := s.addMissingPlayers(report.Team.New)

	if err != nil {
		return err
	}

	err = s.persistTournaments(report.Tournament)

	if err != nil {
		return err
	}

	return s.persistTeams(report.Team)
}
