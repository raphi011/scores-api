package main

import (
	"encoding/json"
	"fmt"
	"go-test/dtos"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func MatchShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["matchId"])

	if err != nil {
		panic(err)
	}

	match := getMatch(matchID)

	WriteJson(w, match)
}

func MatchIndex(w http.ResponseWriter, r *http.Request) {
	matches := getMatches()

	WriteJson(w, matches)
}

func MatchCreate(w http.ResponseWriter, r *http.Request) {
	var newMatch dtos.CreateMatchDto

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newMatch); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	match, _ := createMatch(
		newMatch.Player1ID,
		newMatch.Player2ID,
		newMatch.Player3ID,
		newMatch.Player4ID,
		newMatch.ScoreTeam1,
		newMatch.ScoreTeam2,
	)

	w.WriteHeader(http.StatusCreated)
	WriteJson(w, match)
}

func PlayerCreate(w http.ResponseWriter, r *http.Request) {
	var newPlayer dtos.CreatePlayerDto

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newPlayer); err != nil {
		panic(err)
	}

	player, _ := createPlayer(newPlayer.Name)

	w.WriteHeader(http.StatusCreated)
	WriteJson(w, player)
}

func PlayerIndex(w http.ResponseWriter, r *http.Request) {
	players := getPlayers()

	WriteJson(w, players)
}
