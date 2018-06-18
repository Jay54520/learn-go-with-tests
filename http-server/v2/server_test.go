package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestGETPlayers(t *testing.T) {

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func assertResponseBody(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func newGetScoreRequest(player string) (*http.Request, error)  {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
}
