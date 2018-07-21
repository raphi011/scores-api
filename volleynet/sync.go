package volleynet

type PlayerSyncInformation struct {
	IsNew     bool
	OldPlayer *Player
	NewPlayer *Player
}

func SyncPlayers(persisted []Player, current ...Player) []PlayerSyncInformation {
	ps := []PlayerSyncInformation{}
	for i, _ := range current {
		newPlayer := &current[i]
		oldPlayer := FindPlayer(persisted, newPlayer.ID)

		ps = append(ps, PlayerSyncInformation{
			NewPlayer: newPlayer,
			OldPlayer: oldPlayer,
			IsNew:     oldPlayer == nil,
		})
	}

	return ps
}

type TournamentSyncInformation struct {
	IsNew         bool
	SyncType      string
	OldTournament *FullTournament
	NewTournament *Tournament
}

const (
	SyncTournamentNoUpdate           = "SyncTournamentNoUpdate"
	SyncTournamentUpcomingToCanceled = "SyncTournamentUpcomingToCanceled"
	SyncTournamentUpcoming           = "SyncTournamentUpcoming"
	SyncTournamentUpcomingToDone     = "SyncTournamentUpcomingToDone"
	SyncTournamentNew                = "SyncTournamentNew"
)

func tournamentSyncType(persisted *FullTournament, current *Tournament) string {
	if persisted == nil {
		return SyncTournamentNew
	}
	if persisted.Status != StatusUpcoming {
		return SyncTournamentNoUpdate
	}
	if current.Status == StatusCanceled {
		return SyncTournamentUpcomingToCanceled
	}
	if current.Status == StatusUpcoming {
		return SyncTournamentUpcoming
	}
	if current.Status == StatusDone {
		return SyncTournamentUpcomingToDone
	}

	return ""
}

func SyncTournaments(persisted []FullTournament, current ...Tournament) []TournamentSyncInformation {
	ts := []TournamentSyncInformation{}
	for i, _ := range current {
		newTournament := &current[i]
		oldTournament := FindTournament(persisted, newTournament.ID)

		syncType := tournamentSyncType(oldTournament, newTournament)

		if syncType == SyncTournamentNoUpdate {
			continue
		}

		ts = append(ts, TournamentSyncInformation{
			OldTournament: oldTournament,
			NewTournament: newTournament,
			IsNew:         oldTournament == nil,
			SyncType:      syncType,
		})
	}

	return ts
}

const (
	SyncTeamNew      = "SyncTeamNew"
	SyncTeamUpcoming = "SyncTeamUpcoming"
	SyncTeamDone     = "SyncTeamDone"
	SyncTeamNoUpdate = "SyncTeamNoUpdate"
)

type TournamentTeamSyncInformation struct {
	IsNew    bool
	SyncType string
	OldTeam  *TournamentTeam
	NewTeam  *TournamentTeam
}

func tournamentTeamSyncType(tournamentSyncType string, persisted, current *TournamentTeam) string {
	if tournamentSyncType == SyncTournamentNew {
		return SyncTeamNew
	}
	if tournamentSyncType == SyncTournamentUpcoming {
		return SyncTeamUpcoming
	}
	if tournamentSyncType == SyncTournamentUpcomingToDone {
		return SyncTeamDone
	}
	if tournamentSyncType == SyncTournamentNoUpdate {
		return SyncTeamNoUpdate
	}

	return ""
}

func SyncTournamentTeams(tournamentSyncType string, persisted, current []TournamentTeam) []TournamentTeamSyncInformation {
	ts := []TournamentTeamSyncInformation{}
	for i, _ := range current {
		newTeam := &current[i]
		oldTeam := FindTeam(persisted, newTeam.TournamentID, newTeam.Player1.ID, newTeam.Player2.ID)
		syncType := tournamentTeamSyncType(tournamentSyncType, oldTeam, newTeam)

		if syncType == SyncTeamNoUpdate {
			continue
		}

		ts = append(ts, TournamentTeamSyncInformation{
			OldTeam:  oldTeam,
			NewTeam:  newTeam,
			IsNew:    oldTeam == nil,
			SyncType: syncType,
		})
	}

	return ts
}
