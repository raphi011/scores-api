package sync

import (
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/raphi011/scores-backend/events"
	"github.com/raphi011/scores-backend/repo/sql"
	"github.com/raphi011/scores-backend/test"
	"github.com/raphi011/scores-backend/volleynet"
	"github.com/raphi011/scores-backend/volleynet/mocks"
	"github.com/raphi011/scores-backend/volleynet/scrape"
)

func syncMock(t *testing.T) (*mocks.ClientMock, *Service, *sqlx.DB) {
	repos, db := sql.RepositoriesTest(t)
	clientMock := new(mocks.ClientMock)

	service := &Service{
		Log: log.New(),

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
		{ID: 1, TotalPoints: 125, LadderRank: 60, Gender: gender},
	}

	clientMock.On("Ladder", gender).Return(clientPlayers, nil)

	report, err := service.Ladder(gender)

	test.Check(t, "service.Ladder() err: %v", err)
	test.Assert(t, "Service.Ladder(\"M\") want: .UpdatedPlayers = 1, got: %d", report.UpdatedPlayers == 1, report.UpdatedPlayers)
}

func TestSyncTournamentInformation(t *testing.T) {
	response, _ := os.Open("../testdata/upcoming.html")
	tournament, _ := scrape.Tournament(response, test.MustParseDate("30.05.2018"), &volleynet.TournamentInfo{Status: volleynet.StatusUpcoming, ID: 22231})

	syncInfos := Tournaments(tournament, &volleynet.TournamentInfo{ID: 22231, Status: volleynet.StatusUpcoming})

	if syncInfos.Type != SyncTournamentUpcoming {
		t.Fatalf("Service.Tournaments() want: %s, got: %s", SyncTournamentUpcoming, syncInfos.Type)
	}
}

func TestSyncTournaments(t *testing.T) {
	clientMock, service, db := syncMock(t)

	clientTournaments := []*volleynet.TournamentInfo{{
		ID:     1,
		Status: volleynet.StatusUpcoming,
		Start:  time.Now(),
		End:    time.Now(),
	}}

	clientFullTournament := []*volleynet.Tournament{
		{TournamentInfo: volleynet.TournamentInfo{
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
	league := "amateur-league"
	season := 2018

	clientMock.On("Tournaments", gender, league, season).Return(clientTournaments, nil)
	clientMock.On("ComplementTournament", clientTournaments[0]).Return(clientFullTournament[0], nil)

	err := service.Tournaments("M", "amateur-league", 2018)

	test.Check(t, "service.Tournaments() err: %v", err)
}
