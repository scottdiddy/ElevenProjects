package controllers

import (
	"encoding/json"
	"moviesProject/src/services"
	"moviesProject/src/utils"
	"net/http"
)

func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	userTokenChannnel := make(chan string)
	go services.CreateMovieService(w, r, userTokenChannnel)
	responseMessage := utils.ResponseMessage("Successfully created movie. Unique bearer token below", <-userTokenChannnel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseMessage)
}
