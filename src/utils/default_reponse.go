package utils

import "moviesProject/src/database/models"

func ResponseMessage(message string, data any) models.Response_Data {
	return models.Response_Data{
		Message: message,
		Data:    data,
	}
}
