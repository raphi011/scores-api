package scores

type Match struct {
	Model
	Team1       *Team `json:"team1"`
	Team1ID     uint  `json:"team1Id"`
	Team2       *Team `json:"team2"`
	Team2ID     uint  `json:"team2Id"`
	ScoreTeam1  int   `json:"scoreTeam1"`
	ScoreTeam2  int   `json:"scoreTeam2"`
	CreatedByID uint  `json:"createdById"`
	CreatedBy   *User `json:"createdBy"`
}

type Matches []Match

type MatchService interface {
	Match(matchID uint) (*Match, error)
	PlayerMatches(playerID uint) (*Matches, error)
	Matches() (*Matches, error)
	Create(*Match) error
	Delete(matchID uint) error
}
