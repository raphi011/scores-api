package volleynet

import (
	"testing"
)

func syncMock() (*ClientMock, *VolleynetServiceMock, *SyncService) {
	clientMock := new(ClientMock)
	volleynetMock := new(VolleynetServiceMock)

	syncService := &SyncService{
		Client:           clientMock,
		VolleynetService: volleynetMock,
	}

	return clientMock, volleynetMock, syncService
}

func TestSyncLadder(t *testing.T) {
	persistedPlayers := []Player{Player{PlayerInfo: PlayerInfo{ID: 1}, TotalPoints: 100, Rank: 96}}
	clientPlayers := []Player{Player{PlayerInfo: PlayerInfo{ID: 1}, TotalPoints: 125, Rank: 60}}

	clientMock, volleynetMock, syncService := syncMock()

	clientMock.On("Ladder", "M").Return(clientPlayers, nil)
	volleynetMock.On("AllPlayers").Return(persistedPlayers, nil)
	volleynetMock.On("UpdatePlayer", &Player{
		PlayerInfo:  PlayerInfo{ID: 1},
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

func TestSyncTournaments(t *testing.T) {
	clientTournaments := []Tournament{Tournament{
		ID:     1,
		Status: "upcoming",
	}}
	persistedTournaments := []FullTournament{
		FullTournament{
			Tournament: Tournament{ID: 1, Status: "upcoming"}},
	}
	clientFullTournament := &FullTournament{}

	clientMock, volleynetMock, syncService := syncMock()

	gender := "M"
	league := "AMATEUR LEAGUE"
	season := 2018

	clientMock.On("AllTournaments", gender, league, season).Return(clientTournaments, nil)
	volleynetMock.On("SeasonTournaments", season).Return(persistedTournaments, nil)
	clientMock.On("ComplementTournament", clientTournaments[0]).Return(clientFullTournament, nil)
	volleynetMock.On("UpdateTournament", &FullTournament{
		Tournament: Tournament{ID: 1},
	}).Return(nil)
	volleynetMock.On("TournamentTeams", 1).Return([]TournamentTeam{}, nil)
	volleynetMock.On("AllPlayers").Return([]Player{}, nil)

	report, err := syncService.Tournaments("M", "AMATEUR LEAGUE", 2018)

	if err != nil {
		t.Error(err)
	}

	if report.UpdatedTournaments != 1 {
		t.Errorf("SyncService.Tournaments(\"M\") want: .UpdatedTournaments = 1, got: %d", report.UpdatedTournaments)
	}
}
