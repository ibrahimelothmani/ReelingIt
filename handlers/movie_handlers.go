package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ibrahimelothmani/ReelingIt/logger"
	"github.com/ibrahimelothmani/ReelingIt/models"
)

type MovieHandler struct {
	logger  *logger.Logger
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:      1,
			TMDB_ID: 1000,
			Title:   "Hero",
		},
		{
			ID:      2,
			TMDB_ID: 2000,
			Title:   "Hero",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		// TODO : LOG ERROR
		// fmt.Println("ERROR...")
		http.Error(w, "Internal Server ERROR", http.StatusInternalServerError)
	}

}
