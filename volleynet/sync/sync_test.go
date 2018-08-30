package sync

import (
	"os"
	"testing"

	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/volleynet/mocks"
	"github.com/raphi011/scores/volleynet/parse"
)

func syncMock() (*mocks.ClientMock, *mocks.VolleynetServiceMock, *SyncService) {
	clientMock := new(mocks.ClientMock)
	volleynetMock := new(mocks.VolleynetServiceMock)

	syncService := &SyncService{
		Client:           clientMock,
		VolleynetService: volleynetMock,
	}

	return clientMock, volleynetMock, syncService
}

func TestSyncLadder(t *testing.T) {
	persistedPlayers := []volleynet.Player{volleynet.Player{PlayerInfo: volleynet.PlayerInfo{ID: 1}, TotalPoints: 100, Rank: 96}}
	clientPlayers := []volleynet.Player{volleynet.Player{PlayerInfo: volleynet.PlayerInfo{ID: 1}, TotalPoints: 125, Rank: 60}}

	clientMock, volleynetMock, syncService := syncMock()

	clientMock.On("Ladder", "M").Return(clientPlayers, nil)
	volleynetMock.On("AllPlayers").Return(persistedPlayers, nil)
	volleynetMock.On("UpdatePlayer", &volleynet.Player{
		PlayerInfo:  volleynet.PlayerInfo{ID: 1},
		TotalPoints: 125,
		Rank:        60,
	}).Return(nil)

	report, err := syncService.Ladder("M")

	if err != nil {
		t.Error(err)
	}

	if report.UpdatedPlayers != 1 {
		t.Errorf("SyncService.Ladder(\"M\") want: .UpdatedPlayers = 1, got: %d", report.UpdatedPlayers)
	}
}

func TestSyncTournamentInformation(t *testing.T) {
	response, _ := os.Open("../testdata/upcoming.html")
	tournament, _ := parse.FullTournament(response, volleynet.Tournament{Status: volleynet.StatusUpcoming, ID: 22231})

	syncInfos := SyncTournaments([]volleynet.FullTournament{*tournament}, volleynet.Tournament{ID: 22231, Status: volleynet.StatusUpcoming})

	if syncInfos[0].SyncType != SyncTournamentUpcoming {
		t.Fatalf("SyncService.SyncTournaments() want: %s, got: %s", SyncTournamentUpcoming, syncInfos[0].SyncType)
	}
}

func TestSyncTournaments(t *testing.T) {
	clientTournaments := []volleynet.Tournament{volleynet.Tournament{
		ID:     1,
		Status: "upcoming",
	}}
	persistedTournaments := []volleynet.FullTournament{
		volleynet.FullTournament{
			Tournament: volleynet.Tournament{ID: 1, Status: "upcoming"}},
	}
	clientFullTournament := &volleynet.FullTournament{}

	clientMock, volleynetMock, syncService := syncMock()

	gender := "M"
	league := "AMATEUR LEAGUE"
	season := 2018

	clientMock.On("AllTournaments", gender, league, season).Return(clientTournaments, nil)
	volleynetMock.On("SeasonTournaments", season).Return(persistedTournaments, nil)
	clientMock.On("ComplementTournament", clientTournaments[0]).Return(clientFullTournament, nil)
	volleynetMock.On("UpdateTournament", &volleynet.FullTournament{
		Tournament: volleynet.Tournament{ID: 1},
	}).Return(nil)
	volleynetMock.On("TournamentTeams", 1).Return([]volleynet.TournamentTeam{}, nil)
	volleynetMock.On("AllPlayers").Return([]volleynet.Player{}, nil)

	report, err := syncService.Tournaments("M", "AMATEUR LEAGUE", 2018)

	if err != nil {
		t.Error(err)
	}

	if report.UpdatedTournaments != 1 {
		t.Errorf("SyncService.Tournaments(\"M\") want: .UpdatedTournaments = 1, got: %d", report.UpdatedTournaments)
	}
}
