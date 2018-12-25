package sqlite

import (
	"database/sql"
	"time"

	"github.com/raphi011/scores"
)

var _ scores.StatisticRepository = &StatisticRepository{}

// StatisticRepository calculates various statistics from the repository
type StatisticRepository struct {
	DB *sql.DB
}

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

// Players loads statistics for players since a certain time
func (s *StatisticRepository) Players(filter string) ([]scores.PlayerStatistic, error) {
	timeFilter := parseTimeFilter(filter)

	return scanPlayerStatistics(s.DB, query("statistic/select-by-players"), timeFilter)
}

// PlayersByGroup loads statistics for players of a group since a certain time
func (s *StatisticRepository) PlayersByGroup(groupID uint, filter string) ([]scores.PlayerStatistic, error) {
	timeFilter := parseTimeFilter(filter)

	return scanPlayerStatistics(s.DB, query("statistic/select-by-group-player"), timeFilter, groupID)
}

// Player loads statistics for a players
func (s *StatisticRepository) Player(playerID uint, filter string) (*scores.PlayerStatistic, error) {
	timeFilter := parseTimeFilter(filter)

	row := s.DB.QueryRow(query("statistic/select-by-player"), timeFilter, playerID)

	return scanPlayerStatistic(row)
}

// PlayerTeams loads statistics of all teammates of a player since a certain time
func (s *StatisticRepository) PlayerTeams(playerID uint, filter string) ([]scores.PlayerStatistic, error) {
	timeFilter := parseTimeFilter(filter)

	return scanPlayerStatistics(s.DB, query("statistic/select-by-team"), playerID, timeFilter)
}

func scanPlayerStatistics(db *sql.DB, query string, args ...interface{}) ([]scores.PlayerStatistic, error) {
	statistics := []scores.PlayerStatistic{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		statistic, err := scanPlayerStatistic(rows)

		if err != nil {
			return nil, err
		}

		statistics = append(statistics, *statistic)
	}

	return statistics, nil
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
