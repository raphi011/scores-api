package sql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores/volleynet"
)

func TestCreateTournament(t *testing.T) {
	db := setupDB(t)
	tournamentRepo :=  &TournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.FullTournament{
		Tournament: volleynet.Tournament{
		 ID: 1,
		},
		Teams: []volleynet.TournamentTeam{},
	})

	if err != nil {
		t.Fatalf("tournamentRepository.New(), err: %s", err)
	}

	persistedTournament, err := tournamentRepo.Get(tournament.ID)

	if err != nil {
		t.Fatalf("tournamentRepository.Get(), err: %s", err)
	}

	if !cmp.Equal(tournament, persistedTournament) {
		t.Fatalf("tournaments are not equal:\n%s", cmp.Diff(tournament, persistedTournament))
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
		Teams: []volleynet.TournamentTeam{},
	}

	tournament2 := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			ID: 2,
			Season: 2018,
			League: "amateur-tour",
			Format: "m",
		},
		Teams: []volleynet.TournamentTeam{},
	}

	tournament3 := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			ID: 3,
			Season: 2018,
			League: "pro-tour",
			Format: "m",
		},
		Teams: []volleynet.TournamentTeam{},
	}

	tournament4 := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
			ID: 3,
			Season: 2017,
			League: "amateur-tour",
			Format: "m",
		},
		Teams: []volleynet.TournamentTeam{},
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

func TestUpdateTournament(t *testing.T) {
	db := setupDB(t)
	tournamentRepo := &TournamentRepository{DB: db}

	tournament, err := tournamentRepo.New(&volleynet.FullTournament{
		Tournament: volleynet.Tournament{ ID: 1 },
		Teams: []volleynet.TournamentTeam{},
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
