package volleynet

import (
	"github.com/stretchr/testify/mock"
)

type VolleynetServiceMock struct {
	mock.Mock
}

func (m *VolleynetServiceMock) Tournament(tournamentID int) (*FullTournament, error) {
	args := m.Called(tournamentID)

	return args.Get(0).(*FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) AllTournaments() ([]FullTournament, error) {
	args := m.Called()

	return args.Get(0).([]FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) SeasonTournaments(season int) ([]FullTournament, error) {
	args := m.Called(season)

	return args.Get(0).([]FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) GetTournaments(gender, league string, season int) ([]FullTournament, error) {
	args := m.Called(gender, league, season)

	return args.Get(0).([]FullTournament), args.Error(1)
}

func (m *VolleynetServiceMock) NewTournament(t *FullTournament) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) UpdateTournament(t *FullTournament) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) UpdateTournamentTeam(t *TournamentTeam) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) UpdateTournamentTeams(teams []TournamentTeam) error {
	args := m.Called(teams)

	return args.Error(0)
}

func (m *VolleynetServiceMock) NewTeam(t *TournamentTeam) error {
	args := m.Called(t)

	return args.Error(0)
}

func (m *VolleynetServiceMock) NewTeams(teams []TournamentTeam) error {
	args := m.Called(teams)

	return args.Error(0)
}

func (m *VolleynetServiceMock) TournamentTeams(tournamentID int) ([]TournamentTeam, error) {
	args := m.Called(tournamentID)

	return args.Get(0).([]TournamentTeam), args.Error(1)
}

func (m *VolleynetServiceMock) SearchPlayers() ([]Player, error) {
	args := m.Called()

	return args.Get(0).([]Player), args.Error(1)
}

func (m *VolleynetServiceMock) NewPlayer(p *Player) error {
	args := m.Called(p)

	return args.Error(0)
}

func (m *VolleynetServiceMock) AllPlayers() ([]Player, error) {
	args := m.Called()

	return args.Get(0).([]Player), args.Error(1)
}

func (m *VolleynetServiceMock) UpdatePlayer(p *Player) error {
	args := m.Called(p)

	return args.Error(0)
}
