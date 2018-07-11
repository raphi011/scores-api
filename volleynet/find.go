package volleynet

func FindPlayer(players []Player, playerID int) *Player {
	for i, _ := range players {
		p := &players[i]
		if p.ID == playerID {
			return p
		}
	}

	return nil
}

func FindTournament(tournaments []FullTournament, tournamentID int) *FullTournament {
	for i, _ := range tournaments {
		t := &tournaments[i]

		if t.ID == tournamentID {
			return t
		}
	}

	return nil
}

func FindTeam(teams []TournamentTeam, tournamentID, player1ID, player2ID int) *TournamentTeam {
	for i, _ := range teams {
		t := &teams[i]
		if t.TournamentID == tournamentID && t.Player1.ID == player1ID && t.Player2.ID == player2ID {
			return t
		}
	}

	return nil
}
