package mocks

import (
	"github.com/raphi011/scores/volleynet"
	"github.com/stretchr/testify/mock"
)

type VolleynetServiceMock struct {
	mock.Mock
}

func (m *VolleynetServiceMock) Tournament(tournamentID int) (*volleynet.FullTournament, error) {
	args := m.Called(tournamentID)

	return args.Get(0).(*volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) AllTournaments() ([]volleynet.FullTournament, error) {
	args := m.Called()

	return args.Get(0).([]volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) SeasonTournaments(season int) ([]volleynet.FullTournament, error) {
	args := m.Called(season)

	return args.Get(0).([]volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) GetTournaments(gender, league string, season int) ([]volleynet.FullTournament, error) {
	args := m.Called(gender, league, season)

	return args.Get(0).([]volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) NewTournament(t *volleynet.FullTournament) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) UpdateTournament(t *volleynet.FullTournament) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) UpdateTournamentTeam(t *volleynet.TournamentTeam) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) UpdateTournamentTeams(teams []volleynet.TournamentTeam) error {
	args := m.Called(teams)

	return args.Error(0)
}

func (m *VolleynetServiceMock) NewTeam(t *volleynet.TournamentTeam) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) NewTeams(teams []volleynet.TournamentTeam) error {
	args := m.Called(teams)

	return args.Error(0)
}

func (m *VolleynetServiceMock) TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error) {
	args := m.Called(tournamentID)

	return args.Get(0).([]volleynet.TournamentTeam), args.Error(1)
}

func (m *VolleynetServiceMock) SearchPlayers() ([]volleynet.Player, error) {
	args := m.Called()

	return args.Get(0).([]volleynet.Player), args.Error(1)
}

func (m *VolleynetServiceMock) NewPlayer(p *volleynet.Player) error {
	args := m.Called(p)

	return args.Error(0)
}

func (m *VolleynetServiceMock) AllPlayers() ([]volleynet.Player, error) {
	args := m.Called()

	return args.Get(0).([]volleynet.Player), args.Error(1)
}

func (m *VolleynetServiceMock) Player(id int) (*volleynet.Player, error) {
	args := m.Called(id)

	return args.Get(0).(*volleynet.Player), args.Error(1)
}

func (m *VolleynetServiceMock) UpdatePlayer(p *volleynet.Player) error {
	args := m.Called(p)

	return args.Error(0)
}
