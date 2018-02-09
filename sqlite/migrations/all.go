package migrations

import "scores-backend/migrate"

var (
	MigrationSet = []migrate.Migration{V1, V2, V3}
	ResetSet     = []migrate.TableNames{ResetV1, ResetV2, ResetV3}
)
