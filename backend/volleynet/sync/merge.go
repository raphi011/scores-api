package sync

import "github.com/raphi011/scores/volleynet"

// MergeTournamentTeam merges two tournament teams depending on the syncType
// and returns the new TournamentTeam.
func MergeTournamentTeam(persisted, current *volleynet.TournamentTeam) *volleynet.TournamentTeam {
	merged := *persisted
	merged.Deregistered = current.Deregistered

	isTournamentDone := persisted.Rank == 0 && current.Rank > 0

	if isTournamentDone {
		merged.PrizeMoney = current.PrizeMoney
		merged.Rank = current.Rank
		merged.WonPoints = current.WonPoints
	} else {
		merged.Seed = current.Seed
		merged.TotalPoints = current.TotalPoints
	}

	return &merged
}

// MergeTournament merges two tournaments depending on the syncType
// and returns the new Tournament.
func MergeTournament(persisted, current *volleynet.FullTournament) *volleynet.FullTournament {
	merged := *persisted

	if persisted.Status == volleynet.StatusUpcoming && current.Status == volleynet.StatusCanceled {
		merged.Name = current.Name
		merged.HTMLNotes = current.HTMLNotes
	} else {
		merged.Start = current.Start
		merged.End = current.End
		merged.Name = current.Name
		merged.Link = current.Link
		merged.Location = current.Location
		merged.HTMLNotes = current.HTMLNotes
		merged.Mode = current.Mode
		merged.MinTeams = current.MinTeams
		merged.MaxTeams = current.MaxTeams
		merged.MaxPoints = current.MaxPoints
		merged.EndRegistration = current.EndRegistration
		merged.Organiser = current.Organiser
		merged.Phone = current.Phone
		merged.Email = current.Email
		merged.Website = current.Website
		merged.CurrentPoints = current.CurrentPoints
		merged.LivescoringLink = current.LivescoringLink
		merged.Latitude = current.Latitude
		merged.Longitude = current.Longitude
		merged.SignedupTeams = current.SignedupTeams

		if persisted.Status == volleynet.StatusUpcoming && current.Status == volleynet.StatusDone {
			merged.EntryLink = ""
			merged.RegistrationOpen = false

		} else if current.Status == volleynet.StatusUpcoming {
			merged.EntryLink = current.EntryLink
			merged.RegistrationOpen = current.RegistrationOpen
		}
	}

	merged.Status = current.Status

	return &merged
}

// MergePlayer merges two players depending on the syncType
// and returns the new player.
func MergePlayer(persisted, current *volleynet.Player) *volleynet.Player {
	merged := *persisted
	merged.FirstName = current.FirstName
	merged.LastName = current.LastName
	merged.Birthday = current.Birthday
	merged.Gender = current.Gender
	merged.TotalPoints = current.TotalPoints
	merged.Rank = current.Rank
	merged.Club = current.Club
	merged.CountryUnion = current.CountryUnion
	merged.License = current.License

	return &merged
}
