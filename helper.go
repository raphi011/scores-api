package main

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		panic(err)
	}
}
