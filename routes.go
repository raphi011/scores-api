package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MatchIndex",
		"GET",
		"/matches",
		MatchIndex,
	},
	Route{
		"MatchShow",
		"GET",
		"/matches/{matchID}",
		MatchShow,
	},
	Route{
		"PlayerIndex",
		"GET",
		"/players",
		PlayerIndex,
	},
	Route{
		"PlayerCreate",
		"POST",
		"/players",
		PlayerCreate,
	},
	Route{
		"MatchCreate",
		"POST",
		"/matches",
		MatchCreate,
	},
	Route{
		"MatchDelete",
		"DELETE",
		"/matches/{matchID}",
		MatchDelete,
	},
	Route{
		"PlayerStatistic",
		"GET",
		"/players/{playerID}/statistic",
		PlayerStatistic,
	},
}
