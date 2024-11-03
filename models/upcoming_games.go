package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Game struct {
	Name     string `json:"name"`
	Released string `json:"released"`
}

type ApiResponse struct {
	Results []Game `json:"results"`
}

func FetchUpcomingGames() ([]Game, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("RAWG_API_KEY environment variable not set")
	}
	now := time.Now()
	startDate := now.AddDate(0, 0, -7).Format("2006-01-02")
	endDate := now.AddDate(0, 1, 0).Format("2006-01-02")

	url := fmt.Sprintf("https://api.rawg.io/api/games?key=%s&dates=%s,%s&ordering=-added", apiKey, startDate, endDate)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Results, nil
}
