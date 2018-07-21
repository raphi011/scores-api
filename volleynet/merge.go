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
