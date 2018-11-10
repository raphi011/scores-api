package sqlite

import (
	"testing"
)

func TestCreateMatch(t *testing.T) {
	s := createRepositories(t)

	match := newMatch(s)

	match, err := s.matchRepository.Create(match)

	if err != nil {
		t.Error("Can't create match", err)
	} else if match.ID == 0 {
		t.Error("MatchID not assigned")
	}
}

func TestDeleteMatch(t *testing.T) {
	s := createRepositories(t)

	match := newMatch(s)
	match, _ = s.matchRepository.Create(match)

	err := s.matchRepository.Delete(match.ID)

	if err != nil {
		t.Errorf("MatchRepository.Delete() err: %s", err)
	}

	match, err = s.matchRepository.Match(match.ID)

	if err == nil {
		t.Errorf("MatchRepository.Delete() err: %s", err)
	}
}
