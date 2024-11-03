package main

import (
	"log"
	"net/http"
	"videogame-release-site/controllers"
)

func main() {
	http.HandleFunc("/", controllers.UpcomingGamesHandler)
	http.HandleFunc("/most-active", controllers.MostActiveGamesHandler)
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
