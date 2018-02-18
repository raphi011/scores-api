package sqlite

import (
	"database/sql"
	scores "scores-backend"
	"time"
)

var _ scores.MatchService = &MatchService{}

type MatchService struct {
	DB *sql.DB
}

func (s *MatchService) Delete(matchID uint) error {
	_, err := s.DB.Exec("UPDATE matches SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1", matchID)

	return err
}

const (
	matchesInsertSQL = `
		INSERT INTO matches
		(
			created_at,
			team1_player1_id,
			team1_player2_id,
			team2_player1_id,
			team2_player2_id,
			score_team1,
			score_team2,
			created_by_user_id
		)
		VALUES
		(
			CURRENT_TIMESTAMP,
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
	`
)

func (s *MatchService) Create(match *scores.Match) (*scores.Match, error) {
	result, err := s.DB.Exec(matchesInsertSQL,
		match.Team1.Player1ID,
		match.Team1.Player2ID,
		match.Team2.Player1ID,
		match.Team2.Player2ID,
		match.ScoreTeam1,
		match.ScoreTeam2,
		match.CreatedBy.ID,
	)

	if err != nil {
		return nil, err
	}

	ID, _ := result.LastInsertId()

	return s.Match(uint(ID))
}

const (
	matchesBaseSelectSQL = `
	SELECT
		m.id,
		m.created_at,
		m.team1_player1_id,
		p1.name as team1_player1_name,
		COALESCE(u1.profile_image_url, "") as team1_player1_image_url,
		m.team1_player2_id,
		p2.name as team1_player2_name,
		COALESCE(u2.profile_image_url, "") as team1_player2_image_url,
		m.team2_player1_id,
		p3.name as team2_player1_name,
		COALESCE(u3.profile_image_url, "") as team2_player1_image_url,
		m.team2_player2_id,
		p4.name as team2_player2_name,
		COALESCE(u4.profile_image_url, "") as team2_player2_image_url,
		m.score_team1,
		m.score_team2,
		m.created_by_user_id
	FROM matches m
	JOIN players p1 on m.team1_player1_id = p1.id
	JOIN players p2 on m.team1_player2_id = p2.id
	JOIN players p3 on m.team2_player1_id = p3.id
	JOIN players p4 on m.team2_player2_id = p4.id
	LEFT JOIN users u1 on p1.user_id = u1.id
	LEFT JOIN users u2 on p2.user_id = u2.id
	LEFT JOIN users u3 on p3.user_id = u3.id
	LEFT JOIN users u4 on p4.user_id = u4.id
	WHERE m.deleted_at is null
`

	matchesSelectSQL = matchesBaseSelectSQL +
		"AND m.created_at < $1" +
		matchesOrderBySQL +
		" LIMIT $2"

	matchesOrderBySQL = " ORDER BY m.created_at DESC"

	matchesByPlayerSelectSQL = matchesBaseSelectSQL + ` 
 	AND (m.team1_player1_id = $1 OR 
			 m.team1_player2_id = $1 OR 
			 m.team2_player1_id = $1 OR 
			 m.team2_player2_id = $1)` +
		" AND m.created_at < $2" +
		matchesOrderBySQL +
		" LIMIT $3"

	matchSelectSQL = matchesBaseSelectSQL + " and m.id = $1"
)

func scanMatch(scanner scan) (*scores.Match, error) {
	m := &scores.Match{
		Team1: &scores.Team{
			Player1: &scores.Player{},
			Player2: &scores.Player{},
		},
		Team2: &scores.Team{
			Player1: &scores.Player{},
			Player2: &scores.Player{},
		},
		CreatedBy: &scores.User{},
	}

	err := scanner.Scan(
		&m.ID,
		&m.CreatedAt,
		&m.Team1.Player1.ID,
		&m.Team1.Player1.Name,
		&m.Team1.Player1.ProfileImageURL,
		&m.Team1.Player2.ID,
		&m.Team1.Player2.Name,
		&m.Team1.Player2.ProfileImageURL,
		&m.Team2.Player1.ID,
		&m.Team2.Player1.Name,
		&m.Team2.Player1.ProfileImageURL,
		&m.Team2.Player2.ID,
		&m.Team2.Player2.Name,
		&m.Team2.Player2.ProfileImageURL,
		&m.ScoreTeam1,
		&m.ScoreTeam2,
		&m.CreatedBy.ID,
	)

	if err != nil {
		return nil, err
	}

	m.Team1.Player1ID = m.Team1.Player1.ID
	m.Team1.Player2ID = m.Team1.Player2.ID
	m.Team2.Player1ID = m.Team2.Player1.ID
	m.Team2.Player2ID = m.Team2.Player2.ID

	return m, nil
}

func (s *MatchService) Match(ID uint) (*scores.Match, error) {
	row := s.DB.QueryRow(matchSelectSQL, ID)

	match, err := scanMatch(row)

	return match, err
}

func (s *MatchService) Matches(after time.Time, count uint) (scores.Matches, error) {
	matches := scores.Matches{}

	rows, err := s.DB.Query(matchesSelectSQL, after, count)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		match, err := scanMatch(rows)

		if err != nil {
			return nil, err
		}

		matches = append(matches, *match)
	}

	return matches, nil
}

func (s *MatchService) PlayerMatches(playerID uint, after time.Time, count uint) (scores.Matches, error) {
	matches := scores.Matches{}

	rows, err := s.DB.Query(matchesByPlayerSelectSQL, playerID, after, count)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		match, err := scanMatch(rows)

		if err != nil {
			return nil, err
		}

		matches = append(matches, *match)
	}

	return matches, nil
}
