package sql

import (
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// TournamentRepository implements VolleynetRepository interface.
type TournamentRepository struct {
	DB *sqlx.DB
}

// Get loads a tournament by its id.
func (s *TournamentRepository) Get(tournamentID int) (*volleynet.FullTournament, error) {
	return s.scanOne("tournament/select-by-id", tournamentID)
}

// New creates a new tournament.
func (s *TournamentRepository) New(t *volleynet.FullTournament) (*volleynet.FullTournament, error) {
	_, err := exec(s.DB, "tournament/insert", t)

	return t, errors.Wrap(err, "insert tournament")
}

// Update updates a tournament.
func (s *TournamentRepository) Update(t *volleynet.FullTournament) error {
	err := update(s.DB, "tournament/update", t)

	return errors.Wrap(err, "update tournament")
}

// Filter loads all tournaments by season, league and gender.
func (s *TournamentRepository) Filter(
	seasons []int,
	leagues []string,
	formats []string) ([]*volleynet.FullTournament, error) {

	tournaments, err := s.scan("tournament/select-by-filter",
		formats,
		leagues, 
		seasons,
	)

	return tournaments, errors.Wrap(err, "filtered tournaments")
}

func (s *TournamentRepository) scan(queryName string, args ...interface{}) (
	[]*volleynet.FullTournament, error) {

	tournaments := []*volleynet.FullTournament{}

	q, args, err := sqlx.In(loadQuery(queryName), args...)

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
		t := &volleynet.FullTournament{}
		t.Teams = []volleynet.TournamentTeam{}

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

func (s *TournamentRepository) scanOne(query string, args ...interface{}) (
	*volleynet.FullTournament, error) {

	tournaments, err := s.scan(query, args...)

	if err != nil {
		return nil, err
	}

	if len(tournaments) >= 1 {
		return tournaments[0], nil
	}

	return nil, scores.ErrNotFound
}