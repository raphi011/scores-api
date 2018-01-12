package sqlite

/*
func TestGetMatches(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	matchService := &MatchService{DB: db}
	matchService.Create(&scores.Match{Name: "Test1"})
	matchService.Create(&scores.Match{Name: "Test2"})
	players, err := matchService.Players()

	if err != nil {
		t.Errorf("MatchService.Players() err: %s", err)
	} else if len(players) != 2 {
		t.Errorf("MatchService.Players(), want 2 players, got %d ", len(players))
	}
}

func TestCreatePlayer(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	matchService := &MatchService{DB: db}
	player, err := matchService.Create(&scores.Match{Name: "Test"})

	if err != nil {
		t.Error("Can't create player")
	}
	if player.ID == 0 {
		t.Error("PlayerID not assigned")
	}

	playerID := player.ID

	player, err = matchService.Player(playerID)

	if err != nil {
		t.Errorf("matchService.Player() err: %s", err)
	}
	if player.ID != playerID {
		t.Errorf("matchService.Player(), want ID %d, got %d", playerID, player.ID)
	}
}

func TestDeletePlayer(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	matchService := &MatchService{DB: db}
	player, _ := matchService.Create(&scores.Match{Name: "Test"})
	matchService.Create(&scores.Match{Name: "Test2"})

	err := matchService.Delete(player.ID)

	if err != nil {
		t.Errorf("matchService.Delete() err: %s", err)
	}

	players, _ := matchService.Players()
	playerCount := len(players)

	if playerCount != 1 {
		t.Errorf("len(matchService.Players()), want 1, got %d", playerCount)
	}

}
*/
