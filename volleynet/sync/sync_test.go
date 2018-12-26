package sync

import (
	"os"
	"testing"
	"time"

	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/volleynet/mocks"
	"github.com/raphi011/scores/volleynet/parse"
)

func syncMock() (*mocks.ClientMock, *mocks.VolleynetRepositoryMock, *Service) {
	clientMock := new(mocks.ClientMock)
	volleynetMock := new(mocks.VolleynetRepositoryMock)

	Service := &Service{
		Client:           clientMock,
		VolleynetRepository: volleynetMock,
	}

	return clientMock, volleynetMock, Service
}

func TestSyncLadder(t *testing.T) {
	persistedPlayers := []volleynet.Player{volleynet.Player{PlayerInfo: volleynet.PlayerInfo{ID: 1}, TotalPoints: 100, Rank: 96}}
	clientPlayers := []volleynet.Player{volleynet.Player{PlayerInfo: volleynet.PlayerInfo{ID: 1}, TotalPoints: 125, Rank: 60}}

	clientMock, volleynetMock, Service := syncMock()

	clientMock.On("Ladder", "M").Return(clientPlayers, nil)
	volleynetMock.On("AllPlayers").Return(persistedPlayers, nil)
	volleynetMock.On("UpdatePlayer", &volleynet.Player{
		PlayerInfo:  volleynet.PlayerInfo{ID: 1},
		TotalPoints: 125,
		Rank:        60,
	}).Return(nil)

	report, err := Service.Ladder("M")

	if err != nil {
		t.Error(err)
	}

	if report.UpdatedPlayers != 1 {
		t.Errorf("Service.Ladder(\"M\") want: .UpdatedPlayers = 1, got: %d", report.UpdatedPlayers)
	}
}

func TestSyncTournamentInformation(t *testing.T) {
	response, _ := os.Open("../testdata/upcoming.html")
	tournament, _ := parse.FullTournament(response, time.Now(), volleynet.Tournament{Status: volleynet.StatusUpcoming, ID: 22231})

	syncInfos := Tournaments(tournament, &volleynet.Tournament{ID: 22231, Status: volleynet.StatusUpcoming})

	if syncInfos.Type != SyncTournamentUpcoming {
		t.Fatalf("Service.Tournaments() want: %s, got: %s", SyncTournamentUpcoming, syncInfos.Type)
	}
}

func TestSyncTournaments(t *testing.T) {
	clientTournaments := []volleynet.Tournament{volleynet.Tournament{
		ID:     1,
		Status: volleynet.StatusUpcoming,
	}}
	persistedTournaments := []volleynet.FullTournament{
		volleynet.FullTournament{
			Tournament: volleynet.Tournament{ID: 1, Status: volleynet.StatusUpcoming}},
	}
	clientFullTournament := []volleynet.FullTournament{
		volleynet.FullTournament{Tournament: volleynet.Tournament{
			ID:     1,
			Status: volleynet.StatusUpcoming,
			Name:   "New name",
		},
			Teams: []volleynet.TournamentTeam{},
		},
	}

	clientMock, volleynetMock, Service := syncMock()

	gender := "M"
	league := "AMATEUR LEAGUE"
	season := 2018

	clientMock.On("AllTournaments", gender, league, season).Return(clientTournaments, nil)
	volleynetMock.On("Tournament", persistedTournaments[0].ID).Return(&persistedTournaments[0], nil)
	clientMock.On("ComplementMultipleTournaments", clientTournaments).Return(clientFullTournament, nil)
	volleynetMock.On("UpdateTournament", &clientFullTournament[0]).Return(nil)
	volleynetMock.On("TournamentTeams", 1).Return([]volleynet.TournamentTeam{}, nil)
	volleynetMock.On("AllPlayers").Return([]volleynet.Player{}, nil)

	report, err := Service.Tournaments("M", "AMATEUR LEAGUE", 2018)

	if err != nil {
		t.Fatal(err)
	}

	if len(report.Tournament.Update) != 1 {
		t.Fatalf("Service.Tournaments(\"M\") want: .UpdatedTournaments = 1, got: %d", len(report.Tournament.Update))
	}
}
