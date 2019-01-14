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
	playerRepo := &PlayerRepository{DB: db}

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
	tournamentRepo := &TournamentRepository{DB: db}

	for _, tournament := range tournaments {
		persistedTournament, err := tournamentRepo.New(&volleynet.FullTournament{
			Tournament: volleynet.Tournament{
				ID: tournament.ID,
				// Season: 2018,
				// League: "pro-tour",
				// Format: "m",
			},
			// Teams: []volleynet.TournamentTeam{},
		})
		assert(t, "tournamentRepo.New() failed", err)

		newTournaments = append(newTournaments, persistedTournament)
	}

	return newTournaments
}