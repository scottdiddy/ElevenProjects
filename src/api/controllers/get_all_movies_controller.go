package controllers

import (
	"encoding/json"
	"moviesProject/src/services"
	"net/http"
)

func GetMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services.GetMovies())
}
