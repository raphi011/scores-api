package sqlite

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

type VolleynetService interface {
	Tournament(tournamentID int) (*volleynet.FullTournament, error)
	AllTournaments() ([]volleynet.FullTournament, error)
	SeasonTournaments(season int) ([]volleynet.FullTournament, error)
	GetTournaments(gender, league string, season int) ([]volleynet.FullTournament, error)
	NewTournament(t *volleynet.FullTournament) error
	UpdateTournament(t *volleynet.FullTournament) error

	UpdateTournamentTeam(t *volleynet.TournamentTeam) error
	UpdateTournamentTeams(teams []volleynet.TournamentTeam) error
	NewTeam(t *volleynet.TournamentTeam) error
	NewTeams(teams []volleynet.TournamentTeam) error
	DeleteTeam(t *volleynet.TournamentTeam) error
	TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error)

	SearchPlayers() ([]volleynet.Player, error)
	AllPlayers() ([]volleynet.Player, error)
	NewPlayer(p *volleynet.Player) error
	Player(id int) (*volleynet.Player, error)
	UpdatePlayer(p *volleynet.Player) error
	Ladder(gender string) ([]volleynet.Player, error)
}

// VolleynetServiceImpl implements VolleynetService
type VolleynetServiceImpl struct {
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

func (s *VolleynetServiceImpl) Tournament(tournamentID int) (*volleynet.FullTournament, error) {
	row := s.DB.QueryRow(tournamentSelectSQL, tournamentID)

	return scanTournament(row)
}

func (s *VolleynetServiceImpl) AllTournaments() ([]volleynet.FullTournament, error) {
	return scanTournaments(s.DB, tournamentsBaseSelectSQL)
}

func (s *VolleynetServiceImpl) SeasonTournaments(season int) ([]volleynet.FullTournament, error) {
	return scanTournaments(s.DB, tournamentsSeasonSelectSQL, season)
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
	t.max_teams,
	t.end_registration,
	t.organiser,
	t.phone,
	t.email,
	t.web,
	t.current_points,
	t.live_scoring_link,
	t.loc_lat,
	t.loc_lon,
	t.season,
	t.signedup_teams
FROM volleynet_tournaments t`

	tournamentsSeasonSelectSQL = tournamentsBaseSelectSQL + " WHERE t.season = ?"

	tournamentsFilterSelectSQL = tournamentsBaseSelectSQL +
		" WHERE t.gender = ? AND t.league = ? AND t.season = ?"

	tournamentSelectSQL = tournamentsBaseSelectSQL + " WHERE t.id = ?"

	tournamentsInsertSQL = `
INSERT INTO volleynet_tournaments
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
	max_teams,
	end_registration,
	organiser,
	phone,
	email,
	web,
	current_points,
	live_scoring_link,
	loc_lat,
	loc_lon,
	season,
	signedup_teams
)
VALUES
(
	?,
	CURRENT_TIMESTAMP,
	CURRENT_TIMESTAMP,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?
)`

	tournamentsUpdateSQL = `
UPDATE volleynet_tournaments SET
	updated_at = CURRENT_TIMESTAMP,
	gender = ?,
	start = ?,
	end = ?,
	name = ?,
	league = ?,
	link = ?,
	entry_link = ?,
	status = ?,
	registration_open = ?,
	location = ?,
	html_notes = ?,
	mode = ?,
	max_points = ?,
	min_teams = ?,
	max_teams = ?,	
	end_registration = ?,
	organiser = ?,
	phone = ?,
	email = ?,
	web = ?,
	current_points = ?,
	live_scoring_link = ?,
	loc_lat = ?,
	loc_lon = ?,
	season = ?,
	signedup_teams = ?
WHERE id = ?`
)

func (s *VolleynetServiceImpl) GetTournaments(gender, league string, season int) ([]volleynet.FullTournament, error) {
	return scanTournaments(s.DB, tournamentsFilterSelectSQL, gender, league, season)
}

func (s *VolleynetServiceImpl) NewTournament(t *volleynet.FullTournament) error {
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

func (s *VolleynetServiceImpl) UpdateTournamentTeam(t *volleynet.TournamentTeam) error {
	result, err := s.DB.Exec(
		volleynetTeamsUpdateSQL,
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

func (s *VolleynetServiceImpl) UpdateTournament(t *volleynet.FullTournament) error {
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

const (
	volleynetTeamsSelectSQL = `
SELECT
	t.volleynet_tournament_id,
	t.volleynet_player_1_id,
	p1.first_name,
	p1.last_name,
	p1.total_points,
	p1.country_union,
	p1.birthday,
	p1.license,
	p1.gender,
	t.volleynet_player_2_id,
	p2.first_name,
	p2.last_name,
	p2.total_points,
	p2.country_union,
	p2.birthday,
	p2.license,
	p2.gender,
	t.rank,
	t.seed,
	t.total_points,
	t.won_points,
	t.prize_money,
	t.deregistered
FROM volleynet_tournament_teams t
JOIN volleynet_players p1 on p1.id = t.volleynet_player_1_id
JOIN volleynet_players p2 on p2.id = t.volleynet_player_2_id
WHERE t.volleynet_tournament_id = ?`

	volleynetTeamsUpdateSQL = `
UPDATE volleynet_tournament_teams SET
	rank = ?,
	seed = ?,
	total_points = ?,
	won_points = ?,
	prize_money = ?,
	deregistered = ?
WHERE volleynet_tournament_id = ? AND volleynet_player_1_id = ? AND volleynet_player_2_id = ?`

	volleynetTeamsDeleteSQL = `
DELETE FROM volleynet_tournament_teams
WHERE volleynet_tournament_id = ? 
	AND volleynet_player_1_id = ?
	AND volleynet_player_2_id = ?`

	volleynetTeamsInsertSQL = `
INSERT INTO volleynet_tournament_teams
(
	volleynet_tournament_id,
	volleynet_player_1_id,
	volleynet_player_2_id,
	rank,
	seed,
	total_points,
	won_points,
	prize_money,
	deregistered
)
VALUES
(
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?
)`
)

func (s *VolleynetServiceImpl) NewTeam(t *volleynet.TournamentTeam) error {
	_, err := s.DB.Exec(volleynetTeamsInsertSQL,
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

func (s *VolleynetServiceImpl) NewTeams(teams []volleynet.TournamentTeam) error {
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

func (s *VolleynetServiceImpl) UpdateTournamentTeams(teams []volleynet.TournamentTeam) error {
	for _, t := range teams {
		if err := s.UpdateTournamentTeam(&t); err != nil {
			return err
		}
	}

	return nil
}

func (s *VolleynetServiceImpl) DeleteTeam(t *volleynet.TournamentTeam) error {
	_, err := s.DB.Exec(volleynetTeamsDeleteSQL, t.TournamentID, t.Player1.ID, t.Player2.ID)

	return err
}

func (s *VolleynetServiceImpl) TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error) {
	return scanTournamentTeams(s.DB, volleynetTeamsSelectSQL, tournamentID)
}

const (
	volleynetBasePlayersSelectSQL = `
SELECT
	p.id,
	p.first_name,
	p.last_name,
	p.birthday,
	p.gender,
	p.total_points,
	p.rank,
	p.club,
	p.country_union,
	p.license
FROM volleynet_players p`

	volleynetPlayerSelectSQL  = volleynetBasePlayersSelectSQL + " WHERE p.id = ?"
	volleynetPlayersSelectSQL = volleynetBasePlayersSelectSQL

	volleynetPlayerLadderSelectSQL = volleynetBasePlayersSelectSQL +
		" WHERE p.rank > 0 AND p.gender = ? ORDER BY p.rank"

	volleynetPlayersUpdateSQL = `
UPDATE volleynet_players SET
	updated_at = CURRENT_TIMESTAMP,
	first_name = ?,
	last_name = ?,
	birthday = ?,
	gender = ?,
	total_points = ?,
	rank = ?,
	club = ?,
	country_union = ?,
	license = ?
WHERE id = ?`

	volleynetPlayersInsertSQL = `
INSERT INTO volleynet_players
(
	id,
	created_at,
	updated_at,
	first_name,
	last_name,
	birthday,
	gender,
	total_points,
	rank,
	club,
	country_union,
	license
)
VALUES
(
	?,
	CURRENT_TIMESTAMP,
	CURRENT_TIMESTAMP,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?
)`
)

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

func (s *VolleynetServiceImpl) Ladder(gender string) ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, volleynetPlayerLadderSelectSQL, gender)
}

func (s *VolleynetServiceImpl) AllPlayers() ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, volleynetPlayersSelectSQL)
}

func (s *VolleynetServiceImpl) Player(id int) (*volleynet.Player, error) {
	row := s.DB.QueryRow(
		volleynetPlayerSelectSQL,
		id,
	)

	return scanVolleynetPlayer(row)
}

func (s *VolleynetServiceImpl) NewPlayer(p *volleynet.Player) error {
	_, err := s.DB.Exec(volleynetPlayersInsertSQL,
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

func (s *VolleynetServiceImpl) UpdatePlayer(p *volleynet.Player) error {
	result, err := s.DB.Exec(
		volleynetPlayersUpdateSQL,
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

func (s *VolleynetServiceImpl) SearchPlayers() ([]volleynet.Player, error) {
	return nil, nil
}
