package volleynet

func MergeTournamentTeam(syncType string, persisted, current *TournamentTeam) *TournamentTeam {
	persisted.Deregistered = current.Deregistered

	if syncType == SyncTeamDone {
		persisted.PrizeMoney = current.PrizeMoney
		persisted.Rank = current.Rank
		persisted.WonPoints = current.WonPoints
	} else if syncType == SyncTeamUpcoming {
		persisted.Seed = current.Seed
		persisted.TotalPoints = current.TotalPoints
	}

	return persisted
}

func MergeTournament(syncType string, persisted, current *FullTournament) *FullTournament {
	if syncType == SyncTournamentUpcomingToCanceled {
		persisted.Name = current.Name
		persisted.HTMLNotes = current.HTMLNotes
	} else {
		persisted.Start = current.Start
		persisted.End = current.End
		persisted.Name = current.Name
		persisted.Link = current.Link
		persisted.Location = current.Location
		persisted.HTMLNotes = current.HTMLNotes
		persisted.Mode = current.Mode
		persisted.MinTeams = current.MinTeams
		persisted.MaxTeams = current.MaxTeams
		persisted.MaxPoints = current.MaxPoints
		persisted.EndRegistration = current.EndRegistration
		persisted.Organiser = current.Organiser
		persisted.Phone = current.Phone
		persisted.Email = current.Email
		persisted.Web = current.Web
		persisted.CurrentPoints = current.CurrentPoints
		persisted.LivescoringLink = current.LivescoringLink
		persisted.Latitude = current.Latitude
		persisted.Longitude = current.Longitude

		if syncType == SyncTournamentUpcomingToDone {
			persisted.EntryLink = ""
			persisted.RegistrationOpen = false

		} else if syncType == SyncTournamentUpcoming {
			persisted.EntryLink = current.EntryLink
			persisted.RegistrationOpen = current.RegistrationOpen

		}
	}

	persisted.Status = current.Status

	return persisted
}

func MergePlayer(persisted, current *Player) *Player {
	persisted.FirstName = current.FirstName
	persisted.LastName = current.LastName
	persisted.Login = current.Login
	persisted.Birthday = current.Birthday
	persisted.Gender = current.Gender
	persisted.TotalPoints = current.TotalPoints
	persisted.Rank = current.Rank
	persisted.Club = current.Club
	persisted.CountryUnion = current.CountryUnion
	persisted.License = current.License

	return persisted
}

type PlayerSyncInformation struct {
	IsNew     bool
	OldPlayer *Player
	NewPlayer *Player
}

func SyncPlayers(persisted []Player, current ...Player) []PlayerSyncInformation {
	ps := []PlayerSyncInformation{}
	for i, _ := range current {
		newPlayer := &current[i]
		oldPlayer := GetPlayer(persisted, newPlayer.ID)

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
		oldTournament := GetTournament(persisted, newTournament.ID)

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
		oldTeam := GetTeam(persisted, newTeam.TournamentID, newTeam.Player1.ID, newTeam.Player2.ID)
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
