package volleynet

import "github.com/raphi011/scores/volleynet"

func mergeTournament(old, new FullTournament) {

}

func mergeTournamentTeam(old, new TournamentTeam) {

}

func mergePlayer(old, new Player) {

}

func containsPlayer(players []Player, playerID int) bool {
	for _, p := range players {
		if p.ID == playerID {
			return true
		}
	}

	return false
}

type PlayerSyncInformation struct {
	New    bool
	ID     int
	Player *volleynet.Player
}

func SyncPlayers(gender string, persisted []Player, current []Player) error {
	for _, p := range current {
		new := !containsPlayer(persisted, p.ID)

		if new {
			s.NewPlayer(&p)
		} else {
			s.UpdatePlayer(&p)
		}
	}

	return nil
}

func containsTournament(tournaments []FullTournament, tournamentID int) bool {
	for _, t := range tournaments {
		if t.ID == tournamentID {
			return true
		}
	}

	return false
}

type TournamentSyncInformation struct {
	New        bool
	ID         int
	Tournament volleynet.Tournament
}

func syncTournamentInformation(persisted []FullTournament, current ...Tournament) (
	[]TournamentSyncInformation, error) {

	if err != nil {
		return nil, err
	}

	ts := []TournamentSyncInformation{}
	for _, t := range tournaments {
		new := !containsTournament(persisted, t.ID)

		if new {
			// for now only add new tournaments
			ts = append(ts, TournamentSyncInformation{
				ID:         t.ID,
				Tournament: t,
				New:        new,
			})
		}
	}

	return ts, nil
}
