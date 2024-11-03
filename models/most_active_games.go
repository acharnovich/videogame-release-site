package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MostActiveGame struct {
	Name     string `json:"name"`
	Released string `json:"released"`
}

type MostActiveGamesResponse struct {
	Results []MostActiveGame `json:"results"`
}

func FetchMostActiveGames() ([]MostActiveGame, error) {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://api.rawg.io/api/games?key=%s&ordering=-metacritic", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse MostActiveGamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Results, nil
}
