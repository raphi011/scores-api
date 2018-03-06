package migrations

var V3 = []string{
	matchIndexes,
}

var ResetV3 = []string{
	"matches",
	"teams",
	"players",
	"users",
	"groups",
	"groupPlayers",
}

const (
	matchIndexes = `
		CREATE INDEX idx_matches_deleted_at ON matches(deleted_at);
		CREATE INDEX idx_matches_team1_player1_id ON matches(team1_player1_id);
		CREATE INDEX idx_matches_team1_player2_id ON matches(team1_player2_id);
		CREATE INDEX idx_matches_team2_player1_id ON matches(team2_player1_id);
		CREATE INDEX idx_matches_team2_player2_id ON matches(team2_player2_id);
	`
)
