package sqlite

import (
	"database/sql"
	"errors"

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
		&t.EndRegistration,
		&t.Organiser,
		&t.Phone,
		&t.Email,
		&t.Web,
		&t.CurrentPoints,
		&t.LivescoringLink,
		&t.Latitude,
		&t.Longitude,
	)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *VolleynetService) Tournament(tournamentID string) (*volleynet.FullTournament, error) {
	row := s.DB.QueryRow(tournamentSelectSQL, tournamentID)

	return scanTournament(row)
}

func (s *VolleynetService) AllTournaments() ([]volleynet.FullTournament, error) {
	return scanTournaments(s.DB, tournamentsBaseSelectSQL)
}

func (s *VolleynetService) SearchPlayers() ([]volleynet.Player, error) {
	return nil, nil
}

const (
	tournamentsBaseSelectSQL = `
	SELECT
		t.id,
		t.created_at,
		t.updated_at,
		t.gender,
		t.start,
		t.end,
		t.name,
		t.league,
		t.link,
		t.entry_link,
		t.status,
		t.registration_open,
		t.location,
		t.html_notes,
		t.mode,
		t.max_points,
		t.min_teams,
		t.end_registration,
		t.organiser,
		t.phone,
		t.email,
		t.web,
		t.current_points,
		t.live_scoring_link,
		t.loc_lat,
		t.loc_lon
	FROM volleynetTournaments t
	`

	tournamentsFilterSelectSQL = tournamentsBaseSelectSQL + `
	 WHERE t.gender = $1 AND t.league = $2 AND t.season = $3
	`

	tournamentSelectSQL = tournamentsBaseSelectSQL + " WHERE t.id = $1"

	tournamentsInsertSQL = `
		INSERT INTO volleynetTournaments
		(
			id,
			created_at,
			updated_at,
			gender,
			start,
			end,
			name,
			league,
			link,
			entry_link,
			status,
			registration_open,
			location,
			html_notes,
			mode,
			max_points,
			min_teams,
			end_registration,
			organiser,
			phone,
			email,
			web,
			current_points,
			live_scoring_link,
			loc_lat,
			loc_lon,
			season
		)
		VALUES
		(
			$1,
			CURRENT_TIMESTAMP,
			CURRENT_TIMESTAMP,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14,
			$15,
			$16,
			$17,
			$18,
			$19,
			$20,
			$21,
			$22,
			$23,
			$24,
			$25
		)
	`

	tournamentsUpdateSQL = `
		UPDATE volleynetTournaments
			SET updated_at = CURRENT_TIMESTAMP,
			SET gender = $1,
			SET start = $2,
			SET end = $3,
			SET name = $4,
			SET league = $5,
			SET link = $6,
			SET entry_link = $7,
			SET status = $8,
			SET registration_open = $9,
			SET location = $10,
			SET html_notes = $11,
			SET mode = $12,
			SET max_points = $13,
			SET min_teams = $14,
			SET end_registration = $15,
			SET organiser = $16,
			SET phone = $17,
			SET email = $18,
			SET web = $19,
			SET current_points = $20,
			SET live_scoring_link = $21,
			SET loc_lat = $22,
			SET loc_lon = $23,
			SET season = $24
		WHERE id = $25
	
	`
)

func (s *VolleynetService) GetTournaments(gender, league string, season int) ([]volleynet.FullTournament, error) {
	return scanTournaments(s.DB, tournamentsFilterSelectSQL, gender, league, season)
}

type TournamentSyncInformation struct {
	New        bool
	ID         string
	Tournament volleynet.Tournament
}

func (s *VolleynetService) containsTournament(tournaments []volleynet.FullTournament, tournamentID string) bool {
	for _, t := range tournaments {
		if t.ID == tournamentID {
			return true
		}
	}

	return false
}

func (s *VolleynetService) SyncInformation(tournaments []volleynet.Tournament) ([]TournamentSyncInformation, error) {
	persistedTournaments, err := s.AllTournaments()

	if err != nil {
		return nil, err
	}

	ts := []TournamentSyncInformation{}
	for _, t := range tournaments {
		new := !s.containsTournament(persistedTournaments, t.ID)

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

func (s *VolleynetService) NewTournament(t *volleynet.FullTournament) error {
	_, err := s.DB.Exec(tournamentsInsertSQL,
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
	)

	return err
}

func (s *VolleynetService) UpdateTournament(t *volleynet.FullTournament) error {
	result, err := s.DB.Exec(
		tournamentsUpdateSQL,
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
		t.ID)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected != 1 {
		return errors.New("Tournament not found")
	}

	return nil
}
