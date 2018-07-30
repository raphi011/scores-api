package migrations

var V7 = []string{
	alterTournamentTable,
}

var ResetV7 = ResetV6

const (
	alterTournamentTable = `
		ALTER TABLE "volleynetTournaments" ADD COLUMN "signedup_teams" integer NOT NULL DEFAULT 0;
	`
)
