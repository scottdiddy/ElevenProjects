package validation

import (
	"context"
	"encoding/json"
	"moviesProject/src/constants"
	"moviesProject/src/database/models"
	"moviesProject/src/utils"
	"net/http"
)

func ValidatorforCreateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userMovie = models.CreateMovieDTO{}
		if err := json.NewDecoder(r.Body).Decode(&userMovie); err != nil {
			http.Error(w, "Invalid JSON request", http.StatusBadRequest)
			return
		}
		validate := NewValidatorMiddleware()
		if requestDataValidationErrors, err, ok := validate.ValidateStruct(userMovie); err != nil {
			if ok {
				responseMessage := utils.Failed_Validation(requestDataValidationErrors)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(responseMessage)
				return
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		ctx := context.WithValue(r.Context(), constants.CREATE_MOVIE_KEY, userMovie)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
