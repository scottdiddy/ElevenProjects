package controllers

import (
	"encoding/json"
	"moviesProject/src/services"
	"moviesProject/src/utils"
	"net/http"
)

func UpdateMovieController(w http.ResponseWriter, r *http.Request) {
	updatedMovie, err := services.UpdateMovieService(w, r)
	if err != nil {
		return
	}
	responseMessage := utils.ResponseMessage("Update movie request successful", updatedMovie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseMessage)
}
