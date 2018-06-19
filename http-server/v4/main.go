package main

import (
	"net/http"
	"log"
)

type InMemoryPlayerStore struct {

}

func (i InMemoryPlayerStore) GetPlayerScore(name string) (score int)  {
	score = 123
	return
}

func main()  {
	handler := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
