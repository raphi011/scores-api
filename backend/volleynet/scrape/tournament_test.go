package scrape

import (
	"os"
	"testing"
	"time"

	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/volleynet"
)

var tournamentTests = []struct {
	file       string
	now        time.Time
	tournament volleynet.TournamentInfo
	out        *volleynet.Tournament
}{
	{
		"../testdata/done-no-teams.html",
		mustParseDate("02.05.2018"),
		volleynet.TournamentInfo{ID: 21808, Gender: "M", Status: "upcoming"},
		&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
				ID:     21808,
				Phase:  "ABV Tour AMATEUR 2",
				Start:  mustParseDate("01.05.2018"),
				End:    mustParseDate("01.05.2018"),
				Gender: "M",
				Status: "done",
			},
			MaxTeams:  8,
			MinTeams:  8,
			MaxPoints: -1,
			Website:   "www.surfworldcup.at/sport/beach-summer-games/",
			Location:  "Seegelände 7100 Neusiedl am See",
			Mode:      "Double Elimination 8er-Raster",
			Organiser: "Beach Summer Games",
			Phone:     "0680 1412388",
			Email:     "office@beachsummergames.com",
			Teams:     []*volleynet.TournamentTeam{},
		},
	},
	{
		"../testdata/22764-done.html",
		mustParseDate("02.09.2018"),
		volleynet.TournamentInfo{ID: 22764, Gender: "M", Status: "upcoming"},
		&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
				ID:     22764,
				Phase:  "ABV Tour AMATEUR 1",
				Start:  mustParseDate("01.09.2018"),
				End:    mustParseDate("01.09.2018"),
				Gender: "M",
				Status: "done",
			},
			SignedupTeams: 10,
			MaxTeams:      16,
			MinTeams:      4,
			Website:       "www.12ndr.at",
			Location:      "Beachvolleyballplatz Stockerau - Pestalozzigasse 1  2000 Stockerau",
			Mode:          "Double Elimination 16er-Raster",
			MaxPoints:     540,
			Organiser:     "Michael Gahler",
			Phone:         "+43 664 6122639",
			Email:         "Vorstand@12ndr.at",
			Teams: []*volleynet.TournamentTeam{
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 22606, FirstName: "Richard", LastName: "Bosse", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 41275, FirstName: "Raphael", LastName: "Gruber", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       1,
					TournamentID: 22764,
					WonPoints:    36,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 28725, FirstName: "Alexander", LastName: "Jäger", Gender: "M", CountryUnion: "WVV"},
					Player2:      &volleynet.Player{ID: 20436, FirstName: "Bernhard", LastName: "Metzger", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       2,
					TournamentID: 22764,
					WonPoints:    33,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 27201, FirstName: "Alexander", LastName: "Jirgal", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 17623, FirstName: "Luca Maxim", LastName: "Wojnar", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       3,
					TournamentID: 22764,
					WonPoints:    29,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 10198, FirstName: "Markus", LastName: "Mayer", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 42403, FirstName: "Constantin", LastName: "Schieber", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       4,
					TournamentID: 22764,
					WonPoints:    26,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 22913, FirstName: "Herbert", LastName: "Eminger", Gender: "M", CountryUnion: "WVV"},
					Player2:      &volleynet.Player{ID: 33125, FirstName: "Stefan", LastName: "Handschmann", Gender: "M", CountryUnion: "WVV"},
					Deregistered: false,
					Result:       5,
					TournamentID: 22764,
					WonPoints:    22,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 44906, FirstName: "Reinhard", LastName: "Weiskirchner", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 13788, FirstName: "Michael", LastName: "Gahler", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       5,
					TournamentID: 22764,
					WonPoints:    22,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 18427, FirstName: "Michael", LastName: "Haas", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 39945, FirstName: "Werner", LastName: "Schmid", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       7,
					TournamentID: 22764,
					WonPoints:    18,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 36540, FirstName: "Andreas", LastName: "Zelinka", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 51104, FirstName: "Simon", LastName: "Sladek", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       7,
					TournamentID: 22764,
					WonPoints:    18,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 55789, FirstName: "Martin", LastName: "Gschweidl", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 36557, FirstName: "Stefan", LastName: "Müller", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       9,
					TournamentID: 22764,
					WonPoints:    15,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 55596, FirstName: "Thomas", LastName: "Müllner", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 43098, FirstName: "Sebastian", LastName: "Lechner", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       9,
					TournamentID: 22764,
					WonPoints:    15,
				},
			},
		},
	},
	{

		"../testdata/done.html",
		mustParseDate("22.05.2018"),
		volleynet.TournamentInfo{ID: 22228, Gender: "M", Status: "upcoming"},
		&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
				ID:     22228,
				Phase:  "ABV Tour AMATEUR 1",
				Start:  mustParseDate("21.05.2018"),
				End:    mustParseDate("21.05.2018"),
				Gender: "M",
				Status: "done",
			},
			SignedupTeams: 3,
			MaxTeams:      32,
			Website:       "www.beachvolleyclub.at",
			Location:      "Arbeiterstrandbadstraße 87b 1210, Wien",
			Mode:          "Double Elimination 32er-Raster",
			MaxPoints:     339,
			Organiser:     "MOHAMED Tarek Mohie El-Din",
			Phone:         "0699 106 934 19",
			Email:         "tarek.mohamed@outlook.com",
			Teams: []*volleynet.TournamentTeam{
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 1043, FirstName: "Peter", LastName: "Dietl", Gender: "M", CountryUnion: "WVV"},
					Player2:      &volleynet.Player{ID: 39947, FirstName: "Michael", LastName: "Seiser", Gender: "M", CountryUnion: "WVV"},
					Deregistered: false,
					Result:       1,
					TournamentID: 22228,
					WonPoints:    50,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 11072, FirstName: "Christoph", LastName: "Brunnhofer", Gender: "M", CountryUnion: "STVV"},
					Player2:      &volleynet.Player{ID: 27471, FirstName: "Christoph", LastName: "Mittendrein", Gender: "M", CountryUnion: "STVV"},
					Deregistered: false,
					Result:       2,
					TournamentID: 22228,
					WonPoints:    45,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 36552, FirstName: "Dominik", LastName: "Koudela", Gender: "M", CountryUnion: "WVV"},
					Player2:      &volleynet.Player{ID: 18348, FirstName: "Marian", LastName: "Schwinner", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       3,
					TournamentID: 22228,
					WonPoints:    40,
				},
			},
		},
	},
	{
		"../testdata/done-tournament.html",
		mustParseDate("19.08.2018"),
		volleynet.TournamentInfo{ID: 22750, Gender: "M", Status: "upcoming"},
		&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
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
			Website:       "www.beachvolleywien.at",
			Phone:         "0699/81 49 1124",
			Teams: []*volleynet.TournamentTeam{
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 4162, FirstName: "Christoph", LastName: "Haas", Gender: "M", CountryUnion: "STVV"},
					Player2:      &volleynet.Player{ID: 20158, FirstName: "Dominik Karl", LastName: "Blaha", Gender: "M", CountryUnion: "STVV"},
					Deregistered: false,
					Result:       1,
					TournamentID: 22750,
					WonPoints:    80,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 4523, FirstName: "Josef", LastName: "Buchner", Gender: "M", CountryUnion: "SVV"},
					Player2:      &volleynet.Player{ID: 4179, FirstName: "Florian", LastName: "Tatra", Gender: "M", CountryUnion: "SVV"},
					Deregistered: false,
					Result:       2,
					TournamentID: 22750,
					WonPoints:    70,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 13011, FirstName: "Daniel", LastName: "Wimmer", Gender: "M", CountryUnion: "SVV"},
					Player2:      &volleynet.Player{ID: 44897, FirstName: "Maximilian", LastName: "Wieser", Gender: "M", CountryUnion: "SVV"},
					Deregistered: false,
					Result:       3,
					TournamentID: 22750,
					WonPoints:    60,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 39947, FirstName: "Michael", LastName: "Seiser", Gender: "M", CountryUnion: "WVV"},
					Player2:      &volleynet.Player{ID: 1991, FirstName: "Wolfgang", LastName: "Ertl", Gender: "M", CountryUnion: "WVV"},
					Deregistered: false,
					Result:       4,
					TournamentID: 22750,
					WonPoints:    50,
				},
			},
		},
	},

	{
		"../testdata/upcoming.html",
		mustParseDate("19.08.2018"),
		volleynet.TournamentInfo{ID: 22231, Gender: "M", Status: "upcoming"},

		&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
				ID:     22231,
				Phase:  "ABV Tour AMATEUR 1",
				Start:  mustParseDate("31.05.2018"),
				End:    mustParseDate("31.05.2018"),
				Gender: "M",
				Status: "upcoming",
			},
			SignedupTeams: 3,
			MaxTeams:      32,
			Website:       "www.beachvolleyclub.at",
			Location:      "Arbeiterstrandbadstraße 87b 1210, Wien",
			Mode:          "Double Elimination 32er-Raster",
			MaxPoints:     339,
			Organiser:     "Wolfgang Ziegler",
			Phone:         "0699 106 934 19",
			Email:         "ziegler@sportz.at",
			Teams: []*volleynet.TournamentTeam{
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 10938, FirstName: "Sascha", LastName: "Kosatschow", TotalPoints: 303, License: "1", Gender: "M", CountryUnion: "STVV"},
					Player2:      &volleynet.Player{ID: 27103, FirstName: "Johannes", LastName: "Pekar", TotalPoints: 177, License: "1", Gender: "M", CountryUnion: "STVV"},
					Deregistered: false,
					Seed:         1,
					TournamentID: 22231,
					TotalPoints:  480,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 18068, FirstName: "Lukas", LastName: "Wimmer", TotalPoints: 206, License: "1", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 22590, FirstName: "Dominik", LastName: "Rieder", TotalPoints: 228, License: "1", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Seed:         2,
					TournamentID: 22231,
					TotalPoints:  434,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 42649, FirstName: "Niels", LastName: "Antoni", TotalPoints: 162, License: "1", Gender: "M", CountryUnion: "WVV"},
					Player2:      &volleynet.Player{ID: 21259, FirstName: "Liam", LastName: "Ochaya", TotalPoints: 257, License: "1", Gender: "M", CountryUnion: "STVV"},
					Deregistered: false,
					Seed:         3,
					TournamentID: 22231,
					TotalPoints:  419,
				},
			},
		},
	},

	{
		"../testdata/22764-upcoming-wildcard.html",
		mustParseDate("02.09.2018"),
		volleynet.TournamentInfo{ID: 22764, Gender: "M", Status: "upcoming"},
		&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
				ID:     22764,
				Phase:  "ABV Tour AMATEUR 1",
				Start:  mustParseDate("01.09.2018"),
				End:    mustParseDate("01.09.2018"),
				Gender: "M",
				Status: "upcoming",
			},
			SignedupTeams: 11,
			MaxTeams:      16,
			MinTeams:      4,
			Website:       "www.12ndr.at",
			Location:      "Beachvolleyballplatz Stockerau - Pestalozzigasse 1  2000 Stockerau",
			Mode:          "Double Elimination 16er-Raster",
			MaxPoints:     540,
			Organiser:     "Michael Gahler",
			Phone:         "+43 664 6122639",
			Email:         "Vorstand@12ndr.at",
			Teams: []*volleynet.TournamentTeam{
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 22606, FirstName: "Richard", LastName: "Bosse", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 249},
					Player2:      &volleynet.Player{ID: 41275, FirstName: "Raphael", LastName: "Gruber", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 242},
					Deregistered: false,
					Seed:         1,
					TournamentID: 22764,
					TotalPoints:  491,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 27201, FirstName: "Alexander", LastName: "Jirgal", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 184},
					Player2:      &volleynet.Player{ID: 17623, FirstName: "Luca Maxim", LastName: "Wojnar", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 185},
					Deregistered: false,
					Seed:         2,
					TournamentID: 22764,
					TotalPoints:  369,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 54595, FirstName: "Maximilian", LastName: "Rauter", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 155},
					Player2:      &volleynet.Player{ID: 47755, FirstName: "Moritz", LastName: "Hörl", Gender: "M", CountryUnion: "SVV", License: "1", TotalPoints: 187},
					Deregistered: false,
					Seed:         3,
					TournamentID: 22764,
					TotalPoints:  342,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 44906, FirstName: "Reinhard", LastName: "Weiskirchner", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 157},
					Player2:      &volleynet.Player{ID: 13788, FirstName: "Michael", LastName: "Gahler", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 138},
					Deregistered: false,
					Seed:         4,
					TournamentID: 22764,
					TotalPoints:  295,
				},
				&volleynet.TournamentTeam{Player1: &volleynet.Player{ID: 28725, FirstName: "Alexander", LastName: "Jäger", Gender: "M", CountryUnion: "WVV", License: "1", TotalPoints: 214},
					Player2:      &volleynet.Player{ID: 20436, FirstName: "Bernhard", LastName: "Metzger", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 78},
					Deregistered: false,
					Seed:         5,
					TournamentID: 22764,
					TotalPoints:  292,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 22913, FirstName: "Herbert", LastName: "Eminger", Gender: "M", CountryUnion: "WVV", License: "1", TotalPoints: 102},
					Player2:      &volleynet.Player{ID: 33125, FirstName: "Stefan", LastName: "Handschmann", Gender: "M", CountryUnion: "WVV", License: "1", TotalPoints: 122},
					Deregistered: false,
					Seed:         6,
					TournamentID: 22764,
					TotalPoints:  224,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 10198, FirstName: "Markus", LastName: "Mayer", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 42},
					Player2:      &volleynet.Player{ID: 42403, FirstName: "Constantin", LastName: "Schieber", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 153},
					Deregistered: false,
					Seed:         7,
					TournamentID: 22764,
					TotalPoints:  195,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 36540, FirstName: "Andreas", LastName: "Zelinka", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 18},
					Player2:      &volleynet.Player{ID: 51104, FirstName: "Simon", LastName: "Sladek", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 103},
					Deregistered: false,
					Seed:         8,
					TournamentID: 22764,
					TotalPoints:  121,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 18427, FirstName: "Michael", LastName: "Haas", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 98},
					Player2:      &volleynet.Player{ID: 39945, FirstName: "Werner", LastName: "Schmid", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 11},
					Deregistered: false,
					Seed:         9,
					TournamentID: 22764,
					TotalPoints:  109,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 55596, FirstName: "Thomas", LastName: "Müllner", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 25},
					Player2:      &volleynet.Player{ID: 43098, FirstName: "Sebastian", LastName: "Lechner", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 25},
					Deregistered: false,
					Seed:         10,
					TournamentID: 22764,
					TotalPoints:  50,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 55789, FirstName: "Martin", LastName: "Gschweidl", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 0},
					Player2:      &volleynet.Player{ID: 36557, FirstName: "Stefan", LastName: "Müller", Gender: "M", CountryUnion: "NÖVV", License: "1", TotalPoints: 0},
					Deregistered: false,
					Seed:         11,
					TournamentID: 22764,
					TotalPoints:  0,
				},
			},
		},
	},

	{
		"../testdata/done3.html",
		mustParseDate("26.08.2018"),
		volleynet.TournamentInfo{ID: 22616, Gender: "M", Status: "upcoming"},

		&volleynet.Tournament{
			TournamentInfo: volleynet.TournamentInfo{
				ID:     22616,
				Phase:  "ABV Tour AMATEUR 1",
				Start:  mustParseDate("25.08.2018"),
				End:    mustParseDate("25.08.2018"),
				Gender: "M",
				Status: "done",
			},

			SignedupTeams: 4,
			MaxTeams:      16,
			Website:       "",
			Location:      "Strandbad 104 3400 Klosterneuburg",
			Mode:          "Double Elimination 16er-Raster",
			MaxPoints:     540,
			Organiser:     "SCHAFFER Felix",
			Phone:         "0676 587 444 0",
			Email:         "fschaffer@gmx.at",
			Teams: []*volleynet.TournamentTeam{
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 41275, FirstName: "Raphael", LastName: "Gruber", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 22590, FirstName: "Dominik", LastName: "Rieder", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       1,
					TournamentID: 22616,
					WonPoints:    42,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 6724, FirstName: "Robert", LastName: "Kirkovics", Gender: "M", CountryUnion: "NÖVV"},
					Player2:      &volleynet.Player{ID: 13089, FirstName: "Christian", LastName: "Karlin", Gender: "M", CountryUnion: "NÖVV"},
					Deregistered: false,
					Result:       2,
					TournamentID: 22616,
					WonPoints:    38,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 13917, FirstName: "Florian", LastName: "Böhm", Gender: "M", CountryUnion: "BVV"},
					Player2:      &volleynet.Player{ID: 51026, FirstName: "Stefan", LastName: "Dienst", Gender: "M", CountryUnion: "BVV"},
					Deregistered: false,
					Result:       3,
					TournamentID: 22616,
					WonPoints:    34,
				},
				&volleynet.TournamentTeam{
					Player1:      &volleynet.Player{ID: 45125, FirstName: "Bernhard", LastName: "Sirowy", Gender: "M", CountryUnion: "WVV"},
					Player2:      &volleynet.Player{ID: 36552, FirstName: "Dominik", LastName: "Koudela", Gender: "M", CountryUnion: "WVV"},
					Deregistered: false,
					Result:       4,
					TournamentID: 22616,
					WonPoints:    30,
				},
			},
		},
	},
}

func TestTournament(t *testing.T) {
	for _, tt := range tournamentTests {
		t.Run(tt.file, func(t *testing.T) {
			response, _ := os.Open(tt.file)

			tournament, err := Tournament(response, tt.now, &tt.tournament)

			if err != nil {
				t.Fatalf("Tournament() err: %s", err)
			}

			// ignore notes
			tournament.HTMLNotes = ""
			// ignoe currentpoints... TODO
			tournament.CurrentPoints = ""

			test.Compare(t, "Tournament(): %+v", tournament, tt.out)
		})
	}
}
