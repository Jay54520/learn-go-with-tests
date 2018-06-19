package main

import (
	"net/http"
	"fmt"
)

type PlayerStore interface {
	GetPlayerScore(player string) (score int)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}