package sqlite

import (
	"testing"
)

func TestGetMatches(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	t.Skip()
}

func TestCreateMatch(t *testing.T) {
	s := createServices()
	defer ClearTables(s.db)

	match := newMatch(s)

	match, err := s.matchService.Create(match)

	if err != nil {
		t.Error("Can't create match", err)
	} else if match.ID == 0 {
		t.Error("MatchID not assigned")
	}
}

func TestDeleteMatch(t *testing.T) {
	s := createServices()
	defer ClearTables(s.db)

	match := newMatch(s)
	match, _ = s.matchService.Create(match)

	err := s.matchService.Delete(match.ID)

	if err != nil {
		t.Errorf("MatchService.Delete() err: %s", err)
	}

	match, err = s.matchService.Match(match.ID)

	if err == nil {
		t.Errorf("MatchService.Delete() err: %s", err)
	}
}
