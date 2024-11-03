package models

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchUpcomingGames(t *testing.T) {
	// Mock the HTTP response
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"results": [{"name": "Game 1", "released": "2023-11-01"}, {"name": "Game 2", "released": "2023-12-01"}]}`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	games, err := FetchUpcomingGames()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(games) != 2 {
		t.Fatalf("Expected 2 games, got %d", len(games))
	}

	if games[0].Name != "Game 1" {
		t.Errorf("Expected 'Game 1', got %s", games[0].Name)
	}
}
