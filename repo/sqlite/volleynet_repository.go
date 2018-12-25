package sqlite

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

// VolleynetRepository implements VolleynetRepository interface
type VolleynetRepository struct {
	DB *sql.DB
}

// Tournament loads a tournament by its id
func (s *VolleynetRepository) Tournament(tournamentID int) (*volleynet.FullTournament, error) {
	row := s.DB.QueryRow(query("volleynet/select-tournament-by-id"), tournamentID)

	return scanTournament(row)
}

// AllTournaments loads all tournaments
// Note: should only be used for debugging
func (s *VolleynetRepository) AllTournaments() ([]*volleynet.FullTournament, error) {
	return scanTournaments(s.DB, query("volleynet/select-tournament-all"))
}

// SeasonTournaments loads all tournaments of a season
func (s *VolleynetRepository) SeasonTournaments(season int) ([]*volleynet.FullTournament, error) {
	return scanTournaments(s.DB, query("volleynet/select-tournament-by-season"), season)
}

// GetTournaments loads all tournaments by season, league and gender
func (s *VolleynetRepository) GetTournaments(gender, league string, season int) ([]*volleynet.FullTournament, error) {
	return scanTournaments(s.DB, query("volleynet/select-tournament-by-filter"), gender, league, season)
}

// NewTournament creates a new tournament
func (s *VolleynetRepository) NewTournament(t *volleynet.FullTournament) error {
	_, err := s.DB.Exec(query("volleynet/insert-tournament"),
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

// UpdateTournamentTeam updates a tournament team
func (s *VolleynetRepository) UpdateTournamentTeam(t *volleynet.TournamentTeam) error {
	result, err := s.DB.Exec(
		query("volleynet/update-team"),
		t.Rank,
		t.Seed,
		t.TotalPoints,
		t.WonPoints,
		t.PrizeMoney,
		t.Deregistered,
		t.TournamentID,
		t.Player1.ID,
		t.Player2.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Team not found")
	}

	return nil
}

// TournamentsUpdatedSince gets all tournaments that were updated after a certain time
func (s *VolleynetRepository) TournamentsUpdatedSince(updatedSince time.Time) ([]*volleynet.FullTournament, error) {
	return scanTournaments(s.DB,
		query("volleynet/select-tournament-by-updated-since"),
		updatedSince,
	)
}

// UpdateTournament updates a tournament
func (s *VolleynetRepository) UpdateTournament(t *volleynet.FullTournament) error {
	result, err := s.DB.Exec(
		query("volleynet/update-tournament"),
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
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Tournament not found")
	}

	return nil
}

// NewTeam creates a new team
func (s *VolleynetRepository) NewTeam(t *volleynet.TournamentTeam) error {
	_, err := s.DB.Exec(query("volleynet/insert-team"),
		t.TournamentID,
		t.Player1.ID,
		t.Player2.ID,
		t.Rank,
		t.Seed,
		t.TotalPoints,
		t.WonPoints,
		t.PrizeMoney,
		t.Deregistered,
	)

	return err
}

// NewTeams creates multiple new teams
func (s *VolleynetRepository) NewTeams(teams []volleynet.TournamentTeam) error {
	for _, t := range teams {
		err := s.NewTeam(&t)

		if err != nil {
			return err
		}
	}

	return nil
}

func scanTournamentTeams(db *sql.DB, query string, args ...interface{}) ([]volleynet.TournamentTeam, error) {
	teams := []volleynet.TournamentTeam{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		team, err := scanTournamentTeam(rows)

		if err != nil {
			return nil, err
		}

		teams = append(teams, *team)
	}

	return teams, nil
}

func scanTournamentTeam(scanner scan) (*volleynet.TournamentTeam, error) {
	t := &volleynet.TournamentTeam{}
	t.Player1 = &volleynet.Player{}
	t.Player2 = &volleynet.Player{}

	err := scanner.Scan(
		&t.TournamentID,
		&t.Player1.ID,
		&t.Player1.FirstName,
		&t.Player1.LastName,
		&t.Player1.TotalPoints,
		&t.Player1.CountryUnion,
		&t.Player1.Birthday,
		&t.Player1.License,
		&t.Player1.Gender,
		&t.Player2.ID,
		&t.Player2.FirstName,
		&t.Player2.LastName,
		&t.Player2.TotalPoints,
		&t.Player2.CountryUnion,
		&t.Player2.Birthday,
		&t.Player2.License,
		&t.Player2.Gender,
		&t.Rank,
		&t.Seed,
		&t.TotalPoints,
		&t.WonPoints,
		&t.PrizeMoney,
		&t.Deregistered,
	)

	return t, err
}

// UpdateTournamentTeams updates tournament teams
func (s *VolleynetRepository) UpdateTournamentTeams(teams []volleynet.TournamentTeam) error {
	for _, t := range teams {
		if err := s.UpdateTournamentTeam(&t); err != nil {
			return err
		}
	}

	return nil
}

// DeleteTeam deletes a team
func (s *VolleynetRepository) DeleteTeam(t *volleynet.TournamentTeam) error {
	_, err := s.DB.Exec(query("volleynet/delete-team"), t.TournamentID, t.Player1.ID, t.Player2.ID)

	return err
}

// TournamentTeams loads all teams of a tournament
func (s *VolleynetRepository) TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error) {
	return scanTournamentTeams(s.DB,
		query("volleynet/select-team-by-tournament-id"),
		tournamentID)
}

func scanVolleynetPlayers(db *sql.DB, query string, args ...interface{}) ([]volleynet.Player, error) {
	players := []volleynet.Player{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		player, err := scanVolleynetPlayer(rows)

		if err != nil {
			return nil, err
		}

		players = append(players, *player)
	}

	return players, nil
}

func scanVolleynetPlayer(scanner scan) (*volleynet.Player, error) {
	p := &volleynet.Player{}

	err := scanner.Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.Birthday,
		&p.Gender,
		&p.TotalPoints,
		&p.Rank,
		&p.Club,
		&p.CountryUnion,
		&p.License,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

// Ladder gets all players of the passed gender that have a rank
func (s *VolleynetRepository) Ladder(gender string) ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, query("volleynet/select-player-ladder"), gender)
}

// AllPlayers loads all players
// Note: should only be used for debugging
func (s *VolleynetRepository) AllPlayers() ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, query("volleynet/select-player-all"))
}

// Player loads a player
func (s *VolleynetRepository) Player(id int) (*volleynet.Player, error) {
	row := s.DB.QueryRow(
		query("volleynet/select-player-by-id"),
		id,
	)

	return scanVolleynetPlayer(row)
}

// NewPlayer creates a new player
func (s *VolleynetRepository) NewPlayer(p *volleynet.Player) error {
	_, err := s.DB.Exec(query("volleynet/insert-player"),
		p.ID,
		p.FirstName,
		p.LastName,
		p.Birthday,
		p.Gender,
		p.TotalPoints,
		p.Rank,
		p.Club,
		p.CountryUnion,
		p.License,
	)

	return err
}

// UpdatePlayer updates a player
func (s *VolleynetRepository) UpdatePlayer(p *volleynet.Player) error {
	result, err := s.DB.Exec(
		query("volleynet/update-player"),
		p.FirstName,
		p.LastName,
		p.Birthday,
		p.Gender,
		p.TotalPoints,
		p.Rank,
		p.Club,
		p.CountryUnion,
		p.License,
		p.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Player not found")
	}

	return nil
}

// SearchPlayers searches players
// Note: not implemented yet
func (s *VolleynetRepository) SearchPlayers() ([]volleynet.Player, error) {
	return nil, nil
}

func scanTournaments(db *sql.DB, query string, args ...interface{}) ([]*volleynet.FullTournament, error) {
	tournaments := []*volleynet.FullTournament{}
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

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return t, err
}
