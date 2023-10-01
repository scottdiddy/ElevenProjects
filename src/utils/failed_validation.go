package utils

import "moviesProject/src/database/models"

func Failed_Validation(validationErrors map[string]string) models.Response_Data {
	return models.Response_Data{
		Message: "Data validation failed. Re-check fields",
		Data:    validationErrors,
	}

}
