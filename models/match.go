package models

import (
	"github.com/jinzhu/gorm"
)

type Match struct {
	gorm.Model
	Team1      Team
	Team1ID    uint
	Team2      Team
	Team2ID    uint
	ScoreTeam1 int
	ScoreTeam2 int
}

type MatchDto struct {
	ScoreTeam1 int `json:"scoreTeam1"`
	ScoreTeam2 int `json:"scoreTeam2"`
}

type Matches []Match
