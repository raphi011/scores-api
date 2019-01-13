package sql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores/volleynet"
)

func TestCreateTournament(t *testing.T) {
	db := setupDB(t)
	tournamentRepo :=  &TournamentRepository{DB: db}

	tournament := &volleynet.FullTournament{
		Tournament: volleynet.Tournament{
		 ID: 1,
		},
		Teams: []volleynet.TournamentTeam{},
	}
	err := tournamentRepo.New(tournament)

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
