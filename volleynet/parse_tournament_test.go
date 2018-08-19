package volleynet

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func Test_upcoming_games(t *testing.T) {
	t.Skip()

	c := DefaultClient()
	tournaments, err := c.AllTournaments("M", "AMATEUR TOUR", "2018")

	if err != nil {
		t.Error(err)
	} else if len(tournaments) <= 0 {
		t.Error(errors.New("tournaments didn't return anything"))
	}
}

func Test_done_lms_tournament(t *testing.T) {
	response, _ := os.Open("testdata/done-tournament.html")

	tournament, err := parseFullTournament(response, Tournament{ID: 22750, Gender: "M", Status: "upcoming"})

	if err != nil {
		t.Errorf("parseFullTournament() err: %v", err)
	}

	date, _ := time.Parse("2006-01-02", "2018-08-18")

	compare := &FullTournament{
		Tournament: Tournament{
			Phase:  "LMS",
			Status: "done",
			ID:     22750,
			Gender: "M",
			Start:  date,
			End:    date,
		},
		Mode:          "Double Elimination 32er-Raster",
		Location:      "Auf der Schmelz 10 1150 Wien",
		MaxTeams:      32,
		SignedupTeams: 4,
		Organiser:     "Beachvolley Wien",
		Email:         "office@sportz.at",
		Web:           "www.beachvolleywien.at",
		Phone:         "0699/81 49 1124",
		Teams: []TournamentTeam{
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 4162, FirstName: "Christoph", LastName: "Haas"}, Gender: "M", CountryUnion: "STVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 20158, FirstName: "Dominik Karl", LastName: "Blaha"}, Gender: "M", CountryUnion: "STVV"},
				Deregistered: false,
				Rank:         1,
				TournamentID: 22750,
				WonPoints:    80,
			},
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 4523, FirstName: "Josef", LastName: "Buchner"}, Gender: "M", CountryUnion: "SVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 4179, FirstName: "Florian", LastName: "Tatra"}, Gender: "M", CountryUnion: "SVV"},
				Deregistered: false,
				Rank:         2,
				TournamentID: 22750,
				WonPoints:    70,
			},
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 13011, FirstName: "Daniel", LastName: "Wimmer"}, Gender: "M", CountryUnion: "SVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 44897, FirstName: "Maximilian", LastName: "Wieser"}, Gender: "M", CountryUnion: "SVV"},
				Deregistered: false,
				Rank:         3,
				TournamentID: 22750,
				WonPoints:    60,
			},
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 39947, FirstName: "Michael", LastName: "Seiser"}, Gender: "M", CountryUnion: "WVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 1991, FirstName: "Wolfgang", LastName: "Ertl"}, Gender: "M", CountryUnion: "WVV"},
				Deregistered: false,
				Rank:         4,
				TournamentID: 22750,
				WonPoints:    50,
			},
		},
	}

	if !cmp.Equal(tournament, compare) {
		t.Errorf("parseFullTournament() err: %+v", cmp.Diff(tournament, compare))
	}

}

func Test_done_tournament(t *testing.T) {
	response, _ := os.Open("testdata/done.html")

	_, err := parseFullTournament(response, Tournament{Gender: "M", Status: "upcoming"})

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}
}

func Test_upcoming_tournament(t *testing.T) {
	response, _ := os.Open("testdata/upcoming.html")

	_, err := parseFullTournament(response, Tournament{Gender: "M", Status: "upcoming"})

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}
}
