package controllers

import (
	"encoding/json"
	"moviesProject/src/services"
	"moviesProject/src/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userRequestID := mux.Vars(r)
	movie, err := services.GetMovie(userRequestID["id"])
	if err == nil {
		w.WriteHeader(http.StatusOK)
		responseMessage := utils.ResponseMessage("Get movie request successful", movie)
		json.NewEncoder(w).Encode(responseMessage)
	} else {
		w.WriteHeader(http.StatusNotFound)
		responseMessage := utils.InvalidIdResponse(userRequestID["id"])
		json.NewEncoder(w).Encode(responseMessage)
	}
}
