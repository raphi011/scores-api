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
	return s.scanOne("tournament/select-by-id", tournamentID)
}

// All loads all tournaments
// Note: should only be used for debugging
func (s *TournamentRepository) All() ([]*volleynet.FullTournament, error) {
	tournaments, err := s.scan("tournament/select-all")

	return tournaments, errors.Wrap(err, "all tournaments")
}

// Season loads all tournaments of a season
func (s *TournamentRepository) Season(season int) ([]*volleynet.FullTournament, error) {
	tournaments, err := s.scan("tournament/select-by-season", season)

	return tournaments, errors.Wrap(err, "season tournaments")
}

// New creates a new tournament
func (s *TournamentRepository) New(t *volleynet.FullTournament) error {
	_, err := exec(s.DB, "tournament/insert", t)

	return errors.Wrap(err, "insert tournament")
}

// UpdatedSince gets all tournaments that were updated after a certain time
func (s *TournamentRepository) UpdatedSince(updatedSince time.Time) ([]*volleynet.FullTournament, error) {
	tournaments, err := s.scan("tournament/select-by-updated-since", updatedSince)

	return tournaments, errors.Wrap(err, "updated since tournaments")
}

// Update updates a tournament
func (s *TournamentRepository) Update(t *volleynet.FullTournament) error {
	err := update(s.DB, "tournament/update", t)

	return errors.Wrap(err, "update tournament")
}

// Filter loads all tournaments by season, league and gender
func (s *TournamentRepository) Filter(
	seasons []int,
	leagues []string,
	genders []string) ([]*volleynet.FullTournament, error) {

	tournaments, err := s.scan("tournament/select-by-filter", genders, leagues, seasons)

	return tournaments, errors.Wrap(err, "filtered tournaments")
}

func (s *TournamentRepository) scan(queryName string, args ...interface{}) ([]*volleynet.FullTournament, error) {
	tournaments := []*volleynet.FullTournament{}

	q := query(s.DB, queryName)

	rows, err := s.DB.Query(q, args...)

	if err != nil {
		return tournaments, mapError(err)
	}

	defer rows.Close()

	for rows.Next() {
		t := &volleynet.FullTournament{}
		t.Teams = []volleynet.TournamentTeam{}

		err := rows.Scan(
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

		if err != nil {
			return tournaments, mapError(err)
		}

		tournaments = append(tournaments, t)
	}

	return tournaments, nil
}

func (s *TournamentRepository) scanOne(query string, args ...interface{}) (
	*volleynet.FullTournament, error) {

	tournaments, err := s.scan(query, args)

	if err != nil {
		return nil, err
	}

	if len(tournaments) >= 1 {
		return tournaments[0], nil
	}

	return nil, nil
}