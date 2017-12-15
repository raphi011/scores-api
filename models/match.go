package models

import (
	"github.com/jinzhu/gorm"
)

type Match struct {
	gorm.Model
	Team1       Team
	Team1ID     uint
	Team2       Team
	Team2ID     uint
	ScoreTeam1  int
	ScoreTeam2  int
	CreatedByID uint
	CreatedBy   User
}

type Matches []Match

func (m *Match) DeleteMatch(db *gorm.DB) {
	db.Delete(&m)
}

func (m *Match) CreateMatch(
	db *gorm.DB,
	player1ID uint,
	player2ID uint,
	player3ID uint,
	player4ID uint,
	scoreTeam1 int,
	scoreTeam2 int,
	userEmail string) {

	user := &User{}
	team1 := &Team{}
	team2 := &Team{}

	user.GetUserByEmail(db, userEmail)
	team1.GetTeam(db, player1ID, player2ID)
	team2.GetTeam(db, player3ID, player4ID)

	m.Team1 = *team1
	m.Team2 = *team2
	m.ScoreTeam1 = scoreTeam1
	m.ScoreTeam2 = scoreTeam2
	m.CreatedByID = user.ID

	db.Create(&m)
}

func (m *Match) GetMatch(db *gorm.DB, ID uint) {
	db.
		Preload("Team1.Player1").
		Preload("Team1.Player2").
		Preload("Team2.Player1").
		Preload("Team2.Player2").
		Preload("CreatedBy").
		First(&m, ID)
}

func GetMatches(db *gorm.DB) Matches {
	var matches []Match

	db.
		Preload("Team1.Player1").
		Preload("Team1.Player2").
		Preload("Team2.Player1").
		Preload("Team2.Player2").
		Order("created_at desc").
		Find(&matches)

	return matches
}
