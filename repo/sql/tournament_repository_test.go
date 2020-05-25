// +build repository

package sql

import (
	"math/rand"
	"testing"
	"time"

	"github.com/raphi011/scores-backend/repo"
	"github.com/raphi011/scores-backend/test"

	"github.com/raphi011/scores-backend/volleynet"
	"github.com/wawandco/fako"
)

func TestSeasons(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	expected := []string{"2017", "2018"}

	CreateTournaments(t, db,
		T{ID: 1, Season: "2017"},
		T{ID: 2, Season: "2018"},
		T{ID: 3, Season: "2018"},
	)

	actual, err := tournamentRepo.Seasons()

	test.Check(t, "tournamentRepository.Seasons(), err: %v", err)
	test.Compare(t, "Seasons are not equal:\n%s", expected, actual)
}

func TestSubLeagues(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	CreateTournaments(t, db,
		T{ID: 1, SubLeague: "Amateur Tour 1"},
		T{ID: 2, SubLeague: "Amateur Tour 2"},
		T{ID: 3, SubLeague: "LMS"},
		T{ID: 4, SubLeague: "Pro Tour 80"},
	)

	actual, err := tournamentRepo.SubLeagues()

	test.Check(t, "tournamentRepository.SubLeagues(), err: %v", err)
	test.Assert(t, "SubLeagues() expected len() == 4, got: %d", len(actual) == 4, len(actual))
}
func TestLeagues(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	CreateTournaments(t, db,
		T{ID: 1, League: "Amateur Tour"},
		T{ID: 2, League: "Amateur Tour"},
		T{ID: 3, League: "Pro Tour"},
		T{ID: 4, League: "Amateur Tour"},
	)

	actual, err := tournamentRepo.Leagues()

	test.Check(t, "tournamentRepository.Leagues(), err: %v", err)
	test.Assert(t, "Leagues() expected len() == 2, got: %d", len(actual) == 2, len(actual))
}

func TestCreateTournament(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			ID:    1,
			Start: time.Now(),
			End:   time.Now(),
		},
		Teams: []*volleynet.TournamentTeam{},
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

	tournaments := []struct {
		ID        int
		Season    string
		LeagueKey string
		Gender    string
	}{
		{
			ID:        1,
			Season:    "2018",
			LeagueKey: "amateur-tour",
			Gender:    "M",
		},
		{
			ID:        2,
			Season:    "2018",
			LeagueKey: "amateur-tour",
			Gender:    "M",
		},
		{
			ID:        3,
			Season:    "2018",
			LeagueKey: "pro-tour",
			Gender:    "M",
		},
		{
			ID:        4,
			Season:    "2017",
			LeagueKey: "amateur-tour",
			Gender:    "W",
		},
		{
			ID:        5,
			Season:    "2017",
			LeagueKey: "junior-tour",
			Gender:    "M",
		},
	}

	for _, tournament := range tournaments {
		_, err := tournamentRepo.New(&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
				ID:        tournament.ID,
				Season:    tournament.Season,
				LeagueKey: tournament.LeagueKey,
				Gender:    tournament.Gender,
				Start:     time.Now(),
				End:       time.Now(),
			},
			Teams: []*volleynet.TournamentTeam{},
		})

		test.Check(t, "tournamentRepo.New() failed: %v", err)
	}

	got, err := tournamentRepo.Search(repo.TournamentFilter{
		Seasons: []string{"2018"},
		Leagues: []string{"amateur-tour", "pro-tour"},
		Genders: []string{"M"},
	})

	test.Check(t, "tournamentRepository.Search(), err: %s", err)
	test.Assert(t, "tournamentRepository.Search(), want len(tournaments) 3, got %d", len(got) == 3, len(got))
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
		ts, err := tournamentRepo.Search(repo.TournamentFilter{
			Seasons: []string{"2018"},
			Leagues: []string{"amateur-tour"},
			Genders: []string{"M"},
		})

		b.Logf("found %d tournaments", len(ts))

		test.Check(b, "failed to filter tournaemnts: %v", err)
	}
}

func TestUpdateTournament(t *testing.T) {
	db := SetupDB(t)
	tournamentRepo := &tournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.Tournament{
		TournamentInfo: volleynet.TournamentInfo{
			ID:    1,
			Start: time.Now(),
			End:   time.Now(),
		},
		Teams: []*volleynet.TournamentTeam{},
	})
	test.Check(t, "couldn't persist tournament: %v", err)

	tournament.Email = "test!"

	err = tournamentRepo.Update(tournament)
	test.Check(t, "couldnt update tournament: %v", err)

	updatedTournament, err := tournamentRepo.Get(tournament.ID)
	test.Check(t, "couldnt get tournament: %v", err)
	test.Compare(t, "tournaments are not equal:\n%s", tournament, updatedTournament)
}

func randomTournaments(count, run int) []*volleynet.Tournament {
	tournaments := make([]*volleynet.Tournament, count)
	rand.Seed(time.Now().Unix())

	leagues := []string{"amateur-tour", "pro-tour", "junior-tour"}
	seasons := []string{"2017", "2018", "2019"}
	genders := []string{"M", "W"}
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
		tournament.LeagueKey = leagues[rand.Intn(len(leagues))]
		tournament.Season = seasons[rand.Intn(len(seasons))]
		tournament.Gender = genders[rand.Intn(len(genders))]
		tournament.Status = status[rand.Intn(len(status))]
		tournament.Start = time.Now()
		tournament.End = time.Now()

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
