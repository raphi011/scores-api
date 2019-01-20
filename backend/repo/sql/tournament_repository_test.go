package sql

import (
	"math/rand"
	"testing"
	"time"

	"github.com/raphi011/scores/test"

	"github.com/raphi011/scores/volleynet"
	"github.com/wawandco/fako"
)

func TestCreateTournament(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{ID: 1},
		Teams:          []*volleynet.TournamentTeam{},
	})

	if err != nil {
		t.Fatalf("tournamentRepository.New(), err: %s", err)
	}

	persistedTournament, err := tournamentRepo.Get(tournament.ID)

	test.Check(t, "tournamentRepository.Get(), err: %v", err)
	test.Compare(t, "tournaments are not equal:\n%s", tournament, persistedTournament)
}

func TestFilterTournament(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	tournament1 := &volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			ID:     1,
			Season: 2018,
			League: "amateur-tour",
			Gender: "m",
		},
		Teams: []*volleynet.TournamentTeam{},
	}

	tournament2 := &volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			ID:     2,
			Season: 2018,
			League: "amateur-tour",
			Gender: "m",
		},
		Teams: []*volleynet.TournamentTeam{},
	}

	tournament3 := &volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			ID:     3,
			Season: 2018,
			League: "pro-tour",
			Gender: "m",
		},
		Teams: []*volleynet.TournamentTeam{},
	}

	tournament4 := &volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			ID:     3,
			Season: 2017,
			League: "amateur-tour",
			Gender: "m",
		},
		Teams: []*volleynet.TournamentTeam{},
	}

	tournamentRepo.New(tournament1)
	tournamentRepo.New(tournament2)
	tournamentRepo.New(tournament3)
	tournamentRepo.New(tournament4)

	tournaments, err := tournamentRepo.Filter(
		[]int{2018}, []string{"amateur-tour", "pro-tour"}, []string{"m"},
	)

	test.Check(t, "tournamentRepository.Filter(), err: %s", err)
	test.Assert(t, "tournamentRepository.Filter(), want len(tournaments) 3, got %d", len(tournaments) == 3, len(tournaments))
}

func BenchmarkCreateTournament(b *testing.B) {
	db := SetupDB(b)
	tournamentRepo := &tournamentRepository{DB: db}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		tournaments := randomTournaments(1000, n)
		b.StartTimer()

		err := tournamentRepo.NewBatch(tournaments...)
		test.Check(b, "failed to create random tournaments: %v", err)
	}
}

func BenchmarkFilterTournament(b *testing.B) {
	b.StopTimer()
	db := SetupDB(b)
	tournamentRepo := &tournamentRepository{DB: db}

	tournaments := randomTournaments(1000, 0)
	err := tournamentRepo.NewBatch(tournaments...)
	test.Check(b, "failed to create random tournaments: %v", err)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		ts, err := tournamentRepo.Filter(
			[]int{2018},
			[]string{"amateur-tour"},
			[]string{"m"},
		)

		b.Logf("found %d tournaments", len(ts))

		test.Check(b, "failed to filter tournaemnts: %v", err)
	}
}

func TestUpdateTournament(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{ID: 1},
		Teams:          []*volleynet.TournamentTeam{},
	})
	test.Check(t, "couldn't persist tournament: %v", err)

	tournament.Email = "test!"

	err = tournamentRepo.Update(tournament)
	test.Check(t, "couldnt update tournament: %v", err)

	updatedTournament, err := tournamentRepo.Get(tournament.ID)
	test.Check(t, "couldnt get tournament: %v", err)
	test.Compare(t,"tournaments are not equal:\n%s", tournament, updatedTournament)
}

func randomTournaments(count, run int) []*volleynet.Tournament {
	tournaments := make([]*volleynet.Tournament, count)
	rand.Seed(time.Now().Unix())

	leagues := []string{"amateur-tour", "pro-tour", "junior-tour"}
	seasons := []int{2017, 2018, 2019}
	genders := []string{"m", "w"}
	status := []string{
		volleynet.StatusUpcoming,
		volleynet.StatusDone,
		volleynet.StatusCanceled,
	}

	id := run * count

	for i := range tournaments {
		id++

		tournament := &volleynet.TournamentInfo{}
		tournament.ID = id
		tournament.League = leagues[rand.Intn(len(leagues))]
		tournament.Season = seasons[rand.Intn(len(seasons))]
		tournament.Gender = genders[rand.Intn(len(genders))]
		tournament.Status = status[rand.Intn(len(status))]

		fako.Fill(tournament)

		fullTournament := &volleynet.Tournament{}
		fullTournament.SignedupTeams = rand.Intn(32)
		fullTournament.MaxTeams = rand.Intn(32)
		fullTournament.MinTeams = rand.Intn(16)
		fullTournament.MaxPoints = rand.Intn(80)
		fullTournament.CreatedAt = time.Now()

		fako.Fill(fullTournament)
		fullTournament.TournamentInfo = *tournament
		tournaments[i] = fullTournament
	}

	return tournaments
}
