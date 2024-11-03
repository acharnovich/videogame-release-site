package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"videogame-release-site/models"
)

func UpcomingGamesHandler(w http.ResponseWriter, r *http.Request) {
	games, err := models.FetchUpcomingGames()
	if err != nil {
		log.Printf("Error fetching games: %v", err)
		http.Error(w, "Error fetching games", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("views", "upcoming_games.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	if err := tmpl.Execute(w, games); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func MostActiveGamesHandler(w http.ResponseWriter, r *http.Request) {
	games, err := models.FetchMostActiveGames()
	if err != nil {
		log.Printf("Error fetching games: %v", err)
		http.Error(w, "Error fetching games", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("views", "most_active_games.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	if err := tmpl.Execute(w, games); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}
