package sqlite

import (
	"testing"
)

func TestCreateMatch(t *testing.T) {
	s := createServices()
	defer Reset(s.db)

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
	defer Reset(s.db)

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
