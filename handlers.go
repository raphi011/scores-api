package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scores-backend/dtos"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func MatchShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["matchID"])

	if err != nil {
		respondError(w, err.Error(), http.StatusBadRequest)
	}

	match := getMatch(matchID)

	writeJSON(w, match, http.StatusOK)
}

func MatchDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["matchID"])

	if err != nil {
		respondError(w, err.Error(), http.StatusBadRequest)
	}

	deleteMatch(uint(matchID))

	w.WriteHeader(http.StatusNoContent)
}

func MatchIndex(w http.ResponseWriter, r *http.Request) {
	matches := getMatches()

	writeJSON(w, matches, http.StatusOK)
}

func MatchCreate(w http.ResponseWriter, r *http.Request) {
	var newMatch dtos.CreateMatchDto

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newMatch); err != nil {
		respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	match, _ := createMatch(
		newMatch.Player1ID,
		newMatch.Player2ID,
		newMatch.Player3ID,
		newMatch.Player4ID,
		newMatch.ScoreTeam1,
		newMatch.ScoreTeam2,
	)

	writeJSON(w, match, http.StatusCreated)
}

func PlayerCreate(w http.ResponseWriter, r *http.Request) {
	var newPlayer dtos.CreatePlayerDto

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newPlayer); err != nil {
		respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	player, _ := createPlayer(newPlayer.Name)

	writeJSON(w, player, http.StatusCreated)
}

func PlayerIndex(w http.ResponseWriter, r *http.Request) {
	players := getPlayers()

	writeJSON(w, players, http.StatusOK)
}

func PlayerStatistic(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// playerID, err := strconv.Atoi(vars["playerID"])

	// if err != nil {
	// 	respondError(w, err.Error(), http.StatusBadRequest)
	// }

	// statistic := PlayerStatistic(uint(playerID))
	var statistic dtos.Statistic

	writeJSON(w, statistic, http.StatusOK)
}
