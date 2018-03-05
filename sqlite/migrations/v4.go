package migrations

var V4 = []string{
	userPasswordAuthentication,
}

var ResetV4 = []string{
	"matches",
	"teams",
	"players",
	"users",
	"groups",
	"groupPlayers",
}

const (
	userPasswordAuthentication = `
		ALTER TABLE "users" ADD COLUMN "salt" blob;
		ALTER TABLE "users" ADD COLUMN "hash" blob;
		ALTER TABLE "users" ADD COLUMN "iterations" integer;
	`
)
