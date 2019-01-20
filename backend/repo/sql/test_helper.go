package sql

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sql/crud"
	"github.com/raphi011/scores/repo/sql/migrate"
	"github.com/raphi011/scores/volleynet"
)

// RepositoriesTest creates a database that respects the `TEST_DB_PROVIDER`
// and `TEST_DB_CONNECTION` env variables and returns it and all repositories
// with this db.
func RepositoriesTest(t testing.TB) (*repo.Repositories, *sqlx.DB) {
	db := SetupDB(t)

	return &repo.Repositories{
		UserRepo: &userRepository{DB: db},
		PlayerRepo: &playerRepository{DB: db},
		TournamentRepo: &tournamentRepository{DB: db},
		TeamRepo: &teamRepository{DB: db},
	}, db
}

// SetupDB sets up a database connection, runs all migrations and
// clears any existing data.
func SetupDB(t testing.TB) *sqlx.DB {
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
	test.Check(t, "unable to open db: %v", err)

	err = db.Ping()
	test.Check(t, "unable to connect to db: %v", err)

	err = migrate.All(dbProvider, db)
	test.Check(t, "migration failed: %v", err)

	err = crud.Execute(db, "test/delete-all");
	test.Check(t, "db cleanup failed: %v", err)

	return db
}

// P is a helper struct to create players.
type P struct {
	Gender string
	TotalPoints int
	Rank int
	ID int
}

// CreatePlayers is a handy helper to create multiple players.
func CreatePlayers(t testing.TB, db *sqlx.DB, players ...P) []*volleynet.Player {
	newPlayers := []*volleynet.Player{}
	playerRepo := &playerRepository{DB: db}

	for _, p := range players {
		persistedPlayer, err := playerRepo.New(&volleynet.Player{
			ID: p.ID,
			Gender: p.Gender,
			TotalPoints: p.TotalPoints,
			Rank: p.Rank,
		})

		test.Check(t, "playerRepo.New() failed", err)

		newPlayers = append(newPlayers, persistedPlayer)
	}

	return newPlayers
}

// T is a helper struct to create tournaments.
type T struct {
	ID int
	Status string
}

// CreateTournaments is a handy helper to create multiple tournaments.
func CreateTournaments(t testing.TB, db *sqlx.DB, tournaments ...T) []*volleynet.Tournament {
	newTournaments := []*volleynet.Tournament{}
	tournamentRepo := &tournamentRepository{DB: db}

	for _, tournament := range tournaments {
		persistedTournament, err := tournamentRepo.New(&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo {
				ID: tournament.ID,
				Status: tournament.Status,
			},
		})
		test.Check(t, "tournamentRepo.New() failed", err)

		newTournaments = append(newTournaments, persistedTournament)
	}

	return newTournaments
}

// TT is a helper struct to create teams.
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

// CreateTeams is a handy helper to create multiple teams.
func CreateTeams(t testing.TB, db *sqlx.DB, teams ...TT) []*volleynet.TournamentTeam {
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

		test.Check(t, "teamRepo.New() failed", err)

		newTeams = append(newTeams, persistedTeam)
	}

	return newTeams
}
