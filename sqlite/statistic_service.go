package sqlite

import (
	"database/sql"
	scores "scores-backend"
)

var _ scores.StatisticService = &StatisticService{}

type StatisticService struct {
	DB *sql.DB
}

// func statisticsQuery(db *gorm.DB) *gorm.DB {
// 	query :=
// 		db.Table("playerStatistics").Select(`
// 			playerStatistics.id,
// 			users.profile_image_url as profileimage,
// 			max(playerStatistics.name) as pname,
// 			cast((sum(playerStatistics.won) / cast(count(1) as float) * 100) as int) as percentage,
// 			sum(playerStatistics.pointsWon) as wonpoints,
// 			sum(playerStatistics.pointsLost) as lost,
// 			count(1) as played,
// 			sum(playerStatistics.won) as wongames,
// 			(sum(1) - sum(playerStatistics.won)) as lostgames
// 		`).
// 			Group("playerStatistics.id").
// 			Joins("left join users on users.player_id = playerStatistics.id")

// 	return query
// }

func (s *StatisticService) Players(filter string) (*scores.Statistics, error) {
	// var statistics []Statistic

	// timeFilter := time.Now()

	// switch filter {
	// case "week":
	// 	timeFilter = timeFilter.AddDate(0, 0, -7)
	// case "month":
	// 	timeFilter = timeFilter.AddDate(0, -1, 0)
	// case "quarter":
	// 	timeFilter = timeFilter.AddDate(0, -3, 0)
	// case "year":
	// 	timeFilter = timeFilter.AddDate(-1, 0, 0)
	// default: // "all"
	// 	timeFilter = time.Unix(0, 0)
	// }

	// statisticsQuery(db).
	// 	Where("playerStatistics.created_at > ?", timeFilter).
	// 	Order("percentage desc").
	// 	Scan(&statistics)

	// return statistics
	return nil, nil
}

func (s *StatisticService) Player(playerID uint) (*scores.Statistic, error) {
	// statisticsQuery(db).Where("playerStatistics.id = ?", playerID).First(&s)
	return nil, nil
}

func (s *StatisticService) Team(teamID uint) (*scores.Statistic, error) {

	return nil, nil
}

func (s *StatisticService) Teams() (*scores.Statistics, error) {

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

// func GetTeamStatistics(db *gorm.DB) []Statistic {
// 	var statistics []TeamStatistic

// 	timeFilter := time.Now()

// 	switch filter {
// 	case "week":
// 		timeFilter = timeFilter.AddDate(0, 0, -7)
// 	case "month":
// 		timeFilter = timeFilter.AddDate(0, -1, 0)
// 	case "quarter":
// 		timeFilter = timeFilter.AddDate(0, -3, 0)
// 	case "year":
// 		timeFilter = timeFilter.AddDate(-1, 0, 0)
// 	default: // "all"
// 		timeFilter = time.Unix(0, 0)
// 	}

// 	statisticsQuery(db).
// 		Where("playerStatistics.created_at > ?", timeFilter).
// 		Order("percentage desc").
// 		Scan(&statistics)

// 	return statistics
// }
