package mocks

import (
	"github.com/raphi011/scores/volleynet"
	"github.com/stretchr/testify/mock"
)

type VolleynetRepositoryMock struct {
	mock.Mock
}

func (m *VolleynetRepositoryMock) Tournament(tournamentID int) (*volleynet.FullTournament, error) {
	args := m.Called(tournamentID)

	return args.Get(0).(*volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetRepositoryMock) AllTournaments() ([]*volleynet.FullTournament, error) {
	args := m.Called()

	return args.Get(0).([]*volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetRepositoryMock) SeasonTournaments(season int) ([]*volleynet.FullTournament, error) {
	args := m.Called(season)

	return args.Get(0).([]*volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetRepositoryMock) GetTournaments(gender, league string, season int) ([]*volleynet.FullTournament, error) {
	args := m.Called(gender, league, season)

	return args.Get(0).([]*volleynet.FullTournament), args.Error(1)
}

func (m *VolleynetRepositoryMock) NewTournament(t *volleynet.FullTournament) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) UpdateTournament(t *volleynet.FullTournament) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) UpdateTournamentTeam(t *volleynet.TournamentTeam) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) UpdateTournamentTeams(teams []volleynet.TournamentTeam) error {
	args := m.Called(teams)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) NewTeam(t *volleynet.TournamentTeam) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) DeleteTeam(t *volleynet.TournamentTeam) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) NewTeams(teams []volleynet.TournamentTeam) error {
	args := m.Called(teams)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error) {
	args := m.Called(tournamentID)

	return args.Get(0).([]volleynet.TournamentTeam), args.Error(1)
}

func (m *VolleynetRepositoryMock) SearchPlayers() ([]volleynet.Player, error) {
	args := m.Called()

	return args.Get(0).([]volleynet.Player), args.Error(1)
}

func (m *VolleynetRepositoryMock) NewPlayer(p *volleynet.Player) error {
	args := m.Called(p)

	return args.Error(0)
}

func (m *VolleynetRepositoryMock) AllPlayers() ([]volleynet.Player, error) {
	args := m.Called()

	return args.Get(0).([]volleynet.Player), args.Error(1)
}

func (m *VolleynetRepositoryMock) Player(id int) (*volleynet.Player, error) {
	args := m.Called(id)

	return args.Get(0).(*volleynet.Player), args.Error(1)
}

func (m *VolleynetRepositoryMock) UpdatePlayer(p *volleynet.Player) error {
	args := m.Called(p)

	return args.Error(0)
}
