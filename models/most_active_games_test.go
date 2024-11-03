package models

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchMostActiveGames(t *testing.T) {
	// Mock the HTTP response
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"results": [{"name": "Game 1", "released": "2023-11-01"}, {"name": "Game 2", "released": "2023-12-01"}]}`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	games, err := FetchMostActiveGames()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
