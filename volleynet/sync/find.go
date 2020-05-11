package sync

import "github.com/raphi011/scores/volleynet"

// FindPlayer searches a player slice for a player, returns nil if not found
func FindPlayer(players []*volleynet.Player, playerID int) *volleynet.Player {
	for i := range players {
		p := players[i]
		if p.ID == playerID {
			return p
		}
	}

	return nil
}

// FindTournament searches a tournament slice for a tournament, returns nil if not found
func FindTournament(tournaments []*volleynet.Tournament, tournamentID int) *volleynet.Tournament {
	for i := range tournaments {
		t := tournaments[i]

		if t.ID == tournamentID {
			return t
		}
	}

	return nil
}

// FindTeam searches a team slice for a team, returns nil if not found
func FindTeam(teams []*volleynet.TournamentTeam, tournamentID, player1ID, player2ID int) *volleynet.TournamentTeam {
	for i := range teams {
		t := teams[i]
		if t.TournamentID == tournamentID && t.Player1.ID == player1ID && t.Player2.ID == player2ID {
			return t
		}
	}

	return nil
}
