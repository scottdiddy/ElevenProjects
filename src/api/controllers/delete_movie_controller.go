package controllers

import (
	"encoding/json"
	"moviesProject/src/constants"
	"moviesProject/src/services"
	"moviesProject/src/utils"
	"net/http"
)

func DeleteMovieController(w http.ResponseWriter, r *http.Request) {
	contextBody := r.Context().Value(constants.AUTHORIZATION_KEY)
	token, ok := contextBody.(string)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	remainingMovies := services.DeleteMovieService(token)
	responseMessage := utils.ResponseMessage("Delete movie request successful", remainingMovies)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseMessage)
}
