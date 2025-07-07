package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ibrahimelothmani/ReelingIt/handlers"
	"github.com/ibrahimelothmani/ReelingIt/logger"
	"github.com/joho/godotenv"
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

	// Log Initializer
	logInstance := initializeLogger()

	// Database Initializer
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	defer db.Close()

	// Movie Handler
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
