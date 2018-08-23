package volleynet

import "github.com/stretchr/testify/mock"

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) GetTournamentLink(t *Tournament) string {
	args := m.Called(t)

	return args.String(0)
}

func (m *ClientMock) GetAPITournamentLink(t *Tournament) string {
	args := m.Called(t)

	return args.String(0)
}

func (m *ClientMock) Login(username, password string) (*LoginData, error) {
	args := m.Called(username, password)

	return args.Get(0).(*LoginData), args.Error(1)
}

func (m *ClientMock) AllTournaments(gender, league string, year int) ([]Tournament, error) {
	args := m.Called(gender, league, year)

	return args.Get(0).([]Tournament), args.Error(1)
}

func (m *ClientMock) Ladder(gender string) ([]Player, error) {
	args := m.Called(gender)

	return args.Get(0).([]Player), args.Error(1)
}

func (m *ClientMock) ComplementTournament(tournament Tournament) (*FullTournament, error) {
	args := m.Called(tournament)

	return args.Get(0).(*FullTournament), args.Error(1)
}

func (m *ClientMock) TournamentWithdrawal(tournamentID int) error {
	args := m.Called(tournamentID)

	return args.Error(0)
}

func (m *ClientMock) TournamentEntry(playerName string, playerID, tournamentID int) error {
	args := m.Called(playerName, playerID, tournamentID)

	return args.Error(0)
}

func (m *ClientMock) SearchPlayers(firstName, lastName, birthday string) ([]PlayerInfo, error) {
	args := m.Called(firstName, lastName, birthday)

	return args.Get(0).([]PlayerInfo), args.Error(1)
}
