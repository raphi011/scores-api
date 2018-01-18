package sqlite

import "testing"

func TestPlayerStatistic(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	t.Skip()
}
