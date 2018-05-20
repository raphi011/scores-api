package sqlite

import (
	"database/sql"

	"github.com/raphi011/scores/volleynet"
)

type VolleynetService struct {
	DB *sql.DB
}

func scanTournaments(db *sql.DB, query string, args ...interface{}) ([]volleynet.FullTournament, error) {
	tournaments := []volleynet.FullTournament{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		tournament, err := scanTournament(rows)

		if err != nil {
			return nil, err
		}

		tournaments = append(tournaments, *tournament)
	}

	return tournaments, nil
}

func scanTournament(scanner scan) (*volleynet.FullTournament, error) {
	t := &volleynet.FullTournament{}

	err := scanner.Scan(
		&t.ID,
		&t.StartDate,
		&t.EndDate,
		&t.Name,
		&t.League,
		&t.Link,
		&t.EntryLink,
		&t.Status,
		&t.RegistrationOpen,
		&t.Location,
		&t.HTMLNotes,
		&t.Mode,
		&t.MaxPoints,
		&t.Mode,
		&t.MinTeams,
		&t.EndRegistration,
		&t.Organiser,
		&t.Phone,
		&t.Email,
		&t.Web,
		&t.CurrentPoints,
		&t.LivescoringLink,
	)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *VolleynetService) Tournament(tournamentID string) (*volleynet.FullTournament, error) {
	return nil, nil
}

func (s *VolleynetService) AllTournament() (*volleynet.FullTournament, error) {
	return nil, nil
}

func (s *VolleynetService) SearchPlayers() ([]volleynet.Player, error) {
	return nil, nil
}
