package main

import (
	"scores-backend/models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	playerStatisticsView = `
		create view if not exists playerStatistics as
		select 
		p.id,
		p.name,
		m.created_at,
		case when 
			(t1.player1_id = p.id or t1.player2_id = p.id) and (m.score_team1 > m.score_team2) or 
			(t2.player1_id = p.id or t2.player2_id = p.id) and (m.score_team2 > m.score_team1)
		 then 1 else 0 end as won,
		case when t1.player1_id = p.id or t1.player2_id = p.id then m.score_team1 else m.score_team2 end as pointsWon,
		case when t1.player1_id = p.id or t1.player2_id = p.id then m.score_team2 else m.score_team1 end as pointsLost
		from matches m	
		join teams t1 on m.team1_id = t1.id
		join teams t2 on m.team2_id = t2.id
		join players p on t1.player1_id = p.id or t1.player2_id = p.id or t2.player1_id = p.id or t2.player2_id = p.id
		where m.deleted_at is null
	`
)

var db *gorm.DB

func initDb() error {
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")

	db.Exec(playerStatisticsView)
	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.Match{})
	db.AutoMigrate(&models.Team{})
	db.AutoMigrate(&models.User{})

	var count int
	db.First(&models.User{}).Count(&count)

	if count == 0 {
		player1 := models.Player{Name: "Raphi"}
		player2 := models.Player{Name: "Robert"}
		player3 := models.Player{Name: "Lukas"}
		player4 := models.Player{Name: "Richie"}
		player5 := models.Player{Name: "Dominik"}
		player6 := models.Player{Name: "Roman"}

		user1 := models.User{Email: "raphi011@gmail.com", Player: player1}
		user2 := models.User{Email: "", Player: player2}
		user3 := models.User{Email: "", Player: player3}
		user4 := models.User{Email: "Rb1@outlook.at", Player: player4}
		user5 := models.User{Email: "Rieder.dominik@gmail.com", Player: player5}
		user6 := models.User{Email: "", Player: player6}

		db.Create(&user1)
		db.Create(&user2)
		db.Create(&user3)
		db.Create(&user4)
		db.Create(&user5)
		db.Create(&user6)
	}

	return err
}

func getMatches() []models.Match {
	var matches []models.Match
	db.
		Preload("Team1.Player1").
		Preload("Team1.Player2").
		Preload("Team2.Player1").
		Preload("Team2.Player2").
		Order("created_at desc").
		Find(&matches)

	return matches
}

func getPlayer(id uint) models.Player {
	var player models.Player

	db.First(&player, id)

	return player
}

type statistic struct {
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

func playersStatistic(filter string) []statistic {
	var statistics []statistic

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

	statisticsQuery().
		Where("playerStatistics.created_at > ?", timeFilter).
		Order("percentage desc").
		Scan(&statistics)

	return statistics
}

func playerStatistic(ID uint) statistic {
	var s statistic

	statisticsQuery().Where("playerStatistics.id = ?", ID).First(&s)

	return s
}

func statisticsQuery() *gorm.DB {
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

func getMatch(id uint) models.Match {
	var match models.Match
	db.
		Preload("Team1.Player1").
		Preload("Team1.Player2").
		Preload("Team2.Player1").
		Preload("Team2.Player2").
		Preload("CreatedBy").
		First(&match, id)

	return match
}

func getTeam(player1ID, player2ID uint) models.Team {
	if player1ID > player2ID {
		player1ID, player2ID = player2ID, player1ID
	}

	var team models.Team

	db.Where(models.Team{Player1ID: player1ID, Player2ID: player2ID}).FirstOrCreate(&team)

	return team
}

func deleteMatch(match models.Match) {
	db.Delete(&match)
}

func getUserByEmail(email string) models.User {
	var user models.User

	db.Where(&User{Email: email}).First(&user)

	return user
}

func createMatch(
	player1ID uint,
	player2ID uint,
	player3ID uint,
	player4ID uint,
	scoreTeam1 int,
	scoreTeam2 int,
	userEmail string,
) (models.Match, error) {
	user := getUserByEmail(userEmail)
	team1 := getTeam(player1ID, player2ID)
	team2 := getTeam(player3ID, player4ID)

	match := models.Match{
		Team1:       team1,
		Team2:       team2,
		ScoreTeam1:  scoreTeam1,
		ScoreTeam2:  scoreTeam2,
		CreatedByID: user.ID,
	}

	db.Create(&match)

	return match, nil
}

func getPlayers() []models.Player {
	var players []models.Player
	db.Order("name").Find(&players)
	return players
}

func updateUser(user models.User) {
	db.Save(&user)
}

func createPlayer(name string) (models.Player, error) {
	player := models.Player{
		Name: name,
	}
	db.Create(&player)

	return player, nil
}
