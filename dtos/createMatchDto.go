package dtos

type CreateMatchDto struct {
	Player1ID  int `json:"player1Id"`
	Player2ID  int `json:"player2Id"`
	Player3ID  int `json:"player3Id"`
	Player4ID  int `json:"player4Id"`
	ScoreTeam1 int `json:"scoreTeam1"`
	ScoreTeam2 int `json:"scoreTeam2"`
}
