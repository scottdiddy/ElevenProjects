package utils

import "moviesProject/src/database/models"

func InvalidIdResponse(data any) models.Response_Data {
	return models.Response_Data{
		Message: "Invalid Id. Please enter a valid Id",
		Data:    data,
	}
}
