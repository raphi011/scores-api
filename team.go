package scores

import "time"

type Team struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Player1   *Player   `json:"player1"`
	Player1ID uint      `json:"player1Id"`
	Player2   *Player   `json:"player2"`
	Player2ID uint      `json:"player2Id"`
}

type Teams []Team

type TeamService interface {
	ByPlayers(player1ID, player2ID uint) (*Team, error)
	Create(team *Team) (*Team, error)
}
