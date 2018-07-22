package volleynet

// PlayerSyncInformation contains sync information for two `Player`s
type PlayerSyncInformation struct {
	IsNew     bool
	OldPlayer *Player
	NewPlayer *Player
}

// SyncPlayers takes a slice of current and old `Player`s and finds out which
// one is new and which needs to get updated
func SyncPlayers(persisted []Player, current ...Player) []PlayerSyncInformation {
	ps := []PlayerSyncInformation{}
	for i := range current {
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

// TournamentSyncInformation contains sync information for two `Tournament`s
type TournamentSyncInformation struct {
	IsNew         bool
	SyncType      string
	OldTournament *FullTournament
	NewTournament *Tournament
}

// represents the various tournament sync states.
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

// SyncTournaments finds out if and how tournaments have to be synced
func SyncTournaments(persisted []FullTournament, current ...Tournament) []TournamentSyncInformation {
	ts := []TournamentSyncInformation{}
	for i := range current {
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

// represents the various tournament team sync states.
const (
	SyncTeamNew      = "SyncTeamNew"
	SyncTeamUpcoming = "SyncTeamUpcoming"
	SyncTeamDone     = "SyncTeamDone"
	SyncTeamNoUpdate = "SyncTeamNoUpdate"
)

// TournamentTeamSyncInformation contains sync information for two `TournamentTeam`s
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

// SyncTournamentTeams finds out if and how tournament teams have to be synced
func SyncTournamentTeams(tournamentSyncType string, persisted, current []TournamentTeam) []TournamentTeamSyncInformation {
	ts := []TournamentTeamSyncInformation{}
	for i := range current {
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
