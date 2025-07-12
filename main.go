package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ibrahimelothmani/ReelingIt/handlers"
	"github.com/ibrahimelothmani/ReelingIt/logger"
	"github.com/ibrahimelothmani/ReelingIt/providers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	movieRepo, err := providers.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize repository")
	}

	accountRepo, err := providers.NewAccountRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize account repository: %v", err)
	}

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	accountHandler := handlers.NewAccountHandler(accountRepo, logInstance)

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register", accountHandler.Register)
	http.HandleFunc("/api/account/authenticate", accountHandler.Authenticate)

	http.Handle("/api/account/favorites/",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetFavorites)))

	http.Handle("/api/account/watchlist/",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetWatchlist)))

	http.Handle("/api/account/save-to-collection/",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.SaveToCollection)))

	catchAllHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}
	http.HandleFunc("/movies", catchAllHandler)
	http.HandleFunc("/movies/", catchAllHandler)
	http.HandleFunc("/account/", catchAllHandler)

	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Server is running on http://localhost:8080")

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}

}
