package sqlite

import (
	"database/sql"
	"time"

	"github.com/raphi011/scores"
)

var _ scores.MatchRepository = &MatchRepository{}

type MatchRepository struct {
	DB *sql.DB
}

func (s *MatchRepository) Delete(matchID uint) error {
	_, err := s.DB.Exec(query("match/update-delete"), matchID)

	return err
}

func (s *MatchRepository) Create(match *scores.Match) (*scores.Match, error) {
	result, err := s.DB.Exec(query("match/insert"),
		match.Group.ID,
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

	ID, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return s.Match(uint(ID))
}

func (s *MatchRepository) Match(ID uint) (*scores.Match, error) {
	row := s.DB.QueryRow(query("match/select-by-id"), ID)

	return scanMatch(row)
}

func (s *MatchRepository) MatchesAfter(after time.Time, count uint) (scores.Matches, error) {
	return scanMatches(s.DB, query("match/select-after"), after, count)
}

func (s *MatchRepository) GroupMatches(groupID uint, after time.Time, count uint) (scores.Matches, error) {
	return scanMatches(s.DB, query("match/select-group-matches"), groupID, after, count)
}

func (s *MatchRepository) PlayerMatches(playerID uint, after time.Time, count uint) (scores.Matches, error) {
	return scanMatches(s.DB, query("match/select-player-matches"), playerID, after, count)
}

func scanMatches(db *sql.DB, query string, args ...interface{}) (scores.Matches, error) {
	matches := scores.Matches{}
	rows, err := db.Query(query, args...)

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
		&m.GroupID,
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
