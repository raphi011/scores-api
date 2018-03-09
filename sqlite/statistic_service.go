package sqlite

import (
	"database/sql"
	"time"

	"github.com/raphi011/scores"
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
			COALESCE(u.profile_image_url, "") as profileImage,
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

	err := scanner.Scan(
		&s.PlayerID,
		&s.Player.ProfileImageURL,
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

	s.Rank = scores.CalculateRank(int(s.PercentageWon))
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

const (
	playerTeamsStatisticSelectSQL = `
		SELECT 
			MAX(CASE WHEN s.player1_id = $1 THEN s.player2_id ELSE s.player1_id END) AS player_id,
			COALESCE(MAX(CASE WHEN s.player1_id = $1 THEN u2.profile_image_url ELSE u1.profile_image_url END), "") AS profileImage,
			MAX(CASE WHEN s.player1_id = $1 THEN p2.name ELSE p1.name END) AS name,
			CAST((SUM(s.won) / CAST(COUNT(1) AS float) * 100) AS int) AS percentageWon,
			SUM(s.pointsWon) AS pointsWon,
			SUM(s.pointsLost) AS pointsLost,
			COUNT(1) AS played,
			SUM(s.won) AS gamesWon,
			SUM(1) - SUM(s.won) AS gamesLost
		FROM teamStatistics s
		JOIN players p1 ON s.player1_id = p1.id
		JOIN players p2 ON s.player2_id = p2.id
		LEFT JOIN users u1 ON p1.user_id = u1.id 
		LEFT JOIN users u2 ON p2.user_id = u2.id 
		WHERE (s.player1_id = $1 OR s.player2_id = $1) and s.created_at > $2
		GROUP BY s.player1_id, s.player2_id 
		ORDER BY percentageWon DESC
	`
)

func (s *StatisticService) PlayerTeams(playerID uint, filter string) (scores.PlayerStatistics, error) {
	statistics := scores.PlayerStatistics{}

	timeFilter := parseTimeFilter(filter)

	rows, err := s.DB.Query(playerTeamsStatisticSelectSQL, playerID, timeFilter)

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
