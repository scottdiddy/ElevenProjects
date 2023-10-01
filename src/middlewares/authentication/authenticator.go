package authentication

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"moviesProject/src/constants"
	"moviesProject/src/database/repository"
	"moviesProject/src/utils"
)

func MiddlewareAuthenticateUserToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		_, token, _ := strings.Cut(auth, " ")
		_, ok := repository.UsersUUID[token]
		if !ok {
			responseMessage := utils.ResponseMessage("Wrong token. Access denied", token)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(responseMessage)
			return
		}
		ctx := context.WithValue(r.Context(), constants.AUTHORIZATION_KEY, token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
