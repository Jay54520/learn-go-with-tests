package main

import (
	"net/http"
	"fmt"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) (score string) {
	if player == "Pepper" {
		score = "20"
	} else if player == "Floyd" {
		score = "10"
	}
	return
}

type PlayerStore interface {
	GetPlayerScore(player string) (score int)
}
