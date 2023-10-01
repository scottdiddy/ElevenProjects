package services

import (
	"errors"
	"moviesProject/src/database/models"
	"moviesProject/src/database/repository"
)

func GetMovies() map[string]models.Movie {
	return repository.Movies
}
func GetMovie(id string) (models.Movie, error) {
	movie, ok := repository.Movies[id]
	if ok {
		return movie, nil
	} else {
		return models.Movie{}, errors.New("invalid Id")
	}
}
