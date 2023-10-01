package services

import (
	"moviesProject/src/constants"
	"moviesProject/src/database/models"
	"moviesProject/src/database/repository"
	"moviesProject/src/utils/helpers"
	"net/http"
)

func CreateMovieService(w http.ResponseWriter, r *http.Request, userTokenChannel chan string) {
	userToken := helpers.GenerateUUID()
	contextBody := r.Context().Value(constants.CREATE_MOVIE_KEY)
	newMovieDTO, ok := contextBody.(models.CreateMovieDTO)
	if ok {
		userTokenChannel <- userToken
	} else {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var newMovie models.Movie
	newMovie.Id = helpers.GenerateID()
	helpers.CopyStructFields(&newMovie, &newMovieDTO)
	repository.Movies[newMovie.Id] = newMovie
	repository.UsersUUID[userToken] = newMovie.Id
}
