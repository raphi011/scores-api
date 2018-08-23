package volleynet

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
)

// PersistanceService is a way to persist and retrieve tournaments
type PersistanceService interface {
	UpdateTournament(t *FullTournament) error
	NewTournament(t *FullTournament) error
	SeasonTournaments(season int) ([]FullTournament, error)

	AllPlayers() ([]Player, error)
	NewPlayer(p *Player) error
	UpdatePlayer(p *Player) error

	TournamentTeams(tournamentID int) ([]TournamentTeam, error)
	NewTeam(t *TournamentTeam) error
	UpdateTournamentTeam(t *TournamentTeam) error
}

type SyncService struct {
	VolleynetService PersistanceService
	Client           Client
}

type LadderSyncReport struct {
	NewPlayers     int
	UpdatedPlayers int
}

type TournamentSyncReport struct {
	NewTournaments      int
	UpdatedTournaments  int
	CanceledTournaments int
	NewTeams            int
	UpdatedTeams        int
	ScrapeDuration      time.Duration
	Success             bool
	Error               error
}

func (s *SyncService) Tournaments(gender, league string, season int) (*TournamentSyncReport, error) {
	report := &TournamentSyncReport{}

	start := time.Now()
	current, err := s.Client.AllTournaments(gender, league, season)

	if err != nil {
		return nil, errors.Wrap(err, "loading the tournament list failed")
	}

	persisted, err := s.VolleynetService.SeasonTournaments(season)

	syncInformation := SyncTournaments(persisted, current...)

	for _, t := range syncInformation {
		fullTournament, err := s.Client.ComplementTournament(*t.NewTournament)

		if err != nil {
			return nil, errors.Wrap(err, "loading the full tournament failed")
		}

		if t.IsNew {
			report.NewTournaments++
			log.Printf("adding tournament id: %v, name: %v, start: %v",
				fullTournament.ID,
				fullTournament.Name,
				fullTournament.Start)

			err = s.VolleynetService.NewTournament(fullTournament)
		} else {
			report.UpdatedTournaments++
			log.Printf("updating tournament id: %v, name: %v, start: %v, sync: %v",
				fullTournament.ID,
				fullTournament.Name,
				fullTournament.Start,
				t.SyncType,
			)

			mergedTournament := MergeTournament(t.SyncType, t.OldTournament, fullTournament)

			err = s.VolleynetService.UpdateTournament(mergedTournament)
		}

		if err != nil {
			log.Print(err)
			return nil, errors.Wrap(err, "sync tournament failed")
		}

		persistedTeams, err := s.VolleynetService.TournamentTeams(t.NewTournament.ID)

		if err != nil {
			log.Print(err)
			return nil, errors.Wrap(err, "saving a tournament team failed")
		}

		persistedPlayers, err := s.VolleynetService.AllPlayers()

		if err != nil {
			log.Print(err)
			return nil, errors.Wrap(err, "loading players failed")
		}

		syncTournamentTeams := SyncTournamentTeams(t.SyncType, persistedTeams, fullTournament.Teams)

		for _, team := range syncTournamentTeams {
			if team.IsNew {
				persistedPlayers, err = s.addPlayersIfNew(persistedPlayers, team.NewTeam.Player1, team.NewTeam.Player2)

				if err == nil {
					report.NewTeams++

					log.Printf("adding tournament team tournamentid: %v, player1ID: %v, player2ID: %v",
						team.NewTeam.TournamentID,
						team.NewTeam.Player1.ID,
						team.NewTeam.Player2.ID,
					)

					err = s.VolleynetService.NewTeam(team.NewTeam)
				}
			} else {
				report.UpdatedTeams++

				log.Printf("updating tournament team tournamentid: %v, player1ID: %v, player2ID: %v, sync: %v",
					team.NewTeam.TournamentID,
					team.NewTeam.Player1.ID,
					team.NewTeam.Player2.ID,
					t.SyncType,
				)

				if team.OldTeam == nil || team.NewTeam == nil {
					fmt.Print("asdasd")
				}

				mergedTeam := MergeTournamentTeam(team.SyncType, team.OldTeam, team.NewTeam)

				err = s.VolleynetService.UpdateTournamentTeam(mergedTeam)
			}

			if err != nil {
				log.Print(err)
				return nil, errors.Wrap(err, "TODO")
			}
		}

	}

	report.ScrapeDuration = time.Since(start) / time.Millisecond
	report.Success = true

	return nil, nil
}

func (s *SyncService) addPlayersIfNew(persistedPlayers []Player, players ...*Player) (
	[]Player, error) {

	for _, p := range players {
		player := FindPlayer(persistedPlayers, p.ID)

		if player == nil {
			log.Printf("adding missing player id: %v, name: %v",
				p.ID,
				fmt.Sprintf("%v %v", p.FirstName, p.LastName))

			err := s.VolleynetService.NewPlayer(p)

			if err != nil {
				return nil, err
			}
			persistedPlayers = append(persistedPlayers, *p)
		}
	}

	return persistedPlayers, nil
}

func (s *SyncService) Ladder(gender string) (*LadderSyncReport, error) {
	ranks, err := s.Client.Ladder(gender)
	report := &LadderSyncReport{}

	if err != nil {
		return nil, errors.Wrap(err, "loading the ladder failed")
	}

	persisted, err := s.VolleynetService.AllPlayers()

	if err != nil {
		return nil, errors.Wrap(err, "loading persisted players failed")
	}

	syncInfos := SyncPlayers(persisted, ranks...)

	for _, info := range syncInfos {
		if info.IsNew {
			log.Printf("adding player id: %v, name: %s %s",
				info.NewPlayer.ID,
				info.NewPlayer.FirstName,
				info.NewPlayer.LastName)

			err = s.VolleynetService.NewPlayer(info.NewPlayer)
			report.NewPlayers++

		} else {
			merged := MergePlayer(info.OldPlayer, info.NewPlayer)

			log.Printf("updating player id: %d, name: %s %s",
				info.NewPlayer.ID,
				merged.FirstName,
				merged.LastName)

			err = s.VolleynetService.UpdatePlayer(merged)
			report.UpdatedPlayers++

		}

		if err != nil {
			return nil, errors.Wrap(err, "sync player update failed")
		}
	}

	return report, nil
}

// PlayerSyncInformation contains sync information for two `Player`s
type PlayerSyncInformation struct {
	IsNew     bool
	OldPlayer *Player
	NewPlayer *Player
}

// SyncPlayers takes a slice of current and old `Player`s and finds out which
// one is new and which needs to get updated
func SyncPlayers(persisted []Player, current ...Player) []PlayerSyncInformation {
	ps := []PlayerSyncInformation{}
	for i := range current {
		newPlayer := &current[i]
		oldPlayer := FindPlayer(persisted, newPlayer.ID)

		ps = append(ps, PlayerSyncInformation{
			NewPlayer: newPlayer,
			OldPlayer: oldPlayer,
			IsNew:     oldPlayer == nil,
		})
	}

	return ps
}

// TournamentSyncInformation contains sync information for two `Tournament`s
type TournamentSyncInformation struct {
	IsNew         bool
	SyncType      string
	OldTournament *FullTournament
	NewTournament *Tournament
}

// represents the various tournament sync states.
const (
	SyncTournamentNoUpdate           = "SyncTournamentNoUpdate"
	SyncTournamentUpcomingToCanceled = "SyncTournamentUpcomingToCanceled"
	SyncTournamentUpcoming           = "SyncTournamentUpcoming"
	SyncTournamentUpcomingToDone     = "SyncTournamentUpcomingToDone"
	SyncTournamentNew                = "SyncTournamentNew"
)

func tournamentSyncType(persisted *FullTournament, current *Tournament) string {
	if persisted == nil {
		return SyncTournamentNew
	}
	if persisted.Status != StatusUpcoming {
		return SyncTournamentNoUpdate
	}
	if current.Status == StatusCanceled {
		return SyncTournamentUpcomingToCanceled
	}
	if current.Status == StatusUpcoming {
		return SyncTournamentUpcoming
	}
	if current.Status == StatusDone {
		return SyncTournamentUpcomingToDone
	}

	return ""
}

// SyncTournaments finds out if and how tournaments have to be synced
func SyncTournaments(persisted []FullTournament, current ...Tournament) []TournamentSyncInformation {
	ts := []TournamentSyncInformation{}
	for i := range current {
		newTournament := &current[i]
		oldTournament := FindTournament(persisted, newTournament.ID)

		syncType := tournamentSyncType(oldTournament, newTournament)

		if syncType == SyncTournamentNoUpdate {
			continue
		}

		ts = append(ts, TournamentSyncInformation{
			OldTournament: oldTournament,
			NewTournament: newTournament,
			IsNew:         oldTournament == nil,
			SyncType:      syncType,
		})
	}

	return ts
}

// represents the various tournament team sync states.
const (
	SyncTeamNew      = "SyncTeamNew"
	SyncTeamUpcoming = "SyncTeamUpcoming"
	SyncTeamDone     = "SyncTeamDone"
	SyncTeamNoUpdate = "SyncTeamNoUpdate"
)

// TournamentTeamSyncInformation contains sync information for two `TournamentTeam`s
type TournamentTeamSyncInformation struct {
	IsNew    bool
	SyncType string
	OldTeam  *TournamentTeam
	NewTeam  *TournamentTeam
}

func tournamentTeamSyncType(tournamentSyncType string, persisted, current *TournamentTeam) string {
	if tournamentSyncType == SyncTournamentNew {
		return SyncTeamNew
	}
	if tournamentSyncType == SyncTournamentUpcoming {
		return SyncTeamUpcoming
	}
	if tournamentSyncType == SyncTournamentUpcomingToDone {
		return SyncTeamDone
	}
	if tournamentSyncType == SyncTournamentNoUpdate {
		return SyncTeamNoUpdate
	}

	return ""
}

// SyncTournamentTeams finds out if and how tournament teams have to be synced
func SyncTournamentTeams(tournamentSyncType string, persisted, current []TournamentTeam) []TournamentTeamSyncInformation {
	ts := []TournamentTeamSyncInformation{}
	for i := range current {
		newTeam := &current[i]
		oldTeam := FindTeam(persisted, newTeam.TournamentID, newTeam.Player1.ID, newTeam.Player2.ID)
		syncType := tournamentTeamSyncType(tournamentSyncType, oldTeam, newTeam)

		if syncType == SyncTeamNoUpdate {
			continue
		}

		ts = append(ts, TournamentTeamSyncInformation{
			OldTeam:  oldTeam,
			NewTeam:  newTeam,
			IsNew:    oldTeam == nil,
			SyncType: syncType,
		})
	}

	return ts
}
