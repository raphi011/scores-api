package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sql/crud"
	"github.com/raphi011/scores/volleynet"
)

var _ repo.TeamRepository = &teamRepository{}

type teamRepository struct {
	DB *sqlx.DB
}

// New creates a new team.
func (s *teamRepository) New(t *volleynet.TournamentTeam) (*volleynet.TournamentTeam, error) {
	err := crud.Create(s.DB,  "team/insert", t)

	return t, errors.Wrap(err, "insert team")
}

// NewBatch creates a new team.
func (s *teamRepository) NewBatch(teams ...*volleynet.TournamentTeam) error {
	ts := make([]repo.Tracked, len(teams))

	for i, t := range teams {
		ts[i] = t
	}

	err := crud.Create(s.DB,  "team/insert", ts...)

	return errors.Wrap(err, "batch insert team")
}

// Update updates a tournament team.
func (s *teamRepository) Update(t *volleynet.TournamentTeam) error {
	err := crud.Update(s.DB, "team/update", t)

	return errors.Wrap(err, "update team")
}

// UpdateBatch updates a tournament team.
func (s *teamRepository) UpdateBatch(teams ...*volleynet.TournamentTeam) error {
	ts := make([]repo.Tracked, len(teams))

	for i, t := range teams {
		ts[i] = t
	}

	err := crud.Update(s.DB, "team/update", ts...)

	return errors.Wrap(err, "batch update team")
}

// Delete deletes a team.
func (s *teamRepository) Delete(t *volleynet.TournamentTeam) error {
	err := crud.Delete(s.DB, "team/delete", t)

	return errors.Wrap(err, "delete team")
}

// ByTournament loads all teams of a tournament.
func (s *teamRepository) ByTournament(tournamentID int) (
	[]*volleynet.TournamentTeam, error) {

	teams := []*volleynet.TournamentTeam{}
	err := crud.Read(s.DB, "team/select-by-tournament-id", &teams, tournamentID)

	return teams, errors.Wrap(err, "delete team")
}

// func (s *teamRepository) scan(queryName string, args ...interface{}) (
// 	[]*volleynet.TournamentTeam, error) {

// 	teams := []*volleynet.TournamentTeam{}

// 	q := query(s.DB, queryName)

// 	rows, err := s.DB.Query(q, args...)

// 	if err != nil {
// 		return nil, mapError(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		t := &volleynet.TournamentTeam{}
// 		t.Player1 = &volleynet.Player{}
// 		t.Player2 = &volleynet.Player{}

// 		err = rows.Scan(
// 			&t.TournamentID,
// 			&t.Player1.ID,
// 			&t.Player1.FirstName,
// 			&t.Player1.LastName,
// 			&t.Player1.TotalPoints,
// 			&t.Player1.CountryUnion,
// 			&t.Player1.Birthday,
// 			&t.Player1.License,
// 			&t.Player1.Gender,
// 			&t.Player2.ID,
// 			&t.Player2.FirstName,
// 			&t.Player2.LastName,
// 			&t.Player2.TotalPoints,
// 			&t.Player2.CountryUnion,
// 			&t.Player2.Birthday,
// 			&t.Player2.License,
// 			&t.Player2.Gender,
// 			&t.Rank,
// 			&t.Seed,
// 			&t.TotalPoints,
// 			&t.WonPoints,
// 			&t.PrizeMoney,
// 			&t.Deregistered,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		teams = append(teams, t)
// 	}

// 	return teams, nil
// }
