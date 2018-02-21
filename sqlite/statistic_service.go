package sqlite

import (
	"database/sql"
	scores "scores-backend"
	"time"
)

var _ scores.StatisticService = &StatisticService{}

type StatisticService struct {
	DB *sql.DB
}

const (
	statisticFieldsSelectSQL = `
			max(s.name) as name,
			cast((sum(s.won) / cast(count(1) as float) * 100) as int) as percentageWon,
			sum(s.pointsWon) as pointsWon,
			sum(s.pointsLost) as pointsLost,
			count(1) as played,
			sum(s.won) as gamesWon,
			sum(1) - sum(s.won) as gamesLost
	`
	ungroupedPlayerStatisticSelectSQL = `
		SELECT 
			s.player_id,
			u.profile_image_url as profileImage,
	` + statisticFieldsSelectSQL + `
		FROM playerStatistics s
		JOIN players p ON s.player_id = p.id
		LEFT JOIN users u ON p.user_id = u.id 
		WHERE s.created_at > $1
	`
	groupedPlayerStatisticSQL = `
		GROUP BY s.player_id 
		ORDER BY percentageWon DESC
	`
	playersStatisticSelectSQL = ungroupedPlayerStatisticSelectSQL + groupedPlayerStatisticSQL

	playerStatisticSelectSQL = ungroupedPlayerStatisticSelectSQL +
		" and s.player_id = $2 " + groupedPlayerStatisticSQL
)

func parseTimeFilter(filter string) time.Time {
	timeFilter := time.Now()
	year := timeFilter.Year()
	month := timeFilter.Month()
	day := timeFilter.Day()
	loc := timeFilter.Location()

	switch filter {
	case "today":
		timeFilter = time.Date(year, month, day, 0, 0, 0, 0, loc)
	case "month":
		timeFilter = time.Date(year, month-1, day, 0, 0, 0, 0, loc)
	case "thisyear":
		timeFilter = time.Date(year, 1, 1, 0, 0, 0, 0, loc)
	default: // "all"
		timeFilter = time.Unix(0, 0)
	}

	return timeFilter
}

func scanPlayerStatistic(scanner scan) (*scores.PlayerStatistic, error) {
	s := &scores.PlayerStatistic{
		Player: &scores.Player{},
	}

	var profileImageURL sql.NullString

	err := scanner.Scan(
		&s.PlayerID,
		&profileImageURL,
		&s.Player.Name,
		&s.PercentageWon,
		&s.PointsWon,
		&s.PointsLost,
		&s.Played,
		&s.GamesWon,
		&s.GamesLost,
	)

	if err != nil {
		return nil, err
	}

	if profileImageURL.Valid {
		s.Player.ProfileImageURL = profileImageURL.String
	}

	s.Player.ID = s.PlayerID

	return s, nil
}

func (s *StatisticService) Players(filter string) (scores.PlayerStatistics, error) {
	statistics := scores.PlayerStatistics{}

	timeFilter := parseTimeFilter(filter)

	rows, err := s.DB.Query(playersStatisticSelectSQL, timeFilter)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		st, err := scanPlayerStatistic(rows)

		if err != nil {
			return nil, err
		}

		statistics = append(statistics, *st)
	}

	return statistics, nil
}

func (s *StatisticService) Player(playerID uint, filter string) (*scores.PlayerStatistic, error) {
	timeFilter := parseTimeFilter(filter)

	row := s.DB.QueryRow(playerStatisticSelectSQL, timeFilter, playerID)

	st, err := scanPlayerStatistic(row)

	if err != nil {
		return nil, err
	}

	return st, nil
}

func (s *StatisticService) Team(teamID uint) (*scores.TeamStatistic, error) {

	return nil, nil
}

func (s *StatisticService) Teams() (scores.TeamStatistics, error) {

	return nil, nil
}
