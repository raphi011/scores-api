package sql

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

func setupDB(t testing.TB) *sqlx.DB {
	t.Helper()

	dbProvider := "sqlite3"
	connectionString := "file::memory:?_busy_timeout=5000&mode=memory"

	p := os.Getenv("TEST_DB_PROVIDER")
	c := os.Getenv("TEST_DB_CONNECTION")

	if p != "" && c != "" {
		dbProvider = p
		connectionString = c
	}

	db, err := sqlx.Open(dbProvider, connectionString)
	assert(t, "unable to open db: %v", err)

	err = db.Ping()
	assert(t, "unable to connect to db: %v", err)

	err = migrateAll(dbProvider, db)
	assert(t, "migration failed: %v", err)

	_, err = exec(db, "test/delete-all", make(map[string]interface{}));
	assert(t, "db cleanup failed: %v", err)

	return db
}

func assert(t testing.TB, message string, err error) {
	if err != nil {
		t.Fatalf(message, err)
	}
}

type P struct {
	Gender string
	TotalPoints int
	Rank int
	ID int
}

func createPlayers(t testing.TB, db *sqlx.DB, players ...P) []*volleynet.Player {
	newPlayers := []*volleynet.Player{}
	playerRepo := &playerRepository{DB: db}

	for _, p := range players {
		persistedPlayer, err := playerRepo.New(&volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{
				TrackedModel: scores.TrackedModel{ Model: scores.Model{ID: p.ID }},
			},
			Gender: p.Gender,
			TotalPoints: p.TotalPoints,
			Rank: p.Rank,
		})
		assert(t, "playerRepo.New() failed", err)

		newPlayers = append(newPlayers, persistedPlayer)
	}

	return newPlayers
}

type T struct {
	ID int
}

func createTournaments(t testing.TB, db *sqlx.DB, tournaments ...T) []*volleynet.FullTournament {
	newTournaments := []*volleynet.FullTournament{}
	tournamentRepo := &tournamentRepository{DB: db}

	for _, tournament := range tournaments {
		persistedTournament, err := tournamentRepo.New(&volleynet.FullTournament{
			Tournament: volleynet.Tournament{
				ID: tournament.ID,
			},
		})
		assert(t, "tournamentRepo.New() failed", err)

		newTournaments = append(newTournaments, persistedTournament)
	}

	return newTournaments
}

type TT struct {
	TournamentID int     
	TotalPoints  int     
	Seed         int     
	Rank         int     
	WonPoints    int     
	Player1      *volleynet.Player 
	Player2      *volleynet.Player 
	PrizeMoney   float32 
	Deregistered bool    
}

func createTeams(t testing.TB, db *sqlx.DB, teams ...TT) []*volleynet.TournamentTeam {
	newTeams := []*volleynet.TournamentTeam{}
	teamRepo := &teamRepository{DB: db}

	for _, team := range teams {
		persistedTeam, err := teamRepo.New(&volleynet.TournamentTeam{
			TournamentID: team.TournamentID,
			TotalPoints: team.TotalPoints,
			Seed: team.Seed,
			Rank: team.Rank,
			WonPoints: team.WonPoints,
			Player1: team.Player1,
			Player2: team.Player2,
			PrizeMoney: team.PrizeMoney,
			Deregistered: team.Deregistered,
		})

		assert(t, "teamRepo.New() failed", err)

		newTeams = append(newTeams, persistedTeam)
	}

	return newTeams
}
