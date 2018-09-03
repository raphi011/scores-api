package sync

import "github.com/raphi011/scores/volleynet"

// MergeTournamentTeam merges two tournament teams depending on the syncType
// and returns the new TournamentTeam.
func MergeTournamentTeam(persisted, current volleynet.TournamentTeam) volleynet.TournamentTeam {
	persisted.Deregistered = current.Deregistered

	isTournamentDone := persisted.Rank == 0 && current.Rank > 0

	if isTournamentDone {
		persisted.PrizeMoney = current.PrizeMoney
		persisted.Rank = current.Rank
		persisted.WonPoints = current.WonPoints
	} else {
		persisted.Seed = current.Seed
		persisted.TotalPoints = current.TotalPoints
	}

	return persisted
}

// MergeTournament merges two tournaments depending on the syncType
// and returns the new Tournament.
func MergeTournament(persisted, current volleynet.FullTournament) volleynet.FullTournament {
	if persisted.Status == volleynet.StatusUpcoming && current.Status == volleynet.StatusCanceled {
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
		persisted.SignedupTeams = current.SignedupTeams

		if persisted.Status == volleynet.StatusUpcoming && current.Status == volleynet.StatusDone {
			persisted.EntryLink = ""
			persisted.RegistrationOpen = false

		} else if current.Status == volleynet.StatusUpcoming {
			persisted.EntryLink = current.EntryLink
			persisted.RegistrationOpen = current.RegistrationOpen
		}
	}

	persisted.Status = current.Status

	return persisted
}

// MergePlayer merges two players depending on the syncType
// and returns the new player.
func MergePlayer(persisted, current *volleynet.Player) *volleynet.Player {
	persisted.FirstName = current.FirstName
	persisted.LastName = current.LastName
	persisted.Birthday = current.Birthday
	persisted.Gender = current.Gender
	persisted.TotalPoints = current.TotalPoints
	persisted.Rank = current.Rank
	persisted.Club = current.Club
	persisted.CountryUnion = current.CountryUnion
	persisted.License = current.License

	return persisted
}
