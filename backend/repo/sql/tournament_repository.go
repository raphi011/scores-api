package sql

import (
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// tournamentRepository implements VolleynetRepository interface.
type tournamentRepository struct {
	DB *sqlx.DB
}

// Get loads a tournament by its id.
func (s *tournamentRepository) Get(tournamentID int) (*volleynet.Tournament, error) {
	return s.scanOne("tournament/select-by-id", tournamentID)
}

// New creates a new tournament.
func (s *tournamentRepository) New(t *volleynet.Tournament) (*volleynet.Tournament, error) {
	err := s.exec("tournament/insert", t)

	return t, errors.Wrap(err, "insert tournament")
}

// NewBatch creates multiple new tournaments
func (s *tournamentRepository) NewBatch(t ...*volleynet.Tournament) error {
	err := s.exec("tournament/insert", t...)

	return errors.Wrap(err, "insert tournament")
}

// Update updates a tournament.
func (s *tournamentRepository) Update(t *volleynet.Tournament) error {
	err := update(s.DB, "tournament/update", t)

	return errors.Wrap(err, "update tournament")
}

// Filter loads all tournaments by season, league and gender.
func (s *tournamentRepository) Filter(
	seasons []int,
	leagues []string,
	formats []string) ([]*volleynet.Tournament, error) {

	tournaments, err := s.scan("tournament/select-by-filter",
		formats,
		leagues, 
		seasons,
	)

	return tournaments, errors.Wrap(err, "filtered tournaments")
}

func (s *tournamentRepository) scan(queryName string, args ...interface{}) (
	[]*volleynet.Tournament, error) {

	tournaments := []*volleynet.Tournament{}

	q, args, err := sqlx.In(loadQuery(s.DB, queryName), args...)

	if err != nil {
		return tournaments, errors.Wrap(err, "creating query")
	}

	q = s.DB.Rebind(q)

	rows, err := s.DB.Query(q, args...)

	if err != nil {
		return tournaments, mapError(err)
	}

	defer rows.Close()

	for rows.Next() {
		t := &volleynet.Tournament{}
		t.Teams = []*volleynet.TournamentTeam{}

		err := rows.Scan(
			&t.ID,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.Format,
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
			&t.Website,
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

func (s *tournamentRepository) exec(queryName string, entities ...*volleynet.Tournament) error {
	stmt, err := s.DB.PrepareNamed(namedQuery(s.DB, queryName))

	if err != nil {
		return mapError(err)
	}

	for _, entity := range entities {
		_, err := stmt.Exec(entity)

		if err != nil {
			return mapError(err)
		}
	}

	return nil
}

func (s *tournamentRepository) scanOne(query string, args ...interface{}) (
	*volleynet.Tournament, error) {

	tournaments, err := s.scan(query, args...)

	if err != nil {
		return nil, err
	}

	if len(tournaments) >= 1 {
		return tournaments[0], nil
	}

	return nil, scores.ErrNotFound
}