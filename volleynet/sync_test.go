package volleynet

import (
	"testing"
)

func TestSyncLadder(t *testing.T) {
	clientMock := new(ClientMock)
	volleynetMock := new(VolleynetServiceMock)

	syncService := SyncService{
		Client:           clientMock,
		VolleynetService: volleynetMock,
	}

	clientMock.On("Ladder", "M").Return([]Player{}, nil)
	volleynetMock.On("AllPlayers").Return([]Player{}, nil)

	_, err := syncService.Ladder("M")

	if err != nil {
		t.Error(err)
	}
}
