package dtos

type CreateMatchDto struct {
	Player1ID  uint `json:"player1Id"`
	Player2ID  uint `json:"player2Id"`
	Player3ID  uint `json:"player3Id"`
	Player4ID  uint `json:"player4Id"`
	ScoreTeam1 int  `json:"scoreTeam1"`
	ScoreTeam2 int  `json:"scoreTeam2"`
}
