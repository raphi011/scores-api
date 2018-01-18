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
	statisticsSelectSQL = `
		SELECT 
			s.id,
			u.profile_image_url
			max(s.name)
			cast((sum(s.won) / cast(count(1) as float) * 100) as int) as percentage,
			sum(s.pointsWon),
			sum(s.pointsLost),
			count(1),
			sum(s.won),
			(sum(1) - sum(s.won))
		FROM playerStatistics s
		GROUP BY s.id 
		JOIN players p ON s.id = p.id
		LEFT JOIN users u ON p.user_id = u.id 
		WHERE s.created_at > $1
		ORDER BY percentage DESC
	`
)

func (s *StatisticService) Players(filter string) (scores.PlayerStatistics, error) {
	statistics := scores.PlayerStatistics{}

	timeFilter := time.Now()

	switch filter {
	case "week":
		timeFilter = timeFilter.AddDate(0, 0, -7)
	case "month":
		timeFilter = timeFilter.AddDate(0, -1, 0)
	case "quarter":
		timeFilter = timeFilter.AddDate(0, -3, 0)
	case "year":
		timeFilter = timeFilter.AddDate(-1, 0, 0)
	default: // "all"
		timeFilter = time.Unix(0, 0)
	}

	rows, err := s.DB.Query(statisticsSelectSQL, timeFilter)

	for rows.Next() {
		s := scores.PlayerStatistic{
			Player: &scores.Player{},
		}

		err = rows.Scan(
			&s.PlayerID,
			&s.Player.ProfileImageURL,
			&s.Player.Name,
			&s.PercentageWon,
			&s.PointsWon,
			&s.PointsLost,
			&s.GamesWon,
			&s.GamesLost,
		)

		if err != nil {
			return nil, err
		}

		s.Player.ID = s.PlayerID

		statistics = append(statistics, s)
	}

	return statistics, nil
}

func (s *StatisticService) Player(playerID uint) (*scores.PlayerStatistic, error) {
	return nil, nil
}

func (s *StatisticService) Team(teamID uint) (*scores.TeamStatistic, error) {

	return nil, nil
}

func (s *StatisticService) Teams() (scores.TeamStatistics, error) {

	return nil, nil
}

// func teamStatisticsQuery(db *gorm.DB) *gorm.DB {
// 	query :=
// 		db.Table("teamStatistics").Select(`
// 			team_id,
// 			max(teamStatistics.name) as pname,
// 			cast((sum(teamStatistics.won) / cast(count(1) as float) * 100) as int) as percentage,
// 			sum(teamStatistics.pointsWon) as wonpoints,
// 			sum(teamStatistics.pointsLost) as lost,
// 			count(1) as played,
// 			sum(teamStatistics.won) as wongames,
// 			(sum(1) - sum(teamStatistics.won)) as lostgames
// 		`).
// 			Group("team_id")
// 		// Joins("left join users on users.player_id = playerStatistics.id")

// 	return query
// }
