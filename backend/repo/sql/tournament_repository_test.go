package sql

import (
	"testing"
	"time"
	"math/rand"

	"github.com/wawandco/fako"
	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores/volleynet"
)

func TestCreateTournament(t *testing.T) {
	db := setupDB(t)
	tournamentRepo :=  &TournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.FullTournament{
		Tournament: volleynet.Tournament{ ID: 1 },
		Teams: []*volleynet.TournamentTeam{},
	})

	if err != nil {
		t.Fatalf("tournamentRepository.New(), err: %s", err)
	}

	persistedTournament, err := tournamentRepo.Get(tournament.ID)

	if err != nil {
		t.Fatalf("tournamentRepository.Get(), err: %s", err)
	}

	if !cmp.Equal(tournament, persistedTournament) {
		t.Fatalf(
			"tournaments are not equal:\n%s",
			cmp.Diff(tournament, persistedTournament),
		)
	}
}

func TestFilterTournament(t *testing.T) {
	db := setupDB(t)
	tournamentRepo :=  &TournamentRepository{DB: db}

	tournament1 := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			ID: 1,
			Season: 2018,
			League: "amateur-tour",
			Format: "m",
		},
		Teams: []*volleynet.TournamentTeam{},
	}

	tournament2 := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			ID: 2,
			Season: 2018,
			League: "amateur-tour",
			Format: "m",
		},
		Teams: []*volleynet.TournamentTeam{},
	}

	tournament3 := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			ID: 3,
			Season: 2018,
			League: "pro-tour",
			Format: "m",
		},
		Teams: []*volleynet.TournamentTeam{},
	}

	tournament4 := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			ID: 3,
			Season: 2017,
			League: "amateur-tour",
			Format: "m",
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

	if err != nil {
		t.Fatalf("tournamentRepository.Filter(), err: %s", err)
	}

	if len(tournaments) != 3 {
		t.Fatalf("tournamentRepository.Filter(), want len(tournaments) 3, got %d", len(tournaments))
	}
}

func BenchmarkCreateTournament(b *testing.B) {
	db := setupDB(b)
	tournamentRepo := &TournamentRepository{DB: db}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		tournaments := randomTournaments(1000, n)
		b.StartTimer()

		err := tournamentRepo.NewBatch(tournaments...)

		if err != nil {
			assert(b, "failed to create random tournaments: %v", err)
		}
	}
}

func BenchmarkFilterTournament(b *testing.B) {
	b.StopTimer()
	db := setupDB(b)
	tournamentRepo := &TournamentRepository{DB: db}

	tournaments := randomTournaments(1000, 0)
	err := tournamentRepo.NewBatch(tournaments...)
	assert(b, "failed to create random tournaments: %v", err)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		ts, err := tournamentRepo.Filter(
			[]int{2018},
			[]string{"amateur-tour"},
			[]string{"m"},
		)

		b.Logf("found %d tournaments", len(ts))

		assert(b, "failed to filter tournaemnts: %v", err)
	}
}

func TestUpdateTournament(t *testing.T) {
	db := setupDB(t)
	tournamentRepo := &TournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.FullTournament{
		Tournament: volleynet.Tournament{ ID: 1 },
		Teams: []*volleynet.TournamentTeam{},
	})
	assert(t, "couldn't persist tournament: %v", err)

	tournament.Email = "test!"

	err = tournamentRepo.Update(tournament)
	assert(t, "couldnt update tournament: %v", err)

	updatedTournament, err := tournamentRepo.Get(tournament.ID)
	assert(t, "couldnt get tournament: %v", err)

	if diff := cmp.Diff(tournament, updatedTournament); diff != "" {
		t.Fatalf("tournaments are not equal:\n%s", diff)
	}
}

func randomTournaments(count, run int) []*volleynet.FullTournament {
	tournaments := make([]*volleynet.FullTournament, count)
	rand.Seed(time.Now().Unix())

	leagues := []string{"amateur-tour", "pro-tour", "junior-tour"}
	seasons := []int{2017,2018, 2019}
	formats := []string{"m", "w"}
	status := []string{
		volleynet.StatusUpcoming,
		volleynet.StatusDone,
		volleynet.StatusCanceled,
	}

	id := run * count

	for i := range tournaments {
		id++

		tournament := &volleynet.Tournament{}
		tournament.ID = id
		tournament.League = leagues[rand.Intn(len(leagues))]
		tournament.Season  = seasons[rand.Intn(len(seasons))]
		tournament.Format = formats[rand.Intn(len(formats))]
		tournament.Status = status[rand.Intn(len(status))]

		fako.Fill(tournament)

		fullTournament := &volleynet.FullTournament{}
		fullTournament.SignedupTeams = rand.Intn(32)
		fullTournament.MaxTeams = rand.Intn(32)
		fullTournament.MinTeams = rand.Intn(16)
		fullTournament.MaxPoints = rand.Intn(80)
		fullTournament.CreatedAt = time.Now()

		fako.Fill(fullTournament)
		fullTournament.Tournament = *tournament
		tournaments[i] = fullTournament
	}

	return tournaments
}
