package repository

import "moviesProject/src/database/models"

var Movies = make(map[string]models.Movie)
var UsersUUID = make(map[string]string)
