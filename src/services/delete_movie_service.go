package services

import (
	"moviesProject/src/database/models"
	"moviesProject/src/database/repository"
)

func DeleteMovieService(userUUIDToken string) map[string]models.Movie {
	movieToDeleteID := repository.UsersUUID[userUUIDToken]
	delete(repository.Movies, movieToDeleteID)
	return repository.Movies
}
