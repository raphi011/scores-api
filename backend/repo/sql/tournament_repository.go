package sql

import (
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

// TournamentRepository implements VolleynetRepository interface
type TournamentRepository struct {
	DB *sqlx.DB
}

// Get loads a tournament by its id
func (s *TournamentRepository) Get(tournamentID int) (*volleynet.FullTournament, error) {
	row := s.DB.QueryRow(query("tournament/select-by-id"), tournamentID)
	return scanTournament(row)
}

// All loads all tournaments
// Note: should only be used for debugging
func (s *TournamentRepository) All() ([]*volleynet.FullTournament, error) {
	return scanTournaments(s.DB, query("tournament/select-all"))
}

// Season loads all tournaments of a season
func (s *TournamentRepository) Season(season int) ([]*volleynet.FullTournament, error) {
	return scanTournaments(s.DB, query("tournament/select-by-season"), season)
}

// New creates a new tournament
func (s *TournamentRepository) New(t *volleynet.FullTournament) error {
	_, err := s.DB.Exec(query("tournament/insert"),
		t.ID,
		t.Gender,
		t.Start,
		t.End,
		t.Name,
		t.League,
		t.Link,
		t.EntryLink,
		t.Status,
		t.RegistrationOpen,
		t.Location,
		t.HTMLNotes,
		t.Mode,
		t.MaxPoints,
		t.MinTeams,
		t.MaxTeams,
		t.EndRegistration,
		t.Organiser,
		t.Phone,
		t.Email,
		t.Web,
		t.CurrentPoints,
		t.LivescoringLink,
		t.Longitude,
		t.Latitude,
		t.Season,
		t.SignedupTeams,
	)

	return err
}

// UpdatedSince gets all tournaments that were updated after a certain time
func (s *TournamentRepository) UpdatedSince(updatedSince time.Time) ([]*volleynet.FullTournament, error) {
	return scanTournaments(s.DB,
		query("tournament/select-by-updated-since"),
		updatedSince,
	)
}

// Update updates a tournament
func (s *TournamentRepository) Update(t *volleynet.FullTournament) error {
	result, err := s.DB.Exec(
		query("tournament/update"),
		t.Gender,
		t.Start,
		t.End,
		t.Name,
		t.League,
		t.Link,
		t.EntryLink,
		t.Status,
		t.RegistrationOpen,
		t.Location,
		t.HTMLNotes,
		t.Mode,
		t.MaxPoints,
		t.MinTeams,
		t.MaxTeams,
		t.EndRegistration,
		t.Organiser,
		t.Phone,
		t.Email,
		t.Web,
		t.CurrentPoints,
		t.LivescoringLink,
		t.Longitude,
		t.Latitude,
		t.Season,
		t.SignedupTeams,
		t.ID)

	if err != nil {
		return mapError(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return scores.ErrorNotFound
	}

	return nil
}

// Filter loads all tournaments by season, league and gender
func (s *TournamentRepository) Filter(
	season int,
	leagues []string,
	genders []string) ([]*volleynet.FullTournament, error) {

	rawQuery := query("tournament/select-by-filter")
	query, args, err := sqlx.In(rawQuery, gender, league, season)

	return scanTournaments(s.DB, query, args...)
}

func scanTournaments(db *sqlx.DB, query string, args ...interface{}) ([]*volleynet.FullTournament, error) {
	tournaments := []*volleynet.FullTournament{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return tournaments, mapError(err)
	}

	defer rows.Close()

	for rows.Next() {
		tournament, err := scanTournament(rows)

		if err != nil {
			return , err
		}

		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}

func scanTournament(scanner scan) (*volleynet.FullTournament, error) {
	t := &volleynet.FullTournament{}
	t.Teams = []volleynet.TournamentTeam{}

	err := scanner.Scan(
		&t.ID,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.Gender,
		&t.Start,
		&t.End,
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
		&t.MinTeams,
		&t.MaxTeams,
		&t.EndRegistration,
		&t.Organiser,
		&t.Phone,
		&t.Email,
		&t.Web,
		&t.CurrentPoints,
		&t.LivescoringLink,
		&t.Latitude,
		&t.Longitude,
		&t.Season,
		&t.SignedupTeams,
	)

	return t, mapError(err)
}
