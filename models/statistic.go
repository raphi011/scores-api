package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Statistic struct {
	ID           int     `json:"playerId"`
	Pname        string  `json:"name"`
	Played       int     `json:"played"`
	Wongames     int     `json:"gamesWon"`
	Lostgames    int     `json:"gamesLost"`
	Wonpoints    int     `json:"pointsWon"`
	Lost         int     `json:"pointsLost"`
	Percentage   float32 `json:"percentageWon"`
	Profileimage string  `json:"profileImage"`
}

func statisticsQuery(db *gorm.DB) *gorm.DB {
	query :=
		db.Table("playerStatistics").Select(`
			playerStatistics.id,
			users.profile_image_url as profileimage,
			max(playerStatistics.name) as pname,
			cast((sum(playerStatistics.won) / cast(count(1) as float) * 100) as int) as percentage,
			sum(playerStatistics.pointsWon) as wonpoints,
			sum(playerStatistics.pointsLost) as lost,
			count(1) as played,
			sum(playerStatistics.won) as wongames,
			(sum(1) - sum(playerStatistics.won)) as lostgames
		`).
			Group("playerStatistics.id").
			Joins("left join users on users.player_id = playerStatistics.id")

	return query
}

func GetStatistics(db *gorm.DB, filter string) []Statistic {
	var statistics []Statistic

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

	statisticsQuery(db).
		Where("playerStatistics.created_at > ?", timeFilter).
		Order("percentage desc").
		Scan(&statistics)

	return statistics
}

func (s *Statistic) GetStatistic(db *gorm.DB, playerID uint) {
	statisticsQuery(db).Where("playerStatistics.id = ?", playerID).First(&s)
}
