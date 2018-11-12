package sqlite

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores"
)

func TestPlayerStatistic(t *testing.T) {
	expect := []scores.PlayerStatistic{
		scores.PlayerStatistic{
			Player:   &scores.Player{Name: "p1", Model: scores.Model{ID: 1}},
			PlayerID: 1,
			Statistic: scores.Statistic{
				Played:        1,
				GamesWon:      1,
				GamesLost:     0,
				PointsWon:     15,
				PointsLost:    13,
				PercentageWon: 100,
				Rank:          "Hacker",
			},
		},
		scores.PlayerStatistic{
			Player:   &scores.Player{Name: "p2", Model: scores.Model{ID: 2}},
			PlayerID: 2,
			Statistic: scores.Statistic{
				Played:        1,
				GamesWon:      1,
				GamesLost:     0,
				PointsWon:     15,
				PointsLost:    13,
				PercentageWon: 100,
				Rank:          "Hacker",
			},
		},
		scores.PlayerStatistic{
			Player:   &scores.Player{Name: "p3", Model: scores.Model{ID: 3}},
			PlayerID: 3,
			Statistic: scores.Statistic{
				Played:        1,
				GamesWon:      0,
				GamesLost:     1,
				PointsWon:     13,
				PointsLost:    15,
				PercentageWon: 0,
				Rank:          "Curling doper",
			},
		},
		scores.PlayerStatistic{
			Player:   &scores.Player{Name: "p4", Model: scores.Model{ID: 4}},
			PlayerID: 4,
			Statistic: scores.Statistic{
				Played:        1,
				GamesWon:      0,
				GamesLost:     1,
				PointsWon:     13,
				PointsLost:    15,
				PercentageWon: 0,
				Rank:          "Curling doper",
			},
		},
	}
	s := createRepositories(t)

	m := newMatch(s)
	m, _ = s.Match.Create(m)

	filter := "all"
	st, err := s.Statistic.Players(filter)

	if err != nil {
		t.Fatalf("StatisticRepository.Players() err: %s", err)
	}

	if !cmp.Equal(st, expect) {
		t.Fatalf("StatisticRepository.Players()\n%s", cmp.Diff(st, expect))
	}

	len := len(st)
	if len != 4 {
		t.Fatalf("StatisticRepository.Players() want len(p): 4, got %d ", len)
	}

	ps, err := s.Statistic.Player(m.Team1.Player1ID, filter)

	if err != nil {
		t.Fatalf("StatisticRepository.Player() err: %s", err)
	}

	if !cmp.Equal(ps, &expect[0]) {
		t.Fatalf("StatisticRepository.Player()\n%s", cmp.Diff(ps, &expect[0]))
	}
}
