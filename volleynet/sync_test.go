package volleynet

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSyncLadder(t *testing.T) {
	clientPlayers := []Player{Player{PlayerInfo: PlayerInfo{ID: 1}, TotalPoints: 100}}
	persistedPlayers := []Player{Player{PlayerInfo: PlayerInfo{ID: 1}, TotalPoints: 125}}

	clientMock := new(ClientMock)
	volleynetMock := new(VolleynetServiceMock)

	syncService := SyncService{
		Client:           clientMock,
		VolleynetService: volleynetMock,
	}

	clientMock.On("Ladder", "M").Return(clientPlayers, nil)
	volleynetMock.On("AllPlayers").Return(persistedPlayers, nil)
	volleynetMock.On("UpdatePlayer", mock.Anything).Return(nil)

	report, err := syncService.Ladder("M")

	if err != nil {
		t.Error(err)
	}

	if report.UpdatedPlayers != 1 {
		t.Errorf("SyncService.Ladder(\"M\") want: .UpdatedPlayers = 1, got: %d", report.UpdatedPlayers)
	}
}
