package main

import (
	"net/http"
	"log"
)

func main()  {
	handler := &PlayerServer{}
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
