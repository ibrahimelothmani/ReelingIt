package main

import (
	"log"
	"net/http"

	"github.com/ibrahimelothmani/ReelingIt/handlers"
	"github.com/ibrahimelothmani/ReelingIt/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	logInstance.Error("Hello from the Error system", nil)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {

	logInstance := initializeLogger()

	movieHandler := handlers.MovieHandler{}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("public/index.html")))

	// Start server
	const addr = ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}
}
