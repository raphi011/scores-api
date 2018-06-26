package sqlite

import (
	"database/sql"
	"github.com/pkg/errors"

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
	)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *VolleynetService) Tournament(tournamentID int) (*volleynet.FullTournament, error) {
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
		t.season
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
			$25,
			$26
		)
	`

	tournamentsUpdateSQL = `
		UPDATE volleynetTournaments SET
			updated_at = CURRENT_TIMESTAMP,
			gender = $1,
			start = $2,
			end = $3,
			name = $4,
			league = $5,
			link = $6,
			entry_link = $7,
			status = $8,
			registration_open = $9,
			location = $10,
			html_notes = $11,
			mode = $12,
			max_points = $13,
			min_teams = $14,
			max_teams = $15,	
			end_registration = $16,
			organiser = $17,
			phone = $18,
			email = $19,
			web = $20,
			current_points = $21,
			live_scoring_link = $22,
			loc_lat = $23,
			loc_lon = $24,
			season = $25
		WHERE id = $26
	`
)

func (s *VolleynetService) GetTournaments(gender, league string, season int) ([]volleynet.FullTournament, error) {
	return scanTournaments(s.DB, tournamentsFilterSelectSQL, gender, league, season)
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
	)

	return err
}

func (s *VolleynetService) UpdateTournamentTeam(t *volleynet.TournamentTeam) error {
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

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected != 1 {
		return errors.New("Team not found")
	}

	return nil
}

func (s *VolleynetService) UpdateTournamentTeams(teams []volleynet.TournamentTeam) error {
	for _, t := range teams {
		if err := s.UpdateTournamentTeam(&t); err != nil {
			return err
		}
	}

	return nil
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
		FROM volleynetTournamentTeams t
		JOIN volleynetPlayers p1 on p1.id = t.volleynet_player_1_id
		JOIN volleynetPlayers p2 on p2.id = t.volleynet_player_2_id
		WHERE t.volleynet_tournament_id = $1	
	`

	volleynetTeamsUpdateSQL = `
		UPDATE volleynetTournamentTeams SET
			rank = $1,
			seed = $2,
			total_points = $3,
			won_points = $4,
			prize_money = $5,
			deregistered = $6
		WHERE volleynet_tournament_id = $7 AND volleynet_player_1_id = $8 AND volleynet_player_2_id = $9
	`

	volleynetTeamsInsertSQL = `
		INSERT INTO volleynetTournamentTeams
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
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
		)
	`
)

func (s *VolleynetService) NewTeam(t *volleynet.TournamentTeam) error {
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

func (s *VolleynetService) NewTeams(teams []volleynet.TournamentTeam) error {
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

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *VolleynetService) TournamentTeams(tournamentID int) ([]volleynet.TournamentTeam, error) {
	return scanTournamentTeams(s.DB, volleynetTeamsSelectSQL, tournamentID)
}

const (
	volleynetBasePlayersSelectSQL = `
		SELECT
			p.id,
			p.first_name,
			p.last_name,
			p.login,
			p.birthday,
			p.gender,
			p.total_points,
			p.rank,
			p.club,
			p.country_union,
			p.license,
			p.total_points
		FROM volleynetPlayers p
	`

	volleynetPlayerSelectSQL  = volleynetBasePlayersSelectSQL + " WHERE p.id = $1"
	volleynetPlayersSelectSQL = volleynetBasePlayersSelectSQL

	volleynetPlayersUpdateSQL = `
		UPDATE volleynetPlayers SET
			updated_at = CURRENT_TIMESTAMP,
			first_name = $1,
			last_name = $2,
			login = $3,
			birthday = $4,
			gender = $5,
			total_points = $6,
			rank = $7,
			club = $8,
			country_union = $9,
			license = $10,
			total_points = $11
		WHERE id = $12
	`

	volleynetPlayersInsertSQL = `
		INSERT INTO volleynetPlayers
		(
			id,
			created_at,
			updated_at,
			first_name,
			last_name,
			login,
			birthday,
			gender,
			total_points,
			rank,
			club,
			country_union,
			license,
			total_points
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
			$12
		)
	`
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
		&p.Login,
		&p.Birthday,
		&p.Gender,
		&p.TotalPoints,
		&p.Rank,
		&p.Club,
		&p.CountryUnion,
		&p.License,
		&p.TotalPoints,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *VolleynetService) AllPlayers() ([]volleynet.Player, error) {
	return scanVolleynetPlayers(s.DB, volleynetPlayersSelectSQL)
}

func (s *VolleynetService) NewPlayer(p *volleynet.Player) error {
	_, err := s.DB.Exec(volleynetPlayersInsertSQL,
		p.ID,
		p.FirstName,
		p.LastName,
		p.Login,
		p.Birthday,
		p.Gender,
		p.TotalPoints,
		p.Rank,
		p.Club,
		p.CountryUnion,
		p.License,
		p.TotalPoints,
	)

	return err
}

func (s *VolleynetService) UpdatePlayer(p *volleynet.Player) error {
	result, err := s.DB.Exec(
		volleynetPlayersUpdateSQL,
		p.FirstName,
		p.LastName,
		p.Login,
		p.Birthday,
		p.Gender,
		p.TotalPoints,
		p.Rank,
		p.Club,
		p.CountryUnion,
		p.License,
		p.TotalPoints,
		p.ID)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected != 1 {
		return errors.New("Player not found")
	}

	return nil
}
