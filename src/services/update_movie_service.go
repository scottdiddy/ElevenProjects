package services

import (
	"errors"
	"moviesProject/src/constants"
	"moviesProject/src/database/models"
	"moviesProject/src/database/repository"
	"moviesProject/src/utils/helpers"
	"net/http"
)

func UpdateMovieService(w http.ResponseWriter, r *http.Request) (models.Movie, error) {
	contextBody := r.Context().Value(constants.UPDATE_MOVIE_KEY)
	updateFields, ok := contextBody.(models.UpdateMovieDTO)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return models.Movie{}, errors.New("internal server error")
	}
	contextBody = r.Context().Value(constants.AUTHORIZATION_KEY)
	userUUIDToken, ok := contextBody.(string)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return models.Movie{}, errors.New("internal server error")
	}
	movieToUpdateID := repository.UsersUUID[userUUIDToken]
	movieToUpdate := repository.Movies[movieToUpdateID]
	helpers.CopyStructFields(&movieToUpdate, &updateFields)
	return movieToUpdate, nil

}
