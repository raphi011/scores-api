package sqlite

import "testing"

func TestPlayerStatistic(t *testing.T) {
	s := createServices(t)
	defer Reset(s.db)

	m := newMatch(s)
	_, _ = s.matchService.Create(m)

	filter := "all"
	st, err := s.statisticService.Players(filter)

	if err != nil {
		t.Errorf("StatisticService.Players() err: %s", err)
		return
	}

	len := len(st)
	if len != 4 {
		t.Errorf("StatisticService.Players() want len(p): 4, got %d ", len)
	}

	_, err = s.statisticService.Player(m.Team1.Player1ID, filter)

	if err != nil {
		t.Errorf("StatisticService.Player() err: %s", err)
	}
}
