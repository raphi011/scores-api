package sync

import (
	"time"

	"github.com/pkg/errors"
	"github.com/raphi011/scores/events"
	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/volleynet/client"
)

// Repository persists and retrieve tournaments
type Repository interface {
	Tournament(tournamentID int) (*volleynet.FullTournament, error)
	UpdateTournament(t *volleynet.FullTournament) error
	NewTournament(t *volleynet.FullTournament) error
	SeasonTournaments(season int) ([]*volleynet.FullTournament, error)

	AllPlayers() ([]volleynet.Player, error)
	NewPlayer(p *volleynet.Player) error
	UpdatePlayer(p *volleynet.Player) error
	Player(id int) (*volleynet.Player, error)

	TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error)
	NewTeam(t *volleynet.TournamentTeam) error
	DeleteTeam(t *volleynet.TournamentTeam) error
	UpdateTournamentTeam(t *volleynet.TournamentTeam) error
}

// Changes contains metrics of a scrape job
type Changes struct {
	Tournament     *TournamentChanges
	Team           *TeamChanges
	ScrapeDuration time.Duration
	Success        bool
}
type teamRepository interface{
	ByTournament(tournamentID int) ([]volleynet.TournamentTeam, error)
	New(team *volleynet.TournamentTeam) error
	Update(team *volleynet.TournamentTeam) error
	Delete(team *volleynet.TournamentTeam) error
}

type tournamentRepository interface{
	Get(tournamentID int) (*volleynet.FullTournament, error)
	New(tournament *volleynet.FullTournament) (error)
	Update(tournament *volleynet.FullTournament) (error)
}

type playerRepository interface{
	All() ([]volleynet.Player, error)
	Ladder(gender string) ([]volleynet.Player, error)
	Get(playerID int) (*volleynet.Player, error)
	New(player *volleynet.Player) error
	Update(player *volleynet.Player) error
}

// Service allows loading and synchronizing of the volleynetpage.
type Service struct {
	TeamRepository teamRepository
	TournamentRepository tournamentRepository
	PlayerRepository playerRepository
	Client              client.Client
	Subscriptions       events.Publisher
}

// Tournaments loads tournaments of a certain `gender`, `league` and `season` and
// synchronizes + updates them (if necessary) in the repository.
func (s *Service) Tournaments(gender, league string, season int) error {
	report := &Changes{Tournament: &TournamentChanges{}, Team: &TeamChanges{}}
	s.publishStartScrapeEvent("tournaments", time.Now())

	current, err := s.Client.AllTournaments(gender, league, season)

	if err != nil {
		return errors.Wrap(err, "loading the client tournament list failed")
	}

	persistedTournaments := []volleynet.FullTournament{}
	toDownload := []volleynet.Tournament{}

	for _, t := range current {
		persisted, err := s.TournamentRepository.Get(t.ID)

		if err != nil {
			return errors.Wrap(err, "loading the persisted tournament failed")
		}

		syncInfo := Tournaments(persisted, &t)

		if syncInfo.Type == SyncTournamentNoUpdate {
			continue
		} else if syncInfo.Type != SyncTournamentNew {
			persisted.Teams, err = s.TeamRepository.ByTournament(t.ID)

			if err != nil {
				return errors.Wrap(err, "loading the persisted tournament teams failed")
			}

			persistedTournaments = append(persistedTournaments, *persisted)
		}

		toDownload = append(toDownload, t)
	}

	if len(toDownload) == 0 {
		return nil
	}

	currentTournaments, err := s.Client.ComplementMultipleTournaments(toDownload)

	s.syncTournaments(report, persistedTournaments, currentTournaments)

	err = s.persistChanges(report)

	s.publishEndScrapeEvent(report, time.Now())

	return errors.Wrap(err, "sync failed")
}

func (s *Service) persistChanges(report *Changes) error {
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
