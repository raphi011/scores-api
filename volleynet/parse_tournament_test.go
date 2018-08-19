package volleynet

import (
	"os"
	"testing"

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

	compare := &FullTournament{
		Tournament: Tournament{
			Phase:  "LMS",
			Status: "done",
			ID:     22750,
			Gender: "M",
			Start:  mustParseDate("18.08.2018"),
			End:    mustParseDate("18.08.2018"),
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
		t.Errorf("parseFullTournament(): %+v", cmp.Diff(tournament, compare))
	}

}

func Test_done_tournament(t *testing.T) {
	response, _ := os.Open("testdata/done.html")

	tournament, err := parseFullTournament(response, Tournament{ID: 22228, Gender: "M", Status: "upcoming"})

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}

	// ignore notes
	tournament.HTMLNotes = ""

	compare := &FullTournament{
		Tournament: Tournament{
			ID:     22228,
			Phase:  "ABV Tour AMATEUR 1",
			Start:  mustParseDate("21.05.2018"),
			End:    mustParseDate("21.05.2018"),
			Gender: "M",
			Status: "done",
		},
		SignedupTeams: 3,
		MaxTeams:      32,
		Web:           "www.beachvolleyclub.at",
		Location:      "Arbeiterstrandbadstraße 87b 1210, Wien",
		Mode:          "Double Elimination 32er-Raster",
		MaxPoints:     339,
		Organiser:     "MOHAMED Tarek Mohie El-Din",
		Phone:         "0699 106 934 19",
		Email:         "tarek.mohamed@outlook.com",
		Teams: []TournamentTeam{
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 1043, FirstName: "Peter", LastName: "Dietl"}, Gender: "M", CountryUnion: "WVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 39947, FirstName: "Michael", LastName: "Seiser"}, Gender: "M", CountryUnion: "WVV"},
				Deregistered: false,
				Rank:         1,
				TournamentID: 22228,
				WonPoints:    50,
			},
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 11072, FirstName: "Christoph", LastName: "Brunnhofer"}, Gender: "M", CountryUnion: "STVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 27471, FirstName: "Christoph", LastName: "Mittendrein"}, Gender: "M", CountryUnion: "STVV"},
				Deregistered: false,
				Rank:         2,
				TournamentID: 22228,
				WonPoints:    45,
			},
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 36552, FirstName: "Dominik", LastName: "Koudela"}, Gender: "M", CountryUnion: "WVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 18348, FirstName: "Marian", LastName: "Schwinner"}, Gender: "M", CountryUnion: "NÖVV"},
				Deregistered: false,
				Rank:         3,
				TournamentID: 22228,
				WonPoints:    40,
			},
		},
	}

	if !cmp.Equal(tournament, compare) {
		t.Errorf("parseFullTournament(): %+v", cmp.Diff(tournament, compare))
	}
}

func Test_upcoming_tournament(t *testing.T) {
	response, _ := os.Open("testdata/upcoming.html")

	tournament, err := parseFullTournament(response, Tournament{ID: 22231, Gender: "M", Status: "upcoming"})

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}

	// ignore notes
	tournament.HTMLNotes = ""
	// ignoe currentpoints... TODO
	tournament.CurrentPoints = ""

	compare := &FullTournament{
		Tournament: Tournament{
			ID:     22231,
			Phase:  "ABV Tour AMATEUR 1",
			Start:  mustParseDate("31.05.2018"),
			End:    mustParseDate("31.05.2018"),
			Gender: "M",
			Status: "upcoming",
		},
		SignedupTeams: 3,
		MaxTeams:      32,
		Web:           "www.beachvolleyclub.at",
		Location:      "Arbeiterstrandbadstraße 87b 1210, Wien",
		Mode:          "Double Elimination 32er-Raster",
		MaxPoints:     339,
		Organiser:     "Wolfgang Ziegler",
		Phone:         "0699 106 934 19",
		Email:         "ziegler@sportz.at",
		Teams: []TournamentTeam{
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 10938, FirstName: "Sascha", LastName: "Kosatschow"}, TotalPoints: 303, License: "1", Gender: "M", CountryUnion: "STVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 27103, FirstName: "Johannes", LastName: "Pekar"}, TotalPoints: 177, License: "1", Gender: "M", CountryUnion: "STVV"},
				Deregistered: false,
				Seed:         1,
				TournamentID: 22231,
				TotalPoints:  480,
			},
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 18068, FirstName: "Lukas", LastName: "Wimmer"}, TotalPoints: 206, License: "1", Gender: "M", CountryUnion: "NÖVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 22590, FirstName: "Dominik", LastName: "Rieder"}, TotalPoints: 228, License: "1", Gender: "M", CountryUnion: "NÖVV"},
				Deregistered: false,
				Seed:         2,
				TournamentID: 22231,
				TotalPoints:  434,
			},
			TournamentTeam{
				Player1:      &Player{PlayerInfo: PlayerInfo{ID: 42649, FirstName: "Niels", LastName: "Antoni"}, TotalPoints: 162, License: "1", Gender: "M", CountryUnion: "WVV"},
				Player2:      &Player{PlayerInfo: PlayerInfo{ID: 21259, FirstName: "Liam", LastName: "Ochaya"}, TotalPoints: 257, License: "1", Gender: "M", CountryUnion: "STVV"},
				Deregistered: false,
				Seed:         3,
				TournamentID: 22231,
				TotalPoints:  419,
			},
		},
	}

	if !cmp.Equal(tournament, compare) {
		t.Errorf("parseFullTournament(): %+v", cmp.Diff(tournament, compare))
	}
}
