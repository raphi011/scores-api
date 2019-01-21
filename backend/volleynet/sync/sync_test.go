package sync

import (
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores/events"
	"github.com/raphi011/scores/repo/sql"
	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/volleynet/mocks"
	"github.com/raphi011/scores/volleynet/scrape"
)

func syncMock(t *testing.T) (*mocks.ClientMock, *Service, *sqlx.DB) {
	repos, db := sql.RepositoriesTest(t)
	clientMock := new(mocks.ClientMock)

	service := &Service{
		Client:         clientMock,
		PlayerRepo:     repos.PlayerRepo,
		TournamentRepo: repos.TournamentRepo,
		TeamRepo:       repos.TeamRepo,
		Subscriptions:  &events.Broker{},
	}

	return clientMock, service, db
}

func TestSyncLadder(t *testing.T) {
	clientMock, service, db := syncMock(t)
	gender := "M"

	sql.CreatePlayers(t, db,
		sql.P{ID: 1, TotalPoints: 100, LadderRank: 96, Gender: gender},
	)

	clientPlayers := []*volleynet.Player{
		&volleynet.Player{ID: 1, TotalPoints: 125, LadderRank: 60, Gender: gender},
	}

	clientMock.On("Ladder", gender).Return(clientPlayers, nil)

	report, err := service.Ladder(gender)

	test.Check(t, "service.Ladder() err: %v", err)
	test.Assert(t, "Service.Ladder(\"M\") want: .UpdatedPlayers = 1, got: %d", report.UpdatedPlayers == 1, report.UpdatedPlayers)
}

func TestSyncTournamentInformation(t *testing.T) {
	response, _ := os.Open("../testdata/upcoming.html")
	tournament, _ := scrape.Tournament(response, time.Now(), &volleynet.TournamentInfo{Status: volleynet.StatusUpcoming, ID: 22231})

	syncInfos := Tournaments(tournament, &volleynet.TournamentInfo{ID: 22231, Status: volleynet.StatusUpcoming})

	if syncInfos.Type != SyncTournamentUpcoming {
		t.Fatalf("Service.Tournaments() want: %s, got: %s", SyncTournamentUpcoming, syncInfos.Type)
	}
}

func TestSyncTournaments(t *testing.T) {
	clientMock, service, db := syncMock(t)

	clientTournaments := []*volleynet.TournamentInfo{&volleynet.TournamentInfo{
		ID:     1,
		Status: volleynet.StatusUpcoming,
		Start:  time.Now(),
		End:    time.Now(),
	}}

	clientFullTournament := []*volleynet.Tournament{
		&volleynet.Tournament{TournamentInfo: volleynet.TournamentInfo{
			ID:     1,
			Status: volleynet.StatusUpcoming,
			Name:   "New name",
			Start:  time.Now(),
			End:    time.Now(),
		},
			Teams: []*volleynet.TournamentTeam{},
		},
	}

	sql.CreateTournaments(t, db,
		sql.T{ID: 1, Status: volleynet.StatusUpcoming},
	)

	gender := "M"
	league := "AMATEUR LEAGUE"
	season := 2018

	clientMock.On("AllTournaments", gender, league, season).Return(clientTournaments, nil)
	clientMock.On("ComplementMultipleTournaments", clientTournaments).Return(clientFullTournament, nil)

	err := service.Tournaments("M", "AMATEUR LEAGUE", 2018)

	test.Check(t, "service.Tournaments() err: %v", err)
}
