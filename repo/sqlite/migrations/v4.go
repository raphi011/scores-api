package migrations

var V4 = []string{
	userPasswordAuthentication,
}

var ResetV4 = ResetV3

const (
	userPasswordAuthentication = `
		ALTER TABLE "users" ADD COLUMN "salt" blob;
		ALTER TABLE "users" ADD COLUMN "hash" blob;
		ALTER TABLE "users" ADD COLUMN "iterations" integer;
	`
)
